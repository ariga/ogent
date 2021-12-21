package ogent

import (
	"errors"
	"fmt"
	"go/format"
	"os"
	"path/filepath"

	"entgo.io/contrib/entoas"
	"entgo.io/ent/entc"
	"entgo.io/ent/entc/gen"
	"github.com/ogen-go/ogen"
	ogengen "github.com/ogen-go/ogen/gen"
)

type (
	// Config is injected into the code generation templates.
	Config struct {
		// Target holds the filepath to write the ogen assets to.
		Target string
		// The Views created by entoas.
		Views map[string]*entoas.View
	}
	// Extension implements entc.Extension interface providing integration with ogen.
	Extension struct {
		entc.DefaultExtension
		// The OpenAPI Specification to generate handler implementations for.
		spec *ogen.Spec
		// Code generation configuration.
		cfg *Config
	}
	// ExtensionOption allows managing Extension configuration using functional arguments
	ExtensionOption func(*Extension) error
)

// NewExtension returns a new ogent extension with default configuration.
func NewExtension(spec *ogen.Spec, opts ...ExtensionOption) (*Extension, error) {
	if spec == nil {
		return nil, errors.New("ogent: spec cannot be nil")
	}
	ex := &Extension{spec: spec, cfg: new(Config)}
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
		ex.cfg.Target = t
		return nil
	}
}

// Hooks of the extension.
func (ex Extension) Hooks() []gen.Hook {
	return []gen.Hook{
		ex.ogen,
	}
}

// Templates of the extension.
func (Extension) Templates() []*gen.Template {
	return []*gen.Template{Templates}
}

// Annotations of the extension.
func (ex Extension) Annotations() []entc.Annotation {
	return []entc.Annotation{ex.cfg}
}

func (ex Extension) ogen(next gen.Generator) gen.Generator {
	return gen.GenerateFunc(func(g *gen.Graph) error {
		// Ensure target exists.
		t := ex.cfg.Target
		if t == "" {
			t = filepath.Join(g.Target, "ogent")
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
		// Initialize viewsCache.
		ex.cfg.Views, err = entoas.Views(g)
		if err != nil {
			return err
		}
		// Let ent create all of its assets.
		if err := next.Generate(g); err != nil {
			return err
		}
		// Run the ogen code generator.
		generator, err := ogengen.NewGenerator(ex.spec, ogengen.Options{})
		if err != nil {
			return err
		}
		return generator.WriteSource(formatFS{t}, "ogent")
	})
}

// Name implements the entc.Annotation interface.
func (Config) Name() string { return "Ogent" }

type formatFS struct{ Root string }

func (f formatFS) WriteFile(name string, content []byte) error {
	buf, err := format.Source(content)
	if err != nil {
		return err
	}
	return os.WriteFile(filepath.Join(f.Root, name), buf, 0600)
}

var _ entc.Annotation = (*Config)(nil)
