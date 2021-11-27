package ogent

import (
	"errors"
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
func NewExtension(src io.Reader, opts ...ExtensionOption) (*Extension, error) {
	if src == nil {
		return nil, errors.New("ogent: src cannot be nil")
	}
	ex := &Extension{src: src}
	for _, opt := range opts {
		if err := opt(ex); err != nil {
			return nil, err
		}
	}
	return ex, nil
}
