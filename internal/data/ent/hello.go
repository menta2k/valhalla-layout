// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"entgo.io/ent/dialect/sql"
	"github.com/menta2l/valhalla-layout/internal/data/ent/hello"
)

// Hello is the model entity for the Hello schema.
type Hello struct {
	config
	// ID of the ent.
	ID int `json:"id,omitempty"`
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Hello) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case hello.FieldID:
			values[i] = new(sql.NullInt64)
		default:
			return nil, fmt.Errorf("unexpected column %q for type Hello", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Hello fields.
func (h *Hello) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case hello.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			h.ID = int(value.Int64)
		}
	}
	return nil
}

// Update returns a builder for updating this Hello.
// Note that you need to call Hello.Unwrap() before calling this method if this Hello
// was returned from a transaction, and the transaction was committed or rolled back.
func (h *Hello) Update() *HelloUpdateOne {
	return (&HelloClient{config: h.config}).UpdateOne(h)
}

// Unwrap unwraps the Hello entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (h *Hello) Unwrap() *Hello {
	tx, ok := h.config.driver.(*txDriver)
	if !ok {
		panic("ent: Hello is not a transactional entity")
	}
	h.config.driver = tx.drv
	return h
}

// String implements the fmt.Stringer.
func (h *Hello) String() string {
	var builder strings.Builder
	builder.WriteString("Hello(")
	builder.WriteString(fmt.Sprintf("id=%v", h.ID))
	builder.WriteByte(')')
	return builder.String()
}

// Hellos is a parsable slice of Hello.
type Hellos []*Hello

func (h Hellos) config(cfg config) {
	for _i := range h {
		h[_i].config = cfg
	}
}
