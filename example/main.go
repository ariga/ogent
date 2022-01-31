package main

import (
	"context"
	"flag"
	"log"
	"net/http"

	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql/schema"
	"github.com/ariga/ogent/example/pets"
	"github.com/ariga/ogent/example/pets/ogent"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	var args struct {
		Addr string
		DSN  string
	}
	flag.StringVar(&args.Addr, "addr", ":8080", "http address to listen")
	flag.StringVar(&args.DSN, "dsn", "file:ent?mode=memory&cache=shared&_fk=1", "dsn of database")
	flag.Parse()
	// Create ent client.
	client, err := pets.Open(dialect.SQLite, args.DSN)
	if err != nil {
		log.Fatal(err)
	}
	// Run the migrations.
	if err := client.Schema.Create(context.Background(), schema.WithAtlas(true)); err != nil {
		log.Fatal(err)
	}
	// Start listening.
	if err := http.ListenAndServe(args.Addr, ogent.NewServer(ogent.NewOgentHandler(client))); err != nil {
		log.Fatal(err)
	}
}
