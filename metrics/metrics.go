package metrics

import (
	"github.com/prometheus/client_golang/prometheus"
)

var (
	Status200 = prometheus.NewCounter(prometheus.CounterOpts{
		Name: "http_200_status_code_response",
	})
	Status401 = prometheus.NewCounter(prometheus.CounterOpts{
		Name: "http_401_status_code_response",
	})
	Status404 = prometheus.NewCounter(prometheus.CounterOpts{
		Name: "http_404_status_code_response",
	})
)
