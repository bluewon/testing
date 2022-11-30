package postgres

import (
	"github.com/bluewon/testing/internal/user"
	"gorm.io/gorm"
)

type Database struct {
	connecion *gorm.DB
}

func New(connection *gorm.DB) *Database {
	return &Database{connecion: connection}
}

func (d *Database) GetUserByEmail(email string) (*user.User, error) {
	user := new(user.User)
	if err := d.connecion.Model(user).Where("email = ?", email).First(user).Error; err != nil {
		return nil, err
	}
	return user, nil
}
