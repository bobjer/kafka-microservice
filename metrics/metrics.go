package metrics

import (
	"net/http"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var (
	ProcessedMessages = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "processed_messages_total",
			Help: "Total number of processed messages",
		},
		[]string{"status"},
	)
)

func Init() {
	prometheus.MustRegister(ProcessedMessages)
}

func Handler() http.Handler {
	return promhttp.Handler()
}
