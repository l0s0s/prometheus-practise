package helper

import (
	"bytes"
	"context"
	"net/http"
)

// MakeRequest make request to test server.
func MakeRequest(url string) (*http.Response, error) {
	request, err := http.NewRequestWithContext(context.Background(), http.MethodGet, url, &bytes.Reader{})
	if err != nil {
		return nil, err
	}

	client := http.Client{}

	return client.Do(request)
}
