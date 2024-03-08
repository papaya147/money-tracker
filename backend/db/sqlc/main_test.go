package db

import (
	"os"
	"testing"

	"github.com/papaya147/money-tracker/backend/util"
)

var testQueries Store

func TestMain(m *testing.M) {
	config := util.LoadEnv("../../")

	conn := util.CreatePostgresPool(config.POSTGRES_DSN)

	testQueries = NewStore(conn)

	os.Exit(m.Run())
}
