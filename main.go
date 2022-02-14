package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/l0s0s/prometheus-practise/metrics"
	"github.com/l0s0s/prometheus-practise/mock"
)

func main() {
	r := gin.New()

	metricsHandler := metrics.NewHandler()
	mockHandler := mock.NewHandler()

	metricsHandler.BindRoutes(&r.RouterGroup)
	r.Use(metrics.ReadResponseStatusCode())
	mockHandler.BindRoutes(&r.RouterGroup)

	server := http.Server{
		Addr:    ":8080",
		Handler: r,
	}

	if err := server.ListenAndServe(); err != nil {
		log.Panic(err)
	}
}
