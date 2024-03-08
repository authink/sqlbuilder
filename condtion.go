package sqlbuilder

import "fmt"

type Condition interface {
	fmt.Stringer
}

type Equal struct {
	Left  Field
	Right Field
}

// String implements Condition.
func (e Equal) String() string {
	var right = e.Right
	if right == "" {
		right = e.Left.Named()
	}
	return fmt.Sprintf("%s = %s", e.Left, right)
}

var _ Condition = Equal{}
