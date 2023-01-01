package main

import (
	"context"
	"database/sql"
	"flag"
	"fmt"
	"log"
	"net/http"
	"strings"

	"ariga.io/ogent/example/pets/ent"
	"ariga.io/ogent/example/pets/ent/ogent"
	"entgo.io/ent/dialect"
	entsql "entgo.io/ent/dialect/sql"
	"github.com/go-faster/jx"
	_ "github.com/mattn/go-sqlite3"
)

// handler wraps the generated ogent.OgentHandler and overrides / adds http.Handler methods.
type handler struct {
	*ogent.OgentHandler
	db     *sql.DB
	client *ent.Client
}

// DBHealth sends a ping to the database and returns either HTTP 503 when it is not reachable or HTTP 204 when it is.
func (h handler) DBHealth(context.Context) (ogent.DBHealthRes, error) {
	if err := h.db.Ping(); err != nil {
		return &ogent.DBHealthServiceUnavailable{}, nil
	}
	return &ogent.DBHealthNoContent{}, nil
}

// CreateCategory "overrides" the generated ogent generated CreateCategory method.
func (h handler) CreateCategory(ctx context.Context, req *ogent.CreateCategoryReq) (ogent.CreateCategoryRes, error) {
	b := h.client.Category.Create()
	// Add the name field. Sanitize it beforehand.
	b.SetName(strings.TrimSpace(req.Name))
	// Persist to storage.
	e, err := b.Save(ctx)
	if err != nil {
		var e jx.Encoder
		e.Str(err.Error())
		switch {
		case ent.IsNotSingular(err):
			return &ogent.R409{
				Code:   http.StatusConflict,
				Status: http.StatusText(http.StatusConflict),
				Errors: e.Bytes(),
			}, nil
		case ent.IsConstraintError(err):
			return &ogent.R409{
				Code:   http.StatusConflict,
				Status: http.StatusText(http.StatusConflict),
				Errors: e.Bytes(),
			}, nil
		default:
			// Let the server handle the error.
			return nil, err
		}
	}
	// Do some other custom logic like dispatching an event to send some mails.
	fmt.Println("I sent an email")
	return ogent.NewCategoryCreate(e), nil
}

func main() {
	var args struct {
		Addr string
		DSN  string
	}
	flag.StringVar(&args.Addr, "addr", ":8080", "http address to listen")
	flag.StringVar(&args.DSN, "dsn", "file:ent?mode=memory&cache=shared&_fk=1", "dsn of database")
	flag.Parse()
	// Create ent client.
	drv, err := entsql.Open(dialect.SQLite, args.DSN)
	if err != nil {
		log.Fatal(err)
	}
	client := ent.NewClient(ent.Driver(drv))
	// Run the migrations.
	if err := client.Schema.Create(context.Background()); err != nil {
		log.Fatal(err)
	}
	// Start listening.
	srv, err := ogent.NewServer(handler{
		OgentHandler: ogent.NewOgentHandler(client),
		db:           drv.DB(),
		client:       client,
	})
	if err != nil {
		log.Fatal(err)
	}
	if err := http.ListenAndServe(args.Addr, srv); err != nil {
		log.Fatal(err)
	}
}
