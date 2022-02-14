package metrics

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

var (
	// Status200 is the response metric with code 200.
	Status200 = promauto.NewCounter(prometheus.CounterOpts{
		Name: "http_200_status_code_response",
	})
	// Status401 is the response metric with code 401.
	Status401 = promauto.NewCounter(prometheus.CounterOpts{
		Name: "http_401_status_code_response",
	})
	// Status404  is the response metric with code 404.
	Status404 = promauto.NewCounter(prometheus.CounterOpts{
		Name: "http_404_status_code_response",
	})
)
