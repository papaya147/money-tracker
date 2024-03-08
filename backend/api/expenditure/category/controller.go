package category

import (
	"github.com/go-chi/chi/v5"
	db "github.com/papaya147/money-tracker/backend/db/sqlc"
	"github.com/papaya147/money-tracker/backend/util"
)

type Controller struct {
	config *util.Config
	store  db.Store
}

func NewController(config *util.Config, store db.Store) *Controller {
	controller := &Controller{
		config: config,
		store:  store,
	}

	return controller
}

func (c *Controller) Routes() *chi.Mux {
	router := chi.NewMux()

	router.Post("/", c.create)

	return router
}
