package schema

import (
	"database/sql/driver"
	"fmt"
	"net/url"

	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// AllTypes holds the schema definition for the AllTypes entity.
type AllTypes struct {
	ent.Schema
}

// Fields of the AllTypes.
func (AllTypes) Fields() []ent.Field {
	return []ent.Field{
		field.Int("int"),
		field.Int8("int8"),
		field.Int16("int16"),
		field.Int32("int32"),
		field.Int64("int64"),
		field.Uint("uint"),
		field.Uint8("uint8"),
		field.Uint16("uint16"),
		field.Uint32("uint32"),
		field.Uint64("uint64"),
		field.Float32("float32"),
		field.Float("float64"),
		field.String("string_type"),
		field.Bool("bool"),
		field.UUID("uuid", uuid.Nil),
		field.Time("time"),
		field.Text("text"),
		field.Enum("state").
			Values("on", "off"),
		// field.Strings("strings"),
		// field.Ints("ints"),
		// field.Floats("floats"),
		field.Bytes("bytes"),
		// field.JSON("nicknames", []string{}),
		// field.JSON("json_slice", []http.Dir{}).
		// 	Annotations(entoas.Schema(ogen.String().AsArray())),
		// field.JSON("json_obj", url.URL{}).
		// 	Annotations(entoas.Schema(ogen.String())),
		// field.Other("other", &Link{}).
		// 	SchemaType(map[string]string{dialect.Postgres: "varchar"}).
		// 	Default(DefaultLink()).
		// 	Annotations(entoas.Schema(ogen.String())),
	}
}

type Link struct{ *url.URL }

func DefaultLink() *Link {
	u, _ := url.Parse("127.0.0.1")
	return &Link{URL: u}
}

// Scan implements the Scanner interface.
func (l *Link) Scan(value interface{}) (err error) {
	switch v := value.(type) {
	case nil:
	case []byte:
		l.URL, err = url.Parse(string(v))
	case string:
		l.URL, err = url.Parse(v)
	default:
		err = fmt.Errorf("unexpected type %T", v)
	}
	return
}

// Value implements the driver Valuer interface.
func (l Link) Value() (driver.Value, error) {
	if l.URL == nil {
		return nil, nil
	}
	return l.String(), nil
}
