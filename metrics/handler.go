package metrics

import (
	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func init() {
	prometheus.MustRegister(
		Status200,
		Status401,
		Status404,
	)
}

type Handler struct{}

func NewHandler() *Handler {
	return &Handler{}
}

func (h *Handler) Metrics(c *gin.Context) {
	promhttp.Handler().ServeHTTP(c.Writer, c.Request)
}

func (h *Handler) BindRoutes(group *gin.RouterGroup) {
	group.GET("/metrics", h.Metrics)
}
