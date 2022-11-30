package postgres_test

import (
	"crypto/sha1"
	"fmt"
	"os"
	"testing"

	"github.com/bluewon/testing/db/postgres"
	"github.com/bluewon/testing/internal/user"
	"github.com/bluewon/testing/utils"
	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
	pd "gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	dbConnection *gorm.DB
)

func TestMain(m *testing.M) {
	os.Exit(m.Run())
}

func setup(t *testing.T) (*postgres.Database, func()) {

	err := godotenv.Load()
	if err != nil {
		t.Fatal("Error loading .env file")
	}
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=UTC",
		os.Getenv("HOST_DB_TEST"),
		os.Getenv("USERNAME_TEST"),
		os.Getenv("PASSWORD_TEST"),
		os.Getenv("DATABASE_TEST"),
		os.Getenv("PORT_TEST"))

	db, err := gorm.Open(pd.Open(dsn), &gorm.Config{})
	assert.Nil(t, err)

	err = db.AutoMigrate(&user.User{})
	assert.Nil(t, err)

	dbConnection = db

	return postgres.New(db), func() {
		//TearDown
		//drop table
		//clean up
	}
}

func createUser(t *testing.T) *user.User {
	salt := utils.RandStringRunes(6)
	pwd := utils.RandStringRunes(9)
	h := sha1.New()
	h.Write([]byte(pwd + salt))
	hash := h.Sum(nil)

	users := &user.User{
		ID:       1,
		Email:    utils.RandStringRunes(4) + "@gmail.com",
		Password: pwd,
		Name:     utils.RandStringRunes(4),
		Salt:     salt,
		Hash:     fmt.Sprintf("%x", hash),
	}
	err := dbConnection.Create(users).Error
	assert.Nil(t, err)
	return users
}
