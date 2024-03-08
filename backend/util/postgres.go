package util

import (
	"context"
	"fmt"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
)

func CreatePostgresPool(dsn string) *pgxpool.Pool {
	count := 0
	for {
		db, err := pgxpool.New(context.Background(), dsn)
		if err != nil {
			count++
		} else {
			fmt.Println("connected to postgres!")
			return db
		}
		if count == 5 {
			fmt.Println("unable to connect to postgres...", err)
			fmt.Println("retying in 5 seconds...")
			time.Sleep(time.Second * 5)
			count = 0
		}
	}
}
