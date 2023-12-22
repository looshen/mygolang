package prometheus

import (
	"time"

	"github.com/prometheus/client_golang/prometheus"
)

const (
	WorkNamespace = "default"
)

type ExcutionTimer struct {
	histo *prometheus.HistogramVec
	start time.Time
	last  time.Time
}

var (
	funcMetric = prometheus.NewHistogramVec(prometheus.HistogramOpts{
		Namespace: WorkNamespace,
		Name:      "execution_latency_seconds",
		Help:      "prometheus help",
		Buckets:   prometheus.ExponentialBuckets(0.001, 2, 15),
	}, []string{"step"})
)

func Register() {
	prometheus.Register(funcMetric)
}

func NewExcutionTimer(histo *prometheus.HistogramVec) *ExcutionTimer {
	now := time.Now()
	return &ExcutionTimer{
		histo: histo,
		start: now,
		last:  now,
	}
}

func NewTimer() *ExcutionTimer {
	return NewExcutionTimer(funcMetric)
}

func (t *ExcutionTimer) ObserveTotal() {
	(*t.histo).WithLabelValues("total").Observe(time.Now().Sub(t.start).Seconds())
}
