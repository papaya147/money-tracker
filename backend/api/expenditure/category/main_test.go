package category

import (
	"os"
	"testing"

	db "github.com/papaya147/money-tracker/backend/db/sqlc"
	"github.com/papaya147/money-tracker/backend/util"
)

var testHandler *Controller

func TestMain(m *testing.M) {
	config := util.LoadEnv("../../../")

	store := db.NewMockStore()

	testHandler = NewController(config, store)

	os.Exit(m.Run())
}
