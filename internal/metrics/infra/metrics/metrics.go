package metrics

import (
	"github.com/prometheus/client_golang/prometheus"
)

var metrics = []prometheus.Collector{}

func Register() {
	prometheus.MustRegister(metrics...)
}

func Add(cs ...prometheus.Collector) {
	metrics = append(metrics, cs...)
}
