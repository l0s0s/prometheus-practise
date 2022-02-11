package metrics

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func ReadResponseStatusCode() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()

		switch c.Writer.Status() {
		case http.StatusOK:
			Status200.Inc()
		case http.StatusUnauthorized:
			Status401.Inc()
		case http.StatusNotFound:
			Status404.Inc()
		}
	}
}
