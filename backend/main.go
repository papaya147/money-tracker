package main

import (
	"errors"
	"fmt"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
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
	fmt.Println(*config)

	conn := util.CreatePostgresPool(config.POSTGRES_DSN)
	defer conn.Close()

	fmt.Print("attempting database migration...")
	if err := runDbMigration(config); err != nil {
		fmt.Println(" database migration failed with error:", err)
	} else {
		fmt.Println(" database migration was successful!")
	}

	store := db.NewStore(conn)

	app := api.NewServer(config, store)

	fmt.Printf("starting server on port %s...", config.PORT)
	if err := app.Start(); err != nil {
		fmt.Println(" error starting server: ", err)
	}
}

func runDbMigration(config *util.Config) error {
	migration, err := migrate.New(config.MIGRATION_URL, config.POSTGRES_DSN)
	if err != nil {
		return err
	}

	if err := migration.Up(); err != nil && !errors.Is(err, migrate.ErrNoChange) {
		return err
	}

	return nil
}
