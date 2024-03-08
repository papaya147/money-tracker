package main

import (
	"fmt"

	"github.com/papaya147/money-tracker/backend/api"
	db "github.com/papaya147/money-tracker/backend/db/sqlc"
	"github.com/papaya147/money-tracker/backend/util"
)

// @title Money Tracker APIs
// @version 1.0
// @description The comprehensive list of all Money Tracker APIs
// @host localhost:4000
// @BasePath /api/v1
func main() {
	config := util.LoadEnv(".")

	conn := util.CreatePostgresPool(config.POSTGRES_DSN)
	defer conn.Close()

	store := db.NewStore(conn)

	app := api.NewServer(config, store)

	fmt.Printf("starting server on port %s\n", config.PORT)
	if err := app.Start(); err != nil {
		fmt.Printf("error starting server: %s\n", err)
	}
}
