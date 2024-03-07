package sqlbuilder

import "fmt"

type FieldModifier interface {
	fmt.Stringer
	Count() Field
	As(string) Field
}

type Field string

// As implements FieldModifier.
func (f Field) As(alias string) Field {
	return Field(fmt.Sprintf("%s AS %s", f, alias))
}

// Count implements FieldModifier.
func (f Field) Count() Field {
	return Field(fmt.Sprintf("COUNT(%s)", f))
}

// String implements FieldModifier.
func (f Field) String() string {
	return string(f)
}

var _ FieldModifier = Field("")
