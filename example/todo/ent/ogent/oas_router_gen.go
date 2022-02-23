// Code generated by ogen, DO NOT EDIT.

package ogent

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"math"
	"math/bits"
	"net"
	"net/http"
	"net/url"
	"regexp"
	"sort"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/go-faster/errors"
	"github.com/go-faster/jx"
	"github.com/google/uuid"
	"github.com/ogen-go/ogen/conv"
	ht "github.com/ogen-go/ogen/http"
	"github.com/ogen-go/ogen/json"
	"github.com/ogen-go/ogen/otelogen"
	"github.com/ogen-go/ogen/uri"
	"github.com/ogen-go/ogen/validate"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/codes"
	"go.opentelemetry.io/otel/metric"
	"go.opentelemetry.io/otel/trace"
)

// No-op definition for keeping imports.
var (
	_ = context.Background()
	_ = fmt.Stringer(nil)
	_ = strings.Builder{}
	_ = errors.Is
	_ = sort.Ints
	_ = http.MethodGet
	_ = io.Copy
	_ = json.Marshal
	_ = bytes.NewReader
	_ = strconv.ParseInt
	_ = time.Time{}
	_ = conv.ToInt32
	_ = uuid.UUID{}
	_ = uri.PathEncoder{}
	_ = url.URL{}
	_ = math.Mod
	_ = bits.LeadingZeros64
	_ = validate.Int{}
	_ = ht.NewRequest
	_ = net.IP{}
	_ = otelogen.Version
	_ = trace.TraceIDFromHex
	_ = otel.GetTracerProvider
	_ = metric.NewNoopMeterProvider
	_ = regexp.MustCompile
	_ = jx.Null
	_ = sync.Pool{}
	_ = codes.Unset
)

func (s *Server) notFound(w http.ResponseWriter, r *http.Request) {
	http.NotFound(w, r)
}

// ServeHTTP serves http request as defined by OpenAPI v3 specification,
// calling handler that matches the path or returning not found error.
func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	elem := r.URL.Path
	if len(elem) == 0 {
		s.notFound(w, r)
		return
	}
	args := [1]string{}
	// Static code generated router with unwrapped path search.
	switch r.Method {
	case "DELETE":
		if len(elem) == 0 {
			break
		}
		switch elem[0] {
		case '/': // Prefix: "/todos/"
			if l := len("/todos/"); len(elem) >= l && elem[0:l] == "/todos/" {
				elem = elem[l:]
			} else {
				break
			}

			// Param: "id"
			// Leaf parameter
			args[0] = elem
			elem = ""

			if len(elem) == 0 {
				// Leaf: DeleteTodo
				s.handleDeleteTodoRequest([1]string{
					args[0],
				}, w, r)

				return
			}
		}
	case "GET":
		if len(elem) == 0 {
			break
		}
		switch elem[0] {
		case '/': // Prefix: "/todos"
			if l := len("/todos"); len(elem) >= l && elem[0:l] == "/todos" {
				elem = elem[l:]
			} else {
				break
			}

			if len(elem) == 0 {
				s.handleListTodoRequest([0]string{}, w, r)

				return
			}
			switch elem[0] {
			case '/': // Prefix: "/"
				if l := len("/"); len(elem) >= l && elem[0:l] == "/" {
					elem = elem[l:]
				} else {
					break
				}

				// Param: "id"
				// Leaf parameter
				args[0] = elem
				elem = ""

				if len(elem) == 0 {
					// Leaf: ReadTodo
					s.handleReadTodoRequest([1]string{
						args[0],
					}, w, r)

					return
				}
			}
		}
	case "PATCH":
		if len(elem) == 0 {
			break
		}
		switch elem[0] {
		case '/': // Prefix: "/todos/"
			if l := len("/todos/"); len(elem) >= l && elem[0:l] == "/todos/" {
				elem = elem[l:]
			} else {
				break
			}

			// Param: "id"
			// Match until "/"
			idx := strings.IndexByte(elem, '/')
			if idx < 0 {
				idx = len(elem)
			}
			args[0] = elem[:idx]
			elem = elem[idx:]

			if len(elem) == 0 {
				s.handleUpdateTodoRequest([1]string{
					args[0],
				}, w, r)

				return
			}
			switch elem[0] {
			case '/': // Prefix: "/done"
				if l := len("/done"); len(elem) >= l && elem[0:l] == "/done" {
					elem = elem[l:]
				} else {
					break
				}

				if len(elem) == 0 {
					// Leaf: MarkDone
					s.handleMarkDoneRequest([1]string{
						args[0],
					}, w, r)

					return
				}
			}
		}
	case "POST":
		if len(elem) == 0 {
			break
		}
		switch elem[0] {
		case '/': // Prefix: "/todos"
			if l := len("/todos"); len(elem) >= l && elem[0:l] == "/todos" {
				elem = elem[l:]
			} else {
				break
			}

			if len(elem) == 0 {
				// Leaf: CreateTodo
				s.handleCreateTodoRequest([0]string{}, w, r)

				return
			}
		}
	}
	s.notFound(w, r)
}

// Route is route object.
type Route struct {
	name  string
	count int
	args  [1]string
}

// OperationID returns OpenAPI operationId.
func (r Route) OperationID() string {
	return r.name
}

// Args returns parsed arguments.
func (r Route) Args() []string {
	return r.args[:r.count]
}

// FindRoute finds Route for given method and path.
func (s *Server) FindRoute(method, path string) (r Route, _ bool) {
	var (
		args = [1]string{}
		elem = path
	)
	r.args = args

	// Static code generated router with unwrapped path search.
	switch method {
	case "DELETE":
		if len(elem) == 0 {
			break
		}
		switch elem[0] {
		case '/': // Prefix: "/todos/"
			if l := len("/todos/"); len(elem) >= l && elem[0:l] == "/todos/" {
				elem = elem[l:]
			} else {
				break
			}

			// Param: "id"
			// Leaf parameter
			args[0] = elem
			elem = ""

			if len(elem) == 0 {
				// Leaf: DeleteTodo
				r.name = "DeleteTodo"
				r.args = args
				r.count = 1
				return r, true
			}
		}
	case "GET":
		if len(elem) == 0 {
			break
		}
		switch elem[0] {
		case '/': // Prefix: "/todos"
			if l := len("/todos"); len(elem) >= l && elem[0:l] == "/todos" {
				elem = elem[l:]
			} else {
				break
			}

			if len(elem) == 0 {
				r.name = "ListTodo"
				r.args = args
				r.count = 0
				return r, true
			}
			switch elem[0] {
			case '/': // Prefix: "/"
				if l := len("/"); len(elem) >= l && elem[0:l] == "/" {
					elem = elem[l:]
				} else {
					break
				}

				// Param: "id"
				// Leaf parameter
				args[0] = elem
				elem = ""

				if len(elem) == 0 {
					// Leaf: ReadTodo
					r.name = "ReadTodo"
					r.args = args
					r.count = 1
					return r, true
				}
			}
		}
	case "PATCH":
		if len(elem) == 0 {
			break
		}
		switch elem[0] {
		case '/': // Prefix: "/todos/"
			if l := len("/todos/"); len(elem) >= l && elem[0:l] == "/todos/" {
				elem = elem[l:]
			} else {
				break
			}

			// Param: "id"
			// Match until "/"
			idx := strings.IndexByte(elem, '/')
			if idx < 0 {
				idx = len(elem)
			}
			args[0] = elem[:idx]
			elem = elem[idx:]

			if len(elem) == 0 {
				r.name = "UpdateTodo"
				r.args = args
				r.count = 1
				return r, true
			}
			switch elem[0] {
			case '/': // Prefix: "/done"
				if l := len("/done"); len(elem) >= l && elem[0:l] == "/done" {
					elem = elem[l:]
				} else {
					break
				}

				if len(elem) == 0 {
					// Leaf: MarkDone
					r.name = "MarkDone"
					r.args = args
					r.count = 1
					return r, true
				}
			}
		}
	case "POST":
		if len(elem) == 0 {
			break
		}
		switch elem[0] {
		case '/': // Prefix: "/todos"
			if l := len("/todos"); len(elem) >= l && elem[0:l] == "/todos" {
				elem = elem[l:]
			} else {
				break
			}

			if len(elem) == 0 {
				// Leaf: CreateTodo
				r.name = "CreateTodo"
				r.args = args
				r.count = 0
				return r, true
			}
		}
	}
	return r, false
}