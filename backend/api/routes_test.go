package api

import (
	"net/http"
	"testing"

	"github.com/papaya147/money-tracker/backend/test"
)

func TestRoutes(t *testing.T) {
	router := testApp.router
	apiVersion := 1

	apiRoutes := map[test.Path]test.Method{
		"/swagger/*":             http.MethodGet,
		"/expenditure/category/": http.MethodPost,
	}

	for route, method := range apiRoutes {
		test.RouteExists(t, router, apiVersion, route, method)
	}
}
