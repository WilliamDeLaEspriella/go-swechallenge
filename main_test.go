package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/WilliamDeLaEspriella/go-swechallenge/app"
	"github.com/go-playground/assert"
)

func runTestServer() *httptest.Server {
	var server app.Server
	server.CreateConnection()
	server.CreateTables()
	server.CreateRoutes()
	return httptest.NewServer(server.Routes)
}

func Test_post_api_integration_tests_store_endpoint(t *testing.T) {
	ts := runTestServer()
	defer ts.Close()

	t.Run("it should return 200 when health is ok", func(t *testing.T) {
		resp, err := http.Get(fmt.Sprintf("%s/rating_changes", ts.URL))

		if err != nil {
			t.Fatalf("Expected no error, got %v", err)
		}

		assert.Equal(t, 200, resp.StatusCode)
	})

}
