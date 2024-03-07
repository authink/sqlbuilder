package sqlbuilder

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSelect(t *testing.T) {
	var builder Builder
	sql := builder.
		Select("id", "name").
		From("users").
		Where(&Equal{field: "id"}).
		And(&Equal{field: "name"}).
		String()

	assert.Equal(t, "SELECT id,name FROM users WHERE id = :id AND name = :name", sql)
}

func TestCount(t *testing.T) {
	var builder Builder
	sql := builder.
		Select(Field("id").Count().As("c")).
		From(Table("users").As("u")).
		String()

	assert.Equal(t, "SELECT COUNT(id) AS c FROM users AS u", sql)
}
