package locket

import (
	"errors"
	"fmt"
	"os"
	"time"

	"code.cloudfoundry.org/consuladapter"
	"github.com/hashicorp/consul/api"
	"github.com/pivotal-golang/clock"
	"github.com/pivotal-golang/lager"
)

type registrationRunner struct {
	logger        lager.Logger
	registration  *api.AgentServiceRegistration
	consulClient  consuladapter.Client
	retryInterval time.Duration
	clock         clock.Clock
}

func NewRegistrationRunner(
	logger lager.Logger,
	registration *api.AgentServiceRegistration,
	consulClient consuladapter.Client,
	retryInterval time.Duration,
	clock clock.Clock,
) *registrationRunner {
	return &registrationRunner{
		logger:        logger,
		registration:  registration,
		consulClient:  consulClient,
		retryInterval: retryInterval,
		clock:         clock,
	}
}

func (r *registrationRunner) validateRegistration() error {
	if r.registration.Checks != nil && len(r.registration.Checks) != 0 {
		// Implementing multiple service checks involves some nuance
		// around supporting multiple pollers for updating TTL-based healthchecks.
		// Since we don't need this at the moment, we'll leave this unimplemented.
		return errors.New("Support for multiple service checks not implemented")
	}

	return nil
}

func (r *registrationRunner) Run(signals <-chan os.Signal, ready chan<- struct{}) error {
	logger := r.logger.Session("registration-runner", lager.Data{"service": r.registration.Name})
	logger.Info("starting", lager.Data{"registration": r.registration})
	defer logger.Info("finished")

	// Fail early if the registration is invalid
	err := r.validateRegistration()
	if err != nil {
		logger.Error("failed-invalid-service-registration", err)
		return err
	}

	pollInterval, err := r.calculatePollInterval()
	if err != nil {
		logger.Error("failed-invalid-poll-interval", err)
		return err
	}

	agent := r.consulClient.Agent()
	errChan := make(chan error, 1)
	register := func() {
		logger.Info("attempting-registering-service")
		errChan <- agent.ServiceRegister(r.registration)
	}

	retryTimer := r.clock.NewTimer(0)

	for {
		select {
		case <-signals:
			logger.Info("received-signal")
			return nil
		case <-retryTimer.C():
			go register()
		case err := <-errChan:
			if err != nil {
				logger.Error("failed-registering-service", err)
				retryTimer.Reset(r.retryInterval)
			} else {
				logger.Info("succeeded-registering-service")
				retryTimer.Stop()
				agent.PassTTL(r.checkID(), "")
				close(ready)
				// If we have a TTL-based healthcheck, periodically send a heartbeat to the consul agent
				if pollInterval != 0 {
					return r.pollUntilSignaled(logger, pollInterval, signals)
				}

				return r.waitForSignal(logger, signals)
			}
		}
	}
}

func (r *registrationRunner) calculatePollInterval() (time.Duration, error) {
	if r.registration.Check == nil {
		return 0, nil
	}

	ttl := r.registration.Check.TTL
	if ttl == "" {
		return 0, nil
	}

	ttlDuration, err := time.ParseDuration(ttl)
	if err != nil {
		return 0, err
	}

	return ttlDuration / 2, nil
}

func (r *registrationRunner) waitForSignal(logger lager.Logger, signals <-chan os.Signal) error {
	agent := r.consulClient.Agent()
	<-signals

	logger.Info("deregistering-service")
	return agent.ServiceDeregister(r.registration.ID)
}

func (r *registrationRunner) checkID() string {
	// CheckID is automatically generated by consul based on the service ID
	// https://github.com/hashicorp/consul/blob/71e3901a6592817f9ebfd7f24f4ecff8ef16e7da/command/agent/agent.go#L770-L773
	// Service ID will default to the service name if one wasn't explicitly provided
	// https://github.com/hashicorp/consul/blob/71e3901a6592817f9ebfd7f24f4ecff8ef16e7da/command/agent/agent.go#L724-L726
	checkID := fmt.Sprintf("service:%s", r.registration.ID)
	if r.registration.ID == "" {
		checkID = fmt.Sprintf("service:%s", r.registration.Name)
	}
	return checkID
}

func (r *registrationRunner) unregisterID() string {
	if r.registration.ID == "" {
		return r.registration.Name
	}
	return r.registration.ID
}

func (r *registrationRunner) pollUntilSignaled(logger lager.Logger, interval time.Duration, signals <-chan os.Signal) error {
	logger = logger.Session("poll-until-signaled", lager.Data{"update-interval": interval.String()})
	logger.Info("started")
	defer logger.Info("finished")

	agent := r.consulClient.Agent()
	timer := r.clock.NewTimer(interval)
	checkID := r.checkID()
	errChan := make(chan error, 1)
	register := func() {
		logger.Info("attempting-registering-service")
		errChan <- agent.ServiceRegister(r.registration)
	}

	for {
		select {
		case <-signals:
			logger.Info("deregistering-service")
			return agent.ServiceDeregister(r.unregisterID())
		case <-timer.C():
			logger.Debug("updating-ttl-healthcheck", lager.Data{"checkID": checkID})
			err := agent.PassTTL(checkID, "")
			if err != nil {
				logger.Error("failed-healthcheck-in-consul", err)
				go register()
			}
			timer.Reset(interval)
		case err := <-errChan:
			if err != nil {
				logger.Error("failed-registering-service", err)
			} else {
				logger.Info("succeeded-registering-service")
			}
			timer.Reset(interval)
		}
	}
}
