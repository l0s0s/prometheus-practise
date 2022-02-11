package metrics_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/l0s0s/prometheus-practise/metrics"
	"github.com/l0s0s/prometheus-practise/test/helper"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func newMockMetricsAPI(h *metrics.Handler) *httptest.Server {
	gin.SetMode(gin.ReleaseMode)

	r := gin.New()

	h.BindRoutes(&r.RouterGroup)

	return httptest.NewServer(r)
}

func TestMetrics(t *testing.T) {
	ts := newMockMetricsAPI(metrics.NewHandler())

	defer ts.Close()

	resp, err := helper.MakeRequest(ts.URL + "/metrics")
	require.NoError(t, err)

	assert.Equal(t, http.StatusOK, resp.StatusCode)

	assert.NotEmpty(t, resp.Body)
}
