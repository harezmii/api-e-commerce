// Code generated by entc, DO NOT EDIT.

package ent

import (
	"api/ent/faq"
	"fmt"
	"strings"
	"time"

	"entgo.io/ent/dialect/sql"
)

// Faq is the model entity for the Faq schema.
type Faq struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// Question holds the value of the "question" field.
	Question string `json:"question,omitempty"`
	// Answer holds the value of the "answer" field.
	Answer string `json:"answer,omitempty"`
	// Status holds the value of the "status" field.
	Status bool `json:"status,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"created_at,omitempty"`
	// UpdatedAt holds the value of the "updated_at" field.
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	// DeletedAt holds the value of the "deleted_at" field.
	DeletedAt time.Time `json:"deleted_at,omitempty"`
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Faq) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case faq.FieldStatus:
			values[i] = new(sql.NullBool)
		case faq.FieldID:
			values[i] = new(sql.NullInt64)
		case faq.FieldQuestion, faq.FieldAnswer:
			values[i] = new(sql.NullString)
		case faq.FieldCreatedAt, faq.FieldUpdatedAt, faq.FieldDeletedAt:
			values[i] = new(sql.NullTime)
		default:
			return nil, fmt.Errorf("unexpected column %q for type Faq", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Faq fields.
func (f *Faq) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case faq.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			f.ID = int(value.Int64)
		case faq.FieldQuestion:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field question", values[i])
			} else if value.Valid {
				f.Question = value.String
			}
		case faq.FieldAnswer:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field answer", values[i])
			} else if value.Valid {
				f.Answer = value.String
			}
		case faq.FieldStatus:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field status", values[i])
			} else if value.Valid {
				f.Status = value.Bool
			}
		case faq.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				f.CreatedAt = value.Time
			}
		case faq.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				f.UpdatedAt = value.Time
			}
		case faq.FieldDeletedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field deleted_at", values[i])
			} else if value.Valid {
				f.DeletedAt = value.Time
			}
		}
	}
	return nil
}

// Update returns a builder for updating this Faq.
// Note that you need to call Faq.Unwrap() before calling this method if this Faq
// was returned from a transaction, and the transaction was committed or rolled back.
func (f *Faq) Update() *FaqUpdateOne {
	return (&FaqClient{config: f.config}).UpdateOne(f)
}

// Unwrap unwraps the Faq entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (f *Faq) Unwrap() *Faq {
	tx, ok := f.config.driver.(*txDriver)
	if !ok {
		panic("ent: Faq is not a transactional entity")
	}
	f.config.driver = tx.drv
	return f
}

// String implements the fmt.Stringer.
func (f *Faq) String() string {
	var builder strings.Builder
	builder.WriteString("Faq(")
	builder.WriteString(fmt.Sprintf("id=%v", f.ID))
	builder.WriteString(", question=")
	builder.WriteString(f.Question)
	builder.WriteString(", answer=")
	builder.WriteString(f.Answer)
	builder.WriteString(", status=")
	builder.WriteString(fmt.Sprintf("%v", f.Status))
	builder.WriteString(", created_at=")
	builder.WriteString(f.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", updated_at=")
	builder.WriteString(f.UpdatedAt.Format(time.ANSIC))
	builder.WriteString(", deleted_at=")
	builder.WriteString(f.DeletedAt.Format(time.ANSIC))
	builder.WriteByte(')')
	return builder.String()
}

// Faqs is a parsable slice of Faq.
type Faqs []*Faq

func (f Faqs) config(cfg config) {
	for _i := range f {
		f[_i].config = cfg
	}
}
