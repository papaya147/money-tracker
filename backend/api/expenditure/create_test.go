package expenditure

import (
	"encoding/json"
	"net/http"
	"testing"

	"github.com/papaya147/money-tracker/backend/test"
	"github.com/papaya147/money-tracker/backend/util"
)

func TestCreate(t *testing.T) {
	postBody := map[string]interface{}{
		"paisa":    util.RandomInt32(0, 100000),
		"category": util.RandomString(10),
	}
	body, _ := json.Marshal(postBody)
	test.TestCase(t, http.MethodPost, "/", testHandler.create, body, http.StatusOK)
}
