package mock

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Handler struct {
}

func NewHandler() *Handler {
	return &Handler{}
}

func (h *Handler) Ok(c *gin.Context) {
	c.Status(http.StatusOK)
}

func (h *Handler) Unauthorized(c *gin.Context) {
	c.Status(http.StatusUnauthorized)
}

func (h *Handler) NotFound(c *gin.Context) {
	c.Status(http.StatusNotFound)
}

func (h *Handler) BindRoutes(group *gin.RouterGroup) {
	group.GET("/api/v1/mock/ok", h.Ok)
	group.GET("/api/v1/mock/unauthorized", h.Unauthorized)
	group.GET("/api/v1/mock/notfound", h.NotFound)
}
