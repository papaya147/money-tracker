package expenditure

import (
	"net/http"
	"testing"

	"github.com/papaya147/money-tracker/backend/test"
)

func TestDelete(t *testing.T) {
	test.TestCase(t, http.MethodDelete, "/{expenditure-id}", testHandler.delete, nil, http.StatusOK)
}
