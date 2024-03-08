package sqlbuilder

import "fmt"

type FieldModifier interface {
	fmt.Stringer
	Count() Field
	As(string) Field
	Of(string) Field
	Named() Field
	Assign() Field
}

type Field string

// Assign implements FieldModifier.
func (f Field) Assign() Field {
	return Field(fmt.Sprintf("%s = %s", f, f.Named()))
}

// Named implements FieldModifier.
func (f Field) Named() Field {
	return Field(fmt.Sprintf(":%s", f))
}

// Of implements FieldModifier.
func (f Field) Of(tableAlias string) Field {
	return Field(fmt.Sprintf("%s.%s", tableAlias, f))
}

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
