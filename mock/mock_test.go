package mock_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/l0s0s/prometheus-practise/mock"
	"github.com/l0s0s/prometheus-practise/test/helper"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func startTestServer(h *mock.Handler) *httptest.Server {
	gin.SetMode(gin.ReleaseMode)

	r := gin.New()

	h.BindRoutes(&r.RouterGroup)

	return httptest.NewServer(r)
}

func TestOk(t *testing.T) {
	ts := startTestServer(mock.NewHandler())

	defer ts.Close()

	resp, err := helper.MakeRequest(ts.URL + "/api/v1/mock/ok")
	require.NoError(t, err)

	assert.Equal(t, http.StatusOK, resp.StatusCode)
}

func TestUnauthorized(t *testing.T) {
	ts := startTestServer(mock.NewHandler())

	defer ts.Close()

	resp, err := helper.MakeRequest(ts.URL + "/api/v1/mock/unauthorized")
	require.NoError(t, err)

	assert.Equal(t, http.StatusUnauthorized, resp.StatusCode)
}

func TestNotFound(t *testing.T) {
	ts := startTestServer(mock.NewHandler())

	defer ts.Close()

	resp, err := helper.MakeRequest(ts.URL + "/api/v1/mock/notfound")
	require.NoError(t, err)

	assert.Equal(t, http.StatusNotFound, resp.StatusCode)
}
