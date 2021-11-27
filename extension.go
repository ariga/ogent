package ogent

import (
	"io"

	"entgo.io/ent/entc"
)

type (
	// Extension implements entc.Extension interface providing integration with ogen.
	Extension struct {
		entc.DefaultExtension
		src io.Reader
	}
	// ExtensionOption allows managing Extension configuration using functional arguments
	ExtensionOption func(*Extension) error
)

// NewExtension returns a new ogent extension with default configuration.
func NewExtension(opts ...ExtensionOption) (*Extension, error) {
	ex := &Extension{}
	for _, opt := range opts {
		if err := opt(ex); err != nil {
			return nil, err
		}
	}
	return ex, nil
}

// ReadFrom attempts to read an OpenAPI Specification document from the given reader.
func ReadFrom(src io.Reader) ExtensionOption {
	return func(ex *Extension) error {
		ex.src = src
		return nil
	}
}
