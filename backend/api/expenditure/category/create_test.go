package category

import (
	"encoding/json"
	"net/http"
	"testing"

	"github.com/papaya147/money-tracker/backend/test"
	"github.com/papaya147/money-tracker/backend/util"
)

func TestCreate(t *testing.T) {
	postBody := map[string]interface{}{
		"name": util.RandomString(10),
	}
	body, _ := json.Marshal(postBody)
	test.TestCase(t, http.MethodGet, "/", testHandler.create, body, http.StatusOK)
}
