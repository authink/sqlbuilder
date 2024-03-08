package sqlbuilder

import (
	"fmt"
	"strings"

	"github.com/elliotchance/pie/v2"
)

type Keyword interface {
	fmt.Stringer
	Select(...Field) *Builder
	From(...Table) *Builder
	Where(Condition) *Builder
	And(Condition) *Builder
	Or(Condition) *Builder
	ForUpdate() *Builder
	OrderBy(...Field) *Builder
	Asc() *Builder
	Desc() *Builder
	Limit() *Builder
	InsertInto(Table) *Builder
	Columns(...Field) *Builder
	Update(Table) *Builder
	Set(...Field) *Builder
	DeleteFrom(Table) *Builder
}

type Builder struct {
	buf strings.Builder
}

// Limit implements Keyword.
func (b *Builder) Limit() *Builder {
	b.buf.WriteString(" LIMIT :limit OFFSET :offset")
	return b
}

// DeleteFrom implements Keyword.
func (b *Builder) DeleteFrom(table Table) *Builder {
	b.buf.WriteString("DELETE FROM ")
	b.buf.WriteString(string(table))
	return b
}

// Set implements Keyword.
func (b *Builder) Set(fields ...Field) *Builder {
	var assignFields []Field
	pie.Each(fields, func(field Field) {
		assignFields = append(assignFields, field.Assign())
	})

	b.buf.WriteString(" SET ")
	writeFields(&b.buf, assignFields)
	return b
}

// Update implements Keyword.
func (b *Builder) Update(table Table) *Builder {
	b.buf.WriteString("UPDATE ")
	b.buf.WriteString(string(table))
	return b
}

// Columns implements Keyword.
func (b *Builder) Columns(fields ...Field) *Builder {
	b.buf.WriteString("(")
	writeFields(&b.buf, fields)
	b.buf.WriteRune(')')

	var namedFields []Field
	pie.Each(fields, func(field Field) {
		namedFields = append(namedFields, field.Named())
	})
	b.buf.WriteString(" VALUES(")
	writeFields(&b.buf, namedFields)
	b.buf.WriteRune(')')
	return b
}

// InsertInto implements Keyword.
func (b *Builder) InsertInto(table Table) *Builder {
	b.buf.WriteString("INSERT INTO ")
	b.buf.WriteString(string(table))
	return b
}

// ForUpdate implements Keyword.
func (b *Builder) ForUpdate() *Builder {
	b.buf.WriteString(" FOR UPDATE")
	return b
}

// Asc implements Keyword.
func (b *Builder) Asc() *Builder {
	b.buf.WriteString(" ASC")
	return b
}

// Desc implements Keyword.
func (b *Builder) Desc() *Builder {
	b.buf.WriteString(" DESC")
	return b
}

// Or implements Keyword.
func (b *Builder) Or(cond Condition) *Builder {
	b.buf.WriteString(" OR ")
	b.buf.WriteString(cond.String())
	return b
}

// OrderBy implements Keyword.
func (b *Builder) OrderBy(fields ...Field) *Builder {
	b.buf.WriteString(" ORDER BY ")
	writeFields(&b.buf, fields)
	return b
}

func NewBuilder() *Builder {
	return &Builder{}
}

func (b *Builder) Select(fields ...Field) *Builder {
	b.buf.WriteString("SELECT ")
	writeFields(&b.buf, fields)
	return b
}

func (b *Builder) From(tables ...Table) *Builder {
	b.buf.WriteString(" FROM ")
	writeTables(&b.buf, tables)
	return b
}

func (b *Builder) Where(cond Condition) *Builder {
	b.buf.WriteString(" WHERE ")
	b.buf.WriteString(cond.String())
	return b
}

// And implements Keyword.
func (b *Builder) And(cond Condition) *Builder {
	b.buf.WriteString(" AND ")
	b.buf.WriteString(cond.String())
	return b
}

// String implements Keyword.
func (b *Builder) String() string {
	return b.buf.String()
}

var _ Keyword = (*Builder)(nil)
