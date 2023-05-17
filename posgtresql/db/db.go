package db

import (
	"github.com/go-pg/pg/v10"
	"os"
)

func startDB() (*pg.DB, error) {

	var (
		options *pg.Options
		err     error
	)
	if os.Getenv("ENV") == "PROD" {
		options, err = pg.ParseURL(os.Getenv("DATABASE_URL"))
		if err != nil {
			return nil, err
		}
	} else {
		options = &pg.Options{
			Addr:     "localhost:5432",
			User:     os.Getenv("DB_USER"),
			Password: os.Getenv("DB_PASSWORD"),
		}
	}

	db := pg.Connect(options)
	collection := migrations
}
