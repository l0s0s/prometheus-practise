package mock

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Handler describe /api/v1/mock/* api.
type Handler struct{}

// NewHandler return new instance on Handler.
func NewHandler() *Handler {
	return &Handler{}
}

// Ok return 200 status code.
func (h *Handler) Ok(c *gin.Context) {
	c.Status(http.StatusOK)
}

// Unauthorised return 401 status code.
func (h *Handler) Unauthorised(c *gin.Context) {
	c.Status(http.StatusUnauthorized)
}

// NotFound return 404 status code.
func (h *Handler) NotFound(c *gin.Context) {
	c.Status(http.StatusNotFound)
}

// BindRoutes bind gin routes to handler methods.
func (h *Handler) BindRoutes(group *gin.RouterGroup) {
	group.GET("/api/v1/mock/ok", h.Ok)
	group.GET("/api/v1/mock/unauthorised", h.Unauthorised)
	group.GET("/api/v1/mock/notfound", h.NotFound)
}
