package sqlbuilder

import (
	"strings"

	"github.com/elliotchance/pie/v2"
)

func writeFields(buf *strings.Builder, fields []Field) {
	var i int
	pie.Each(fields, func(field Field) {
		buf.WriteString(string(field))
		i++
		if i < len(fields) {
			buf.WriteRune(',')
		}
	})
}

func writeTables(buf *strings.Builder, tables []Table) {
	var i int
	pie.Each(tables, func(table Table) {
		buf.WriteString(string(table))
		i++
		if i < len(tables) {
			buf.WriteRune(',')
		}
	})
}
