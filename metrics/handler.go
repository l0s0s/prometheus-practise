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

// Handler describe /metrics api.
type Handler struct{}

// NewHandler return new instance of Handler.
func NewHandler() *Handler {
	return &Handler{}
}

// Metrics return inforamtion about all metrics.
func (h *Handler) Metrics(c *gin.Context) {
	promhttp.Handler().ServeHTTP(c.Writer, c.Request)
}

// BindRoutes bind gin routes to handler methods.
func (h *Handler) BindRoutes(group *gin.RouterGroup) {
	group.GET("/metrics", h.Metrics)
}
