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
}

type Builder struct {
	buf strings.Builder
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
	var i int
	pie.Each(fields, func(field Field) {
		b.buf.WriteString(field.String())
		i++
		if i < len(fields) {
			b.buf.WriteRune(',')
		}
	})
	return b
}

func NewBuilder() *Builder {
	return &Builder{}
}

func (b *Builder) Select(fields ...Field) *Builder {
	b.buf.WriteString("SELECT ")
	var i int
	pie.Each(fields, func(field Field) {
		b.buf.WriteString(field.String())
		i++
		if i < len(fields) {
			b.buf.WriteRune(',')
		}
	})
	return b
}

func (b *Builder) From(tables ...Table) *Builder {
	b.buf.WriteString(" FROM ")
	var i int
	pie.Each(tables, func(table Table) {
		b.buf.WriteString(table.String())
		i++
		if i < len(tables) {
			b.buf.WriteRune(',')
		}
	})
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
