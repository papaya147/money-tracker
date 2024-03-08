package api

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/papaya147/money-tracker/backend/api/expenditure"
	db "github.com/papaya147/money-tracker/backend/db/sqlc"
	"github.com/papaya147/money-tracker/backend/util"
)

type server struct {
	config *util.Config
	store  db.Store

	expenditureController *expenditure.Controller
	router                *chi.Mux
}

func NewServer(config *util.Config, store db.Store) *server {
	app := &server{
		config: config,
		store:  store,
	}

	app.expenditureController = expenditure.NewController(app.config, app.store)

	app.loadRoutes()

	return app
}

func (app *server) Start() error {
	srv := &http.Server{
		Addr:    fmt.Sprintf(":%s", app.config.PORT),
		Handler: app.router,
	}

	return srv.ListenAndServe()
}
