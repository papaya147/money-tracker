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

func RouteExists(t *testing.T, routes chi.Router, apiVersion int, path Path, method Method) {
	found := false

	_ = chi.Walk(routes, func(foundMethod string, foundRoute string, handler http.Handler, middlewares ...func(http.Handler) http.Handler) error {
		if fmt.Sprintf("/api/v%d%s", apiVersion, path) == foundRoute && method.String() == foundMethod {
			found = true
		}
		return nil
	})

	if !found {
		t.Errorf("did not find /api/v%d%s with method %s in registered routes", apiVersion, path, method)
	}
}
