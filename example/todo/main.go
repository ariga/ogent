package main

import (
	"context"
	"log"
	"net/http"

	"ariga.io/ogent/example/todo/ent"
	"ariga.io/ogent/example/todo/ent/ogent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql/schema"
	_ "github.com/mattn/go-sqlite3"
)

type handler struct {
	*ogent.OgentHandler
	client *ent.Client
}

func (h handler) MarkDone(ctx context.Context, params ogent.MarkDoneParams) (ogent.MarkDoneNoContent, error) {
	return ogent.MarkDoneNoContent{}, h.client.Todo.UpdateOneID(params.ID).SetDone(true).Exec(ctx)
}

func main() {
	// Create ent client.
	client, err := ent.Open(dialect.SQLite, "file:ent?mode=memory&cache=shared&_fk=1")
	if err != nil {
		log.Fatal(err)
	}
	// Run the migrations.
	if err := client.Schema.Create(context.Background(), schema.WithAtlas(true)); err != nil {
		log.Fatal(err)
	}
	// Create the handler.
	h := handler{
		OgentHandler: ogent.NewOgentHandler(client),
		client:       client,
	}
	// Start listening.
	srv, err := ogent.NewServer(h)
	if err != nil {
		log.Fatal(err)
	}
	if err := http.ListenAndServe(":8180", srv); err != nil {
		log.Fatal(err)
	}
}
