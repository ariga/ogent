package ogent

import (
	"embed"
	"fmt"
	"net/http"
	"path"
	"text/template"

	"entgo.io/contrib/entoas"
	"entgo.io/ent/entc/gen"
	"github.com/stoewer/go-strcase"
)

var (
	//go:embed template
	templateDir embed.FS
	// FuncMap contains extra template functions used by ogent.
	FuncMap = template.FuncMap{
		"edgeOperations": entoas.EdgeOperations,
		"hasParams":      hasParams,
		"hasRequestBody": hasRequestBody,
		"httpVerb":       httpVerb,
		"httpRoute":      httpRoute,
		"kebab":          strcase.KebabCase,
		"nodeOperations": entoas.NodeOperations,
	}
	// Templates holds all templates used by ogent.
	Templates = gen.MustParse(gen.NewTemplate("ogent").Funcs(FuncMap).ParseFS(templateDir, "template/*tmpl"))
)

// hasParams returns if the given entoas.Operation has parameters.
func hasParams(op entoas.Operation) bool {
	return op != entoas.OpCreate
}

// hasRequestBody returns if the given entoas.Operation has a request body.
func hasRequestBody(op entoas.Operation) bool {
	return op == entoas.OpCreate || op == entoas.OpUpdate
}

// httpVerb returns the HTTP httpVerb for the given entoas.Operation.
func httpVerb(op entoas.Operation) (string, error) {
	switch op {
	case entoas.OpCreate:
		return http.MethodPost, nil
	case entoas.OpRead, entoas.OpList:
		return http.MethodGet, nil
	case entoas.OpUpdate:
		return http.MethodPatch, nil
	case entoas.OpDelete:
		return http.MethodDelete, nil
	}
	return "", fmt.Errorf("unknown operation: %q", op)
}

// httpRoute returns the HTTP endpoint for the given entoas.Operation.
func httpRoute(root string, op entoas.Operation) (string, error) {
	switch op {
	case entoas.OpCreate, entoas.OpList:
		return root, nil
	case entoas.OpRead, entoas.OpUpdate, entoas.OpDelete:
		return path.Join(root, "{id}"), nil
	}
	return "", fmt.Errorf("unknown operation: %q", op)
}
