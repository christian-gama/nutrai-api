package http

import (
	"github.com/christian-gama/nutrai-api/internal/metrics/infra/metrics"
	"github.com/prometheus/client_golang/prometheus"
)

func init() {
	metrics.Add(RequestsTotal, RequestsDuration)
}

var (
	RequestsTotal = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "http_requests_total",
			Help: "Number of HTTP requests",
		},
		[]string{"path"},
	)

	RequestsDuration = prometheus.NewHistogramVec(
		prometheus.HistogramOpts{
			Name:    "http_requests_duration_seconds",
			Help:    "Duration of HTTP requests",
			Buckets: prometheus.DefBuckets,
		},
		[]string{"path"},
	)
)
