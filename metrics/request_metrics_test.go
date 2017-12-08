package metrics_test

import (
	"time"

	"code.cloudfoundry.org/clock/fakeclock"
	mfakes "code.cloudfoundry.org/diego-logging-client/testhelpers"
	loggregator "code.cloudfoundry.org/go-loggregator"
	"code.cloudfoundry.org/go-loggregator/rpc/loggregator_v2"
	"code.cloudfoundry.org/lager/lagertest"
	"code.cloudfoundry.org/locket/metrics"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/onsi/gomega/gstruct"
	"github.com/onsi/gomega/types"
	"github.com/tedsuo/ifrit"
	"github.com/tedsuo/ifrit/ginkgomon"
)

var _ = Describe("RequestMetrics", func() {
	type metric struct {
		Name  string
		Value interface{}
		Opts  []loggregator.EmitGaugeOption
	}

	var (
		runner           *metrics.RequestMetricsNotifier
		process          ifrit.Process
		fakeMetronClient *mfakes.FakeIngressClient
		logger           *lagertest.TestLogger
		fakeClock        *fakeclock.FakeClock
		metricsInterval  time.Duration
		metricsChan      chan metric
	)

	BeforeEach(func() {
		metricsChan = make(chan metric, 100)
		fakeMetronClient = new(mfakes.FakeIngressClient)
		fakeMetronClient.SendMetricStub = func(name string, value int, opts ...loggregator.EmitGaugeOption) error {
			metricsChan <- metric{name, value, opts}
			return nil
		}
		fakeMetronClient.SendDurationStub = func(name string, value time.Duration, opts ...loggregator.EmitGaugeOption) error {
			metricsChan <- metric{name, value, opts}
			return nil
		}

		logger = lagertest.NewTestLogger("test")
		fakeClock = fakeclock.NewFakeClock(time.Now())
		metricsInterval = 10 * time.Second
	})

	JustBeforeEach(func() {
		runner = metrics.NewRequestMetricsNotifier(
			logger,
			fakeClock,
			fakeMetronClient,
			metricsInterval,
		)

		process = ifrit.Background(runner)
		Eventually(process.Ready()).Should(BeClosed())
	})

	AfterEach(func() {
		ginkgomon.Interrupt(process)
	})

	Context("when there is traffic to the locket server", func() {
		var requestType = "random-request"

		JustBeforeEach(func() {
			fakeClock.Increment(metricsInterval)
		})

		It("periodically emits the number of requests started", func() {
			runner.IncrementRequestsStartedCounter(requestType, 1)
			fakeClock.Increment(metricsInterval)
			Eventually(metricsChan).Should(Receive(gstruct.MatchAllFields(
				gstruct.Fields{
					"Name":  Equal("RequestsStarted"),
					"Value": Equal(1),
					"Opts":  haveTag("request-type", requestType),
				},
			)))
		})

		It("periodically emits the number of requests started for different request types", func() {
			runner.IncrementRequestsStartedCounter(requestType, 1)
			fakeClock.Increment(metricsInterval)
			Eventually(metricsChan).Should(Receive(gstruct.MatchAllFields(
				gstruct.Fields{
					"Name":  Equal("RequestsStarted"),
					"Value": Equal(1),
					"Opts":  haveTag("request-type", requestType),
				},
			)))
		})

		It("periodically emits the number of requests succeeded", func() {
			runner.IncrementRequestsSucceededCounter(requestType, 1)
			fakeClock.Increment(metricsInterval)
			Eventually(metricsChan).Should(Receive(gstruct.MatchAllFields(
				gstruct.Fields{
					"Name":  Equal("RequestsSucceeded"),
					"Value": Equal(1),
					"Opts":  haveTag("request-type", requestType),
				},
			)))
		})

		It("periodically emits the number of requests failed", func() {
			runner.IncrementRequestsFailedCounter(requestType, 2)
			fakeClock.Increment(metricsInterval)
			Eventually(metricsChan).Should(Receive(gstruct.MatchAllFields(
				gstruct.Fields{
					"Name":  Equal("RequestsFailed"),
					"Value": Equal(2),
					"Opts":  haveTag("request-type", requestType),
				},
			)))
		})

		It("periodically emits the number of requests in flight", func() {
			runner.IncrementRequestsInFlightCounter(requestType, 4)
			fakeClock.Increment(metricsInterval)
			Eventually(metricsChan).Should(Receive(gstruct.MatchAllFields(
				gstruct.Fields{
					"Name":  Equal("RequestsInFlight"),
					"Value": Equal(4),
					"Opts":  haveTag("request-type", requestType),
				},
			)))

			runner.DecrementRequestsInFlightCounter(requestType, 2)
			fakeClock.Increment(metricsInterval)
			Eventually(metricsChan).Should(Receive(gstruct.MatchAllFields(
				gstruct.Fields{
					"Name":  Equal("RequestsInFlight"),
					"Value": Equal(2),
					"Opts":  haveTag("request-type", requestType),
				},
			)))
		})

		It("periodically emits the max latency", func() {
			runner.UpdateLatency(requestType, 5*time.Millisecond)
			fakeClock.Increment(metricsInterval)
			Eventually(metricsChan).Should(Receive(gstruct.MatchAllFields(
				gstruct.Fields{
					"Name":  Equal("RequestLatencyMax"),
					"Value": Equal(5 * time.Millisecond),
					"Opts":  haveTag("request-type", requestType),
				},
			)))

			fakeClock.Increment(metricsInterval)
			Eventually(metricsChan).Should(Receive(gstruct.MatchAllFields(
				gstruct.Fields{
					"Name":  Equal("RequestLatencyMax"),
					"Value": Equal(time.Duration(0)),
					"Opts":  haveTag("request-type", requestType),
				},
			)))
		})
	})
})

func haveTag(name, value string) types.GomegaMatcher {
	return WithTransform(func(opts []loggregator.EmitGaugeOption) map[string]string {
		envelope := &loggregator_v2.Envelope{
			Tags: make(map[string]string),
		}
		for _, opt := range opts {
			opt(envelope)
		}
		return envelope.Tags
	}, Equal(map[string]string{name: value}))
}
