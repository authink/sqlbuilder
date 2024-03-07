package sqlbuilder

import "fmt"

type TableModifier interface {
	fmt.Stringer
	As(string) Table
}

type Table string

// As implements Modifier.
func (f Table) As(alias string) Table {
	return Table(fmt.Sprintf("%s AS %s", f, alias))
}

// String implements Modifier.
func (f Table) String() string {
	return string(f)
}

var _ TableModifier = Table("")
