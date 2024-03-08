package expenditure

import (
	"net/http"
	"testing"

	"github.com/papaya147/money-tracker/backend/test"
)

func TestGet(t *testing.T) {
	test.TestCase(t, http.MethodGet, "/", testHandler.get, nil, http.StatusOK)
}
