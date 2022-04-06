# ogent

This package is an extension to the awesome [Ent](https://entgo.io) ORM code generator. It utilizes the power
of [`ogen`](https://github.com/ogen-go/ogen) to provide a type-safe, reflection free implementation of the OpenAPI
Specification document generated by another popular Ent extension: [`entoas`](https://github.com/ent/contrib/entoas)
. `ogent` generated the missing handler implementations needed by `ogen` to serve the described API. The data access is
done with Ent. To learn more about [Ent](https://entgo.io), how to connect to different types of databases, run
migrations or work the Graph model head over to their [documentation](https://entgo.io/docs/getting-started).

`ogen` is an opinionated Go code generator for OpenAPI Specification v3 documents. `ogen` generates both server and
client implementations for a given OpenAPI Specification document. The only thing left to do for the user is to
implement an interface to access the data layer of any application. `ogen` has many cool features, one of which is
integration with [OpenTelemetry](https://opentelemetry.io/). `ogent` serves as a bridge between [Ent](https://entgo.io)
and the code generated by [`ogen`](https://github.com/ogen-go/ogen). It uses the configuration by the
popular [Ent](https://entgo.io) extension `entoas`(https://github.com/ent/contrib/entoas) to generate the missing parts
of the `ogent` code.

## Getting Started

The first step is to add the `ogent` package to your Ent project.

```shell
go get ariga.io/ogent@main
```

`ogent` uses the Ent [Extension API](https://entgo.io/docs/extensions) to integrate with Ent’s code-generation. This
requires that you use the `entc` (ent codegen) package as
described [here](https://entgo.io/docs/code-gen#use-entc-as-a-package). Follow the next three steps to enable it and to
configure Ent to work with the `ogent` extension:

1. Create a new Go file name `ent/entc.go` and paste the following content:

```go
//go:build ignore

package main

import (
	"log"

	"ariga.io/ogent"
	"entgo.io/contrib/entoas"
	"entgo.io/ent/entc"
	"entgo.io/ent/entc/gen"
	"github.com/ogen-go/ogen"
)

func main() {
	spec := new(ogen.Spec)
	oas, err := entoas.NewExtension(entoas.Spec(spec))
	if err != nil {
		log.Fatalf("creating entoas extension: %v", err)
	}
	ogent, err := ogent.NewExtension(spec)
	if err != nil {
		log.Fatalf("creating ogent extension: %v", err)
	}
	err = entc.Generate("./schema", &gen.Config{}, entc.Extensions(ogent, oas))
	if err != nil {
		log.Fatalf("running ent codegen: %v", err)
	}
}
```

Note, that the order in which you register the extensions do matter. If you have more extensions enabled, make sure that
the `entoas` extension is executed before `ogent`.

2. Edit the `ent/generate.go` file to execute the `ent/entc.go` instead.

```go
package ent

//go:generate go run -mod=mod entc.go
```

3. Run the code generator.

```shell
go generate ./...
```

You should see a bunch of files generated by the Ent code generator. If you are new to Ent and want to learn more about
them, have a look at their [docs](https://entgo.io/docs).

If you want to follow along with the next part you can have a look at
the [example project](https://github.com/ariga/ogent/tree/master/example/pets) within this repository.

The files we are interested in reside in the `ent/ogent` directory. All the files ending in `_gen.go` are generated
by `ogen`. The file named `oas_server_gen.go` contains the interface generated by `ogen` that a user needs to implement
in order to use `ogen`. `ogent` adds an implementation for that handler in the file `ogent.go`. To see how you can
define what routes to generate and what edges to eager load please head over to
the `entoas` [documentation](https://github.com/ent/contrib/entoas).

`ogent` generates two more files in that directory: `ogent.go` containing the implementation code for the `ogen`
interface and `responses.go` containing helpers to map Ent structs to `ogen` structs.

## Adding custom logic

In order to add HTTP endpoints to your application and document them as well in the OpenAPI Specification document you
can extend the spec `entoas` generates. Additionally `ogen` provides an easy-to-use API to add or edit parts of the
spec. Assume we want to add another route to the document to add a DB health check reachable under the `/db-health`
endpoint. Edit the `ent/entc.go` to include the endpoint in the OpenAPI Specification:

```go
//go:build ignore

package main

import (
	"log"

	"entgo.io/contrib/entoas"
	"entgo.io/ent/entc"
	"entgo.io/ent/entc/gen"
	"ariga.io/ogent"
	"github.com/ogen-go/ogen"
)

func main() {
	spec := new(ogen.Spec)
	oas, err := entoas.NewExtension(
		entoas.Spec(spec),
		entoas.Mutations(func(graph *gen.Graph, spec *ogen.Spec) error {
			spec.AddPathItem("/db-health", ogen.NewPathItem().
				SetDescription("Check the servers DB status").
				SetGet(ogen.NewOperation().
					SetOperationID("DBHealth").
					SetSummary("Ping the database and report").
					AddResponse("204", ogen.NewResponse().SetDescription("DB is reachable")).
					AddResponse("503", ogen.NewResponse().SetDescription("DB is not reachable")),
				),
			)
			return nil
		}),
	)
	if err != nil {
		log.Fatalf("creating entoas extension: %v", err)
	}
	ogent, err := ogent.NewExtension(spec)
	if err != nil {
		log.Fatalf("creating ogent extension: %v", err)
	}
	err = entc.Generate("./schema", &gen.Config{}, entc.Extensions(ogent, oas))
	if err != nil {
		log.Fatalf("running ent codegen: %v", err)
	}
}
```

`ogen` will add another method to the interface and all you have to do is add an implementation for that. The
following `main.go` would do that:

```go
package main

import (
	"context"
	"database/sql"
	"flag"
	"log"
	"net/http"

	"entgo.io/ent/dialect"
	entsql "entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/schema"
	_ "github.com/mattn/go-sqlite3"

	"<your-project>/ent"
	"<your-project>/ent/ogent"
)

type handler struct {
	*ogent.OgentHandler
	db *sql.DB
}

func (h handler) DBHealth(_ context.Context) (ogent.DBHealthRes, error) {
	if err := h.db.Ping(); err != nil {
		return &ogent.DBHealthServiceUnavailable{}, nil
	}
	return &ogent.DBHealthNoContent{}, nil
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
	if err := client.Schema.Create(context.Background(), schema.WithAtlas(true)); err != nil {
		log.Fatal(err)
	}
	// Start listening.
	srv := ogent.NewServer(handler{
		OgentHandler: ogent.NewOgentHandler(client),
		db:           drv.DB(),
	})
	if err := http.ListenAndServe(args.Addr, srv); err != nil {
		log.Fatal(err)
	}
}
```

## Customizing Templates

Since `ogent` is written as an extension to Ent, you can pass in custom templates to customize the generated code by
using the `Templates()` option.

```go
//go:build ignore

package main

import (
	"log"

	"ariga.io/ogent"
	"entgo.io/contrib/entoas"
	"entgo.io/ent/entc"
	"entgo.io/ent/entc/gen"
	"github.com/ogen-go/ogen"
)

var noPagination = gen.MustParse(gen.NewTemplate("").Parse(`
{{ define "ogent/ogent/helper/list/paginate" }}
// Skip pagination
{{ end }}
`))

func main() {
	spec := new(ogen.Spec)
	oas, err := entoas.NewExtension(entoas.Spec(spec))
	if err != nil {
		log.Fatalf("creating entoas extension: %v", err)
	}
	ogent, err := ogent.NewExtension(spec, ogent.Templates(noPagination))
	if err != nil {
		log.Fatalf("creating ogent extension: %v", err)
	}
	err = entc.Generate("./schema", &gen.Config{}, entc.Extensions(ogent, oas))
	if err != nil {
		log.Fatalf("running ent codegen: %v", err)
	}
}
```
