module github.com/ariga/ogent

go 1.17

require (
	entgo.io/ent v0.9.2-0.20211216115003-1c263c7abd10
	github.com/go-faster/errors v0.5.0
	github.com/go-faster/jx v0.25.0
	github.com/google/uuid v1.3.0
	github.com/ogen-go/ogen v0.0.0-00010101000000-000000000000
	go.opentelemetry.io/otel v1.3.0
	go.opentelemetry.io/otel/metric v0.26.0
	go.opentelemetry.io/otel/trace v1.3.0
)

require golang.org/x/sys v0.0.0-20211113001501-0c823b97ae02 // indirect

require (
	entgo.io/contrib v0.0.0-00010101000000-000000000000 // indirect
	github.com/fatih/color v1.13.0 // indirect
	github.com/go-logr/logr v1.2.1 // indirect
	github.com/go-logr/stdr v1.2.0 // indirect
	github.com/go-openapi/inflect v0.19.0 // indirect
	github.com/goccy/go-yaml v1.9.4 // indirect
	github.com/mattn/go-colorable v0.1.12 // indirect
	github.com/mattn/go-isatty v0.0.14 // indirect
	golang.org/x/crypto v0.0.0-20210513164829-c07d793c2f9a // indirect
	golang.org/x/mod v0.5.1 // indirect
	golang.org/x/tools v0.1.8 // indirect
	golang.org/x/xerrors v0.0.0-20200804184101-5ec99f83aff1 // indirect
)

replace github.com/ogen-go/ogen => github.com/masseelch/ogen v0.0.0-20211217142721-bb02574f0aab

replace entgo.io/contrib => entgo.io/contrib v0.2.1-0.20211217200019-533305064701
