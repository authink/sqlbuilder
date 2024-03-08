package sqlbuilder

import "fmt"

type TableModifier interface {
	fmt.Stringer
	As(Table) Table
	Field(string) Field
}

type Table string

// As implements Modifier.
func (t Table) As(alias Table) Table {
	return Table(fmt.Sprintf("%s AS %s", t, alias))
}

// String implements Modifier.
func (t Table) String() string {
	return string(t)
}

// Field implements TableModifier.
func (t Table) Field(name string) Field {
	return Field(fmt.Sprintf("%s.%s", t, name))
}

var _ TableModifier = Table("")
