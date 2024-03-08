package category

import (
	"encoding/json"
	"net/http"
	"testing"

	"github.com/papaya147/money-tracker/backend/test"
	"github.com/papaya147/money-tracker/backend/util"
)

func TestUpdate(t *testing.T) {
	postBody := map[string]interface{}{
		"new_name": util.RandomString(10),
	}
	body, _ := json.Marshal(postBody)
	test.TestCase(t, http.MethodPut, "/{category-id}", testHandler.update, body, http.StatusOK)
}
