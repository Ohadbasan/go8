// Code generated by ent, DO NOT EDIT.

package gen

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/gmhafiz/go8/ent/gen/author"
)

// Author is the model entity for the Author schema.
type Author struct {
	config `json:"-"`
	// ID of the ent.
	ID uint `json:"id,omitempty"`
	// FirstName holds the value of the "first_name" field.
	FirstName string `json:"first_name,omitempty"`
	// MiddleName holds the value of the "middle_name" field.
	MiddleName string `json:"middle_name,omitempty"`
	// LastName holds the value of the "last_name" field.
	LastName string `json:"last_name,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"-"`
	// UpdatedAt holds the value of the "updated_at" field.
	UpdatedAt time.Time `json:"-"`
	// DeletedAt holds the value of the "deleted_at" field.
	DeletedAt *time.Time `json:"-"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the AuthorQuery when eager-loading is set.
	Edges        AuthorEdges `json:"edges"`
	selectValues sql.SelectValues
}

// AuthorEdges holds the relations/edges for other nodes in the graph.
type AuthorEdges struct {
	// Books holds the value of the books edge.
	Books []*Book `json:"books,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
}

// BooksOrErr returns the Books value or an error if the edge
// was not loaded in eager-loading.
func (e AuthorEdges) BooksOrErr() ([]*Book, error) {
	if e.loadedTypes[0] {
		return e.Books, nil
	}
	return nil, &NotLoadedError{edge: "books"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Author) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case author.FieldID:
			values[i] = new(sql.NullInt64)
		case author.FieldFirstName, author.FieldMiddleName, author.FieldLastName:
			values[i] = new(sql.NullString)
		case author.FieldCreatedAt, author.FieldUpdatedAt, author.FieldDeletedAt:
			values[i] = new(sql.NullTime)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Author fields.
func (a *Author) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case author.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			a.ID = uint(value.Int64)
		case author.FieldFirstName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field first_name", values[i])
			} else if value.Valid {
				a.FirstName = value.String
			}
		case author.FieldMiddleName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field middle_name", values[i])
			} else if value.Valid {
				a.MiddleName = value.String
			}
		case author.FieldLastName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field last_name", values[i])
			} else if value.Valid {
				a.LastName = value.String
			}
		case author.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				a.CreatedAt = value.Time
			}
		case author.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				a.UpdatedAt = value.Time
			}
		case author.FieldDeletedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field deleted_at", values[i])
			} else if value.Valid {
				a.DeletedAt = new(time.Time)
				*a.DeletedAt = value.Time
			}
		default:
			a.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the Author.
// This includes values selected through modifiers, order, etc.
func (a *Author) Value(name string) (ent.Value, error) {
	return a.selectValues.Get(name)
}

// QueryBooks queries the "books" edge of the Author entity.
func (a *Author) QueryBooks() *BookQuery {
	return NewAuthorClient(a.config).QueryBooks(a)
}

// Update returns a builder for updating this Author.
// Note that you need to call Author.Unwrap() before calling this method if this Author
// was returned from a transaction, and the transaction was committed or rolled back.
func (a *Author) Update() *AuthorUpdateOne {
	return NewAuthorClient(a.config).UpdateOne(a)
}

// Unwrap unwraps the Author entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (a *Author) Unwrap() *Author {
	_tx, ok := a.config.driver.(*txDriver)
	if !ok {
		panic("gen: Author is not a transactional entity")
	}
	a.config.driver = _tx.drv
	return a
}

// String implements the fmt.Stringer.
func (a *Author) String() string {
	var builder strings.Builder
	builder.WriteString("Author(")
	builder.WriteString(fmt.Sprintf("id=%v, ", a.ID))
	builder.WriteString("first_name=")
	builder.WriteString(a.FirstName)
	builder.WriteString(", ")
	builder.WriteString("middle_name=")
	builder.WriteString(a.MiddleName)
	builder.WriteString(", ")
	builder.WriteString("last_name=")
	builder.WriteString(a.LastName)
	builder.WriteString(", ")
	builder.WriteString("created_at=")
	builder.WriteString(a.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(a.UpdatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	if v := a.DeletedAt; v != nil {
		builder.WriteString("deleted_at=")
		builder.WriteString(v.Format(time.ANSIC))
	}
	builder.WriteByte(')')
	return builder.String()
}

// Authors is a parsable slice of Author.
type Authors []*Author
