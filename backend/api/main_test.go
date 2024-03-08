package api

import (
	"os"
	"testing"

	"github.com/papaya147/money-tracker/backend/api/expenditure"
	db "github.com/papaya147/money-tracker/backend/db/sqlc"
	"github.com/papaya147/money-tracker/backend/util"
)

var testApp *server

func TestMain(m *testing.M) {
	config := util.LoadEnv("../")

	testApp = &server{
		config:                config,
		store:                 db.NewMockStore(),
		expenditureController: expenditure.NewController(nil, nil),
	}

	testApp.loadRoutes()

	os.Exit(m.Run())
}
