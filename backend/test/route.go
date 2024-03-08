package test

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/go-chi/chi/v5"
)

type Path string
type Method string

func (m Method) String() string {
	return string(m)
}

type Route struct {
	Path   Path
	Method Method
}

func RouteExists(t *testing.T, routes chi.Router, apiVersion int, route Route) {
	found := false

	_ = chi.Walk(routes, func(foundMethod string, foundRoute string, handler http.Handler, middlewares ...func(http.Handler) http.Handler) error {
		if fmt.Sprintf("/api/v%d%s", apiVersion, route.Path) == foundRoute && route.Method.String() == foundMethod {
			found = true
		}
		return nil
	})

	if !found {
		t.Errorf("did not find /api/v%d%s with method %s in registered routes", apiVersion, route.Path, route.Method)
	}
}
