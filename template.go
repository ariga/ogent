package ogent

import (
	"embed"
	"fmt"
	"net/http"
	"path"
	"strings"
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
		"isCreate":       isCreate,
		"isRead":         isRead,
		"isUpdate":       isUpdate,
		"isDelete":       isDelete,
		"isList":         isList,
		"kebab":          strcase.KebabCase,
		"nodeOperations": entoas.NodeOperations,
		"replaceAll":     strings.ReplaceAll,
		"setFieldExpr":   setFieldExpr,
		"viewName":       entoas.ViewName,
		"viewNameEdge":   entoas.ViewNameEdge,
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

// isCreate returns if the given entoas.Operation is entoas.OpCreate.
func isCreate(op entoas.Operation) bool { return op == entoas.OpCreate }

// isRead returns if the given entoas.Operation is entoas.OpRead.
func isRead(op entoas.Operation) bool { return op == entoas.OpRead }

// isUpdate returns if the given entoas.Operation is entoas.OpUpdate.
func isUpdate(op entoas.Operation) bool { return op == entoas.OpUpdate }

// isDelete returns if the given entoas.Operation is entoas.OpDelete.
func isDelete(op entoas.Operation) bool { return op == entoas.OpDelete }

// isList returns if the given entoas.Operation is entoas.OpList.
func isList(op entoas.Operation) bool { return op == entoas.OpList }

// OAS Schema types.
const (
	Integer = "integer"
	Number  = "number"
	String  = "string"
	Boolean = "boolean"
)

// OAS Schema formats.
const (
	None     = ""
	UUID     = "uuid"
	Date     = "date"
	Time     = "time"
	DateTime = "date-time"
	Duration = "duration"
	URI      = "uri"
	IPv4     = "ipv4"
	IPv6     = "ipv6"
	Byte     = "byte"
	Password = "password"
	Int64    = "int64"
	Int32    = "int32"
	Float    = "float"
	Double   = "double"
)

// setFieldExpr returns a Go expression to set the field on a response.
func setFieldExpr(f *gen.Field, ident string) (string, error) {
	if !f.Optional {
		return fmt.Sprintf("%s: %s.%s", f.StructField(), ident, f.StructField()), nil
	}
	t, err := entoas.OgenSchema(f)
	if err != nil {
		return "", err
	}
	buf := new(strings.Builder)
	buf.WriteString(fmt.Sprintf("%s: ", f.StructField()))
	switch t.Type {
	case Integer:
		switch t.Format {
		case Int32:
			buf.WriteString("NewOptInt32(")
		case Int64:
			buf.WriteString("NewOptInt64(")
		case None:
			buf.WriteString("NewOptInt(")
		default:
			return "", fmt.Errorf("unexpected type: %q", t.Format)
		}
	case Number:
		switch t.Format {
		case Float:
			buf.WriteString("NewOptFloat32(")
		case Double, None:
			buf.WriteString("NewOptFloat64(")
		case Int32:
			buf.WriteString("NewOptInt32(")
		case Int64:
			buf.WriteString("NewOptInt64(")
		default:
			return "", fmt.Errorf("unexpected type: %q", t.Format)
		}
	case String:
		switch t.Format {
		case Byte:
			buf.WriteString("NewOptByteSlice(")
		case DateTime, Date, Time:
			buf.WriteString("NewOptTime(")
		case Duration:
			buf.WriteString("NewOptDuration(")
		case UUID:
			buf.WriteString("NewOptUUID(")
		case IPv4, IPv6:
			buf.WriteString("NewOptIP(")
		case URI:
			buf.WriteString("NewOptURL(")
		case Password, None:
			buf.WriteString("NewOptString(")
		default:
			return "", fmt.Errorf("unexpected type: %q", t.Format)
		}
	case Boolean:
		switch t.Format {
		case None:
			buf.WriteString("NewOptBool(")
		default:
			return "", fmt.Errorf("unexpected type: %q", t.Format)
		}
	default:
		return "", fmt.Errorf("unexpected type: %q", t.Format)
	}
	_, _ = fmt.Fprintf(buf, "%s.%s)", ident, f.StructField())
	return buf.String(), nil
}
