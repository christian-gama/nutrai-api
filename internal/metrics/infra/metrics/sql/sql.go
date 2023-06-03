package sql

import (
	"github.com/christian-gama/nutrai-api/internal/metrics/infra/metrics"
	"github.com/prometheus/client_golang/prometheus"
)

func init() {
	metrics.Add(QueriesTotal)
}

var QueriesTotal = prometheus.NewCounter(
	prometheus.CounterOpts{
		Name: "sql_queries_total",
		Help: "Number of database queries (SQL)",
	},
)
