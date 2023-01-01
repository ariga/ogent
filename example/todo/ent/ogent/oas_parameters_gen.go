// Code generated by ogen, DO NOT EDIT.

package ogent

import (
	"net/http"
	"net/url"

	"github.com/go-faster/errors"

	"github.com/ogen-go/ogen/conv"
	"github.com/ogen-go/ogen/middleware"
	"github.com/ogen-go/ogen/ogenerrors"
	"github.com/ogen-go/ogen/uri"
	"github.com/ogen-go/ogen/validate"
)

// DeleteTodoParams is parameters of deleteTodo operation.
type DeleteTodoParams struct {
	// ID of the Todo.
	ID int
}

func unpackDeleteTodoParams(packed middleware.Parameters) (params DeleteTodoParams) {
	{
		key := middleware.ParameterKey{
			Name: "id",
			In:   "path",
		}
		params.ID = packed[key].(int)
	}
	return params
}

func decodeDeleteTodoParams(args [1]string, r *http.Request) (params DeleteTodoParams, _ error) {
	// Decode path: id.
	if err := func() error {
		param, err := url.PathUnescape(args[0])
		if err != nil {
			return errors.Wrap(err, "unescape path")
		}
		if len(param) > 0 {
			d := uri.NewPathDecoder(uri.PathDecoderConfig{
				Param:   "id",
				Value:   param,
				Style:   uri.PathStyleSimple,
				Explode: false,
			})

			if err := func() error {
				val, err := d.DecodeValue()
				if err != nil {
					return err
				}

				c, err := conv.ToInt(val)
				if err != nil {
					return err
				}

				params.ID = c
				return nil
			}(); err != nil {
				return err
			}
		} else {
			return validate.ErrFieldRequired
		}
		return nil
	}(); err != nil {
		return params, &ogenerrors.DecodeParamError{
			Name: "id",
			In:   "path",
			Err:  err,
		}
	}
	return params, nil
}

// ListTodoParams is parameters of listTodo operation.
type ListTodoParams struct {
	// What page to render.
	Page OptInt
	// Item count to render per page.
	ItemsPerPage OptInt
}

func unpackListTodoParams(packed middleware.Parameters) (params ListTodoParams) {
	{
		key := middleware.ParameterKey{
			Name: "page",
			In:   "query",
		}
		if v, ok := packed[key]; ok {
			params.Page = v.(OptInt)
		}
	}
	{
		key := middleware.ParameterKey{
			Name: "itemsPerPage",
			In:   "query",
		}
		if v, ok := packed[key]; ok {
			params.ItemsPerPage = v.(OptInt)
		}
	}
	return params
}

func decodeListTodoParams(args [0]string, r *http.Request) (params ListTodoParams, _ error) {
	q := uri.NewQueryDecoder(r.URL.Query())
	// Decode query: page.
	if err := func() error {
		cfg := uri.QueryParameterDecodingConfig{
			Name:    "page",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}

		if err := q.HasParam(cfg); err == nil {
			if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
				var paramsDotPageVal int
				if err := func() error {
					val, err := d.DecodeValue()
					if err != nil {
						return err
					}

					c, err := conv.ToInt(val)
					if err != nil {
						return err
					}

					paramsDotPageVal = c
					return nil
				}(); err != nil {
					return err
				}
				params.Page.SetTo(paramsDotPageVal)
				return nil
			}); err != nil {
				return err
			}
		}
		return nil
	}(); err != nil {
		return params, &ogenerrors.DecodeParamError{
			Name: "page",
			In:   "query",
			Err:  err,
		}
	}
	// Decode query: itemsPerPage.
	if err := func() error {
		cfg := uri.QueryParameterDecodingConfig{
			Name:    "itemsPerPage",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}

		if err := q.HasParam(cfg); err == nil {
			if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
				var paramsDotItemsPerPageVal int
				if err := func() error {
					val, err := d.DecodeValue()
					if err != nil {
						return err
					}

					c, err := conv.ToInt(val)
					if err != nil {
						return err
					}

					paramsDotItemsPerPageVal = c
					return nil
				}(); err != nil {
					return err
				}
				params.ItemsPerPage.SetTo(paramsDotItemsPerPageVal)
				return nil
			}); err != nil {
				return err
			}
		}
		return nil
	}(); err != nil {
		return params, &ogenerrors.DecodeParamError{
			Name: "itemsPerPage",
			In:   "query",
			Err:  err,
		}
	}
	return params, nil
}

// MarkDoneParams is parameters of markDone operation.
type MarkDoneParams struct {
	ID int
}

func unpackMarkDoneParams(packed middleware.Parameters) (params MarkDoneParams) {
	{
		key := middleware.ParameterKey{
			Name: "id",
			In:   "path",
		}
		params.ID = packed[key].(int)
	}
	return params
}

func decodeMarkDoneParams(args [1]string, r *http.Request) (params MarkDoneParams, _ error) {
	// Decode path: id.
	if err := func() error {
		param, err := url.PathUnescape(args[0])
		if err != nil {
			return errors.Wrap(err, "unescape path")
		}
		if len(param) > 0 {
			d := uri.NewPathDecoder(uri.PathDecoderConfig{
				Param:   "id",
				Value:   param,
				Style:   uri.PathStyleSimple,
				Explode: false,
			})

			if err := func() error {
				val, err := d.DecodeValue()
				if err != nil {
					return err
				}

				c, err := conv.ToInt(val)
				if err != nil {
					return err
				}

				params.ID = c
				return nil
			}(); err != nil {
				return err
			}
		} else {
			return validate.ErrFieldRequired
		}
		return nil
	}(); err != nil {
		return params, &ogenerrors.DecodeParamError{
			Name: "id",
			In:   "path",
			Err:  err,
		}
	}
	return params, nil
}

// ReadTodoParams is parameters of readTodo operation.
type ReadTodoParams struct {
	// ID of the Todo.
	ID int
}

func unpackReadTodoParams(packed middleware.Parameters) (params ReadTodoParams) {
	{
		key := middleware.ParameterKey{
			Name: "id",
			In:   "path",
		}
		params.ID = packed[key].(int)
	}
	return params
}

func decodeReadTodoParams(args [1]string, r *http.Request) (params ReadTodoParams, _ error) {
	// Decode path: id.
	if err := func() error {
		param, err := url.PathUnescape(args[0])
		if err != nil {
			return errors.Wrap(err, "unescape path")
		}
		if len(param) > 0 {
			d := uri.NewPathDecoder(uri.PathDecoderConfig{
				Param:   "id",
				Value:   param,
				Style:   uri.PathStyleSimple,
				Explode: false,
			})

			if err := func() error {
				val, err := d.DecodeValue()
				if err != nil {
					return err
				}

				c, err := conv.ToInt(val)
				if err != nil {
					return err
				}

				params.ID = c
				return nil
			}(); err != nil {
				return err
			}
		} else {
			return validate.ErrFieldRequired
		}
		return nil
	}(); err != nil {
		return params, &ogenerrors.DecodeParamError{
			Name: "id",
			In:   "path",
			Err:  err,
		}
	}
	return params, nil
}

// UpdateTodoParams is parameters of updateTodo operation.
type UpdateTodoParams struct {
	// ID of the Todo.
	ID int
}

func unpackUpdateTodoParams(packed middleware.Parameters) (params UpdateTodoParams) {
	{
		key := middleware.ParameterKey{
			Name: "id",
			In:   "path",
		}
		params.ID = packed[key].(int)
	}
	return params
}

func decodeUpdateTodoParams(args [1]string, r *http.Request) (params UpdateTodoParams, _ error) {
	// Decode path: id.
	if err := func() error {
		param, err := url.PathUnescape(args[0])
		if err != nil {
			return errors.Wrap(err, "unescape path")
		}
		if len(param) > 0 {
			d := uri.NewPathDecoder(uri.PathDecoderConfig{
				Param:   "id",
				Value:   param,
				Style:   uri.PathStyleSimple,
				Explode: false,
			})

			if err := func() error {
				val, err := d.DecodeValue()
				if err != nil {
					return err
				}

				c, err := conv.ToInt(val)
				if err != nil {
					return err
				}

				params.ID = c
				return nil
			}(); err != nil {
				return err
			}
		} else {
			return validate.ErrFieldRequired
		}
		return nil
	}(); err != nil {
		return params, &ogenerrors.DecodeParamError{
			Name: "id",
			In:   "path",
			Err:  err,
		}
	}
	return params, nil
}
