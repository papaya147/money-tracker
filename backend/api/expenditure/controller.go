package expenditure

import (
	"github.com/go-chi/chi/v5"
	"github.com/papaya147/money-tracker/backend/api/expenditure/category"
	db "github.com/papaya147/money-tracker/backend/db/sqlc"
	"github.com/papaya147/money-tracker/backend/util"
)

type Controller struct {
	config *util.Config
	store  db.Store

	categoryController *category.Controller
}

func NewController(config *util.Config, store db.Store) *Controller {
	controller := &Controller{
		config: config,
		store:  store,
	}

	controller.categoryController = category.NewController(config, store)

	return controller
}

func (c *Controller) Routes() *chi.Mux {
	router := chi.NewMux()

	router.Mount("/category", c.categoryController.Routes())
	router.Post("/", c.create)
	router.Get("/", c.get)
	router.Put("/{expenditure-id}", c.update)
	router.Delete("/{expenditure-id}", c.delete)

	return router
}
