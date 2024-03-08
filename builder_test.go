package sqlbuilder

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSelect(t *testing.T) {
	sql := NewBuilder().
		Select("id", "name").
		From("users").
		Where(Equal{Left: "id"}).
		And(Equal{Left: "name"}).
		ForUpdate().
		String()

	assert.Equal(t, "SELECT id,name FROM users WHERE id = :id AND name = :name FOR UPDATE", sql)
}

func TestCount(t *testing.T) {
	var tAlias = "u"
	sql := NewBuilder().
		Select(Field("id").Of(tAlias).Count().As("c")).
		From(Table("users").As(tAlias)).
		String()

	assert.Equal(t, "SELECT COUNT(u.id) AS c FROM users AS u", sql)
}

func TestMultiTables(t *testing.T) {
	var (
		tAlias1 = "u"
		tAlias2 = "a"
	)
	sql := NewBuilder().
		Select(
			Field("id").Of(tAlias1),
			Field("name").Of(tAlias2).As("app_name"),
		).
		From(
			Table("users").As(tAlias1),
			Table("apps").As(tAlias2),
		).
		Where(Equal{
			Left:  Field("id").Of(tAlias1),
			Right: Field("app_id").Of(tAlias2),
		}).
		String()

	assert.Equal(t, "SELECT u.id,a.name AS app_name FROM users AS u,apps AS a WHERE u.id = a.app_id", sql)
}

func TestInsert(t *testing.T) {
	sql := NewBuilder().
		InsertInto("users").
		Columns("id", "name").
		String()

	assert.Equal(t, "INSERT INTO users(id,name) VALUES(:id,:name)", sql)
}

func TestUpdate(t *testing.T) {
	sql := NewBuilder().
		Update("users").
		Set("email", "name").
		Where(Equal{Left: "id"}).
		String()

	assert.Equal(t, "UPDATE users SET email = :email,name = :name WHERE id = :id", sql)
}

func TestDelete(t *testing.T) {
	sql := NewBuilder().
		DeleteFrom("users").
		Where(Equal{Left: "id"}).
		String()

	assert.Equal(t, "DELETE FROM users WHERE id = :id", sql)
}
