package metrics_test

import (
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/l0s0s/prometheus-practise/metrics"
	"github.com/l0s0s/prometheus-practise/mock"
	"github.com/l0s0s/prometheus-practise/test/helper"
	"github.com/prometheus/client_golang/prometheus"
	cm "github.com/prometheus/client_model/go"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func startTestServer() *httptest.Server {
	gin.SetMode(gin.ReleaseMode)

	h := mock.NewHandler()

	r := gin.New()
	r.Use(metrics.ReadResponseStatusCode())

	h.BindRoutes(&r.RouterGroup)

	return httptest.NewServer(r)
}

func getMetricValue(col prometheus.Collector) float64 {
	c := make(chan prometheus.Metric, 1)

	col.Collect(c)

	m := cm.Metric{}

	_ = (<-c).Write(&m)

	return *m.Counter.Value
}

func TestReadResponseStatusCode_RequestToOkEndpoint_increaseStatus200Metric(t *testing.T) {
	ts := startTestServer()

	defer ts.Close()

	_, err := helper.MakeRequest(ts.URL + "/api/v1/mock/ok")
	require.NoError(t, err)

	assert.Equal(t, float64(1), getMetricValue(metrics.Status200))
}

func TestReadResponseStatusCode_RequestToUnauthorithedEndpoint_increaseStatus401Metric(t *testing.T) {
	ts := startTestServer()

	defer ts.Close()

	_, err := helper.MakeRequest(ts.URL + "/api/v1/mock/unauthorithed")
	require.NoError(t, err)

	assert.Equal(t, float64(1), getMetricValue(metrics.Status401))
}

func TestReadResponseStatusCode_RequestToNotFoundEndpoint_increaseStatus404Metric(t *testing.T) {
	ts := startTestServer()

	defer ts.Close()

	_, err := helper.MakeRequest(ts.URL + "/api/v1/mock/notfound")
	require.NoError(t, err)

	assert.Equal(t, float64(1), getMetricValue(metrics.Status404))
}
