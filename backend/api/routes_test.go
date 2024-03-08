package api

import (
	"net/http"
	"testing"

	"github.com/papaya147/money-tracker/backend/test"
)

func TestRoutes(t *testing.T) {
	router := testApp.router
	apiVersion := 1

	apiRoutes := []test.Route{
		{Path: "/swagger/*", Method: http.MethodGet},

		{Path: "/expenditure/", Method: http.MethodPost},

		{Path: "/expenditure/category/", Method: http.MethodPost},
		{Path: "/expenditure/category/bulk", Method: http.MethodPost},
		{Path: "/expenditure/category/", Method: http.MethodGet},
		{Path: "/expenditure/category/{category-name}", Method: http.MethodPut},
		{Path: "/expenditure/category/{category-name}", Method: http.MethodDelete},
	}

	for _, route := range apiRoutes {
		test.RouteExists(t, router, apiVersion, route)
	}
}
