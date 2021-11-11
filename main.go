package main

import (
	"context"
	"database/sql"
	"net/http"
	"urlshort/core/composer"
	"urlshort/storage/pg"
	"urlshort/transport"

	_ "github.com/lib/pq"
)

func main() {
	db := getDBOrPanic()

	s := pg.NewStorage(db)
	c := composer.NewMd5Base64()
	shortener := transport.NewHTTPShortener(s, c)

	http.HandleFunc("/create", shortener.Create)
	http.HandleFunc("/", shortener.Open)

	if err := http.ListenAndServe(":8000", http.DefaultServeMux); err != nil {
		panic(err)
	}
}

func getDBOrPanic() *sql.DB {
	dataSource := "postgres://postgres:mypassword@localhost/postgres?sslmode=disable"
	db, err := sql.Open("postgres", dataSource)
	if err != nil {
		panic(err)
	}

	if err = db.PingContext(context.Background()); err != nil {
		panic(err)
	}
	return db
}
