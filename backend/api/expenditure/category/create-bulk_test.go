package category

import (
	"encoding/json"
	"net/http"
	"testing"

	"github.com/papaya147/money-tracker/backend/test"
	"github.com/papaya147/money-tracker/backend/util"
)

func TestCreateBulk(t *testing.T) {
	postBody := map[string]interface{}{
		"categories": []string{util.RandomString(10), util.RandomString(20)},
	}
	body, _ := json.Marshal(postBody)
	test.TestCase(t, http.MethodPost, "/bulk", testHandler.createBulk, body, http.StatusOK)
}
