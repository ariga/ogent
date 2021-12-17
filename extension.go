package ogent

import (
	"errors"
	"fmt"
	"go/format"
	"os"
	"path/filepath"

	"entgo.io/ent/entc"
	"entgo.io/ent/entc/gen"
	"github.com/ogen-go/ogen"
	ogengen "github.com/ogen-go/ogen/gen"
)

type (
	// Extension implements entc.Extension interface providing integration with ogen.
	Extension struct {
		entc.DefaultExtension
		// The OpenAPI Specification to generate handler implementations for.
		spec *ogen.Spec
		// target holds the filepath to write the ogen assets to.
		target string
		// pkg holds the package name of the ogen assets.
		pkg string
	}
	// ExtensionOption allows managing Extension configuration using functional arguments
	ExtensionOption func(*Extension) error
)

// NewExtension returns a new ogent extension with default configuration.
func NewExtension(spec *ogen.Spec, opts ...ExtensionOption) (*Extension, error) {
	if spec == nil {
		return nil, errors.New("ogent: spec cannot be nil")
	}
	ex := &Extension{spec: spec}
	for _, opt := range opts {
		if err := opt(ex); err != nil {
			return nil, err
		}
	}
	return ex, nil
}

// Target sets the directory the ogen assets are written to.
func Target(t string) ExtensionOption {
	return func(ex *Extension) error {
		ex.target = t
		return nil
	}
}

// Package sets the package name for the ogen assets.
func Package(pkg string) ExtensionOption {
	return func(ex *Extension) error {
		ex.pkg = pkg
		return nil
	}
}

func (ex Extension) Hooks() []gen.Hook {
	return []gen.Hook{
		ex.ogen,
	}
}

func (ex Extension) ogen(next gen.Generator) gen.Generator {
	return gen.GenerateFunc(func(g *gen.Graph) error {
		// Let ent create all of its assets.
		if err := next.Generate(g); err != nil {
			return err
		}
		// Ensure target exists.
		t := ex.target
		if t == "" {
			t = filepath.Join(g.Target, "api")
		}
		_, err := os.Stat(t)
		if err != nil && !os.IsNotExist(err) {
			return fmt.Errorf("ogent: read target dir: %w", err)
		}
		if os.IsNotExist(err) {
			if err := os.MkdirAll(t, 0750); err != nil {
				return fmt.Errorf("ogent: create target dir: %w", err)
			}
		}
		// Ensure there is a package name given.
		pkg := ex.pkg
		if pkg == "" {
			pkg = "api"
		}
		// Run the ogen code generator.
		generator, err := ogengen.NewGenerator(ex.spec, ogengen.Options{})
		if err != nil {
			return err
		}
		return generator.WriteSource(formatFS{t}, pkg)
	})
}

type formatFS struct{ Root string }

func (f formatFS) WriteFile(name string, content []byte) error {
	buf, err := format.Source(content)
	if err != nil {
		return err
	}
	return os.WriteFile(filepath.Join(f.Root, name), buf, 0600)
}
