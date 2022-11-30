package postgres_test

import (
	"testing"

	"github.com/bluewon/testing/utils"
	"github.com/stretchr/testify/assert"
)

func TestGetUserByEmail(t *testing.T) {
	db, cleanFunc := setup(t)
	defer cleanFunc()
	user := createUser(t)
	t.Run("error user not found", func(t *testing.T) {
		_, err := db.GetUserByEmail(utils.RandStringRunes(4) + "@gmail.com")
		assert.NotNil(t, err)
	})
	t.Run("successful", func(t *testing.T) {
		u, err := db.GetUserByEmail(user.Email)
		assert.Nil(t, err)
		assert.Equal(t, user.ID, u.ID)
	})
}
