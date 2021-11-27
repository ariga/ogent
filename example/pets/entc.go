//go:build ignore
// +build ignore

package main

import (
	"io"
	"log"

	"entgo.io/contrib/entoas"
	"entgo.io/ent/entc"
	"entgo.io/ent/entc/gen"
	"github.com/ariga/ogent"
)

func main() {
	r, w := io.Pipe()
	oas, err := entoas.NewExtension(entoas.WriteTo(w))
	if err != nil {
		log.Fatalf("creating entoas extension: %v", err)
	}
	ogen, err := ogent.NewExtension(ogent.ReadFrom(r))
	if err != nil {
		log.Fatalf("creating ogent extension: %v", err)
	}
	err = entc.Generate("./schema", &gen.Config{}, entc.Extensions(oas))
	if err != nil {
		log.Fatalf("running ent codegen: %v", err)
	}
}
