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

type True struct {
	Left Field
}

// String implements Condition.
func (t True) String() string {
	return fmt.Sprintf("%s = 1", t.Left)
}

var _ Condition = True{}

type False struct {
	Left Field
}

// String implements Condition.
func (f False) String() string {
	return fmt.Sprintf("%s = 0", f.Left)
}

var _ Condition = False{}
