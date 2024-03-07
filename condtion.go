package sqlbuilder

import "fmt"

type Condition interface {
	fmt.Stringer
}

type Equal struct {
	field string
}

// String implements Condition.
func (e *Equal) String() string {
	return fmt.Sprintf("%s = :%s", e.field, e.field)
}

var _ Condition = (*Equal)(nil)
