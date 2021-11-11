package main

import (
	"context"
	"database/sql"
	"net/http"
	"urlshort/composer"
	"urlshort/core"
	"urlshort/storage/pg"

	_ "github.com/lib/pq"
)

func main() {
	ctx := context.Background()

	dataSource := "postgres://postgres:mypassword@localhost/postgres?sslmode=disable"
	db, err := sql.Open("postgres", dataSource)
	if err != nil {
		panic(err)
	}

	if err = db.PingContext(ctx); err != nil {
		panic(err)
	}

	s := pg.NewStorage(db)
	c := composer.NewMd5Base64()
	shortener := core.NewShortener(s, c)

	http.HandleFunc("/create", shortener.Create)

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})

	if err := http.ListenAndServe(":8000", http.DefaultServeMux); err != nil {
		panic(err)
	}
}
