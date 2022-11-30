package user_test

import (
	"fmt"
	"os"
	"testing"

	"github.com/bluewon/testing/db/postgres"
	"github.com/bluewon/testing/internal/user"
	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
	pd "gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func TestGetUserByEmail(t *testing.T) {
	//Arrange
	req := user.LoginRequest{
		Email: "whTH@gmail.com",
		Pwd:   "dsfdf",
	}
	srv, cleanFunc := setupSrv(t)
	defer cleanFunc()

	t.Run("Test incorrect pwd", func(t *testing.T) {
		// Act
		_, err := srv.Login(&req)
		if err != nil {
			// Assert
			assert.EqualError(t, err, "password incorrect")
		}
	})
}

func setupSrv(t *testing.T) (*user.Service, func()) {
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
	srv := user.NewService(postgres.New(db))
	return srv, func() {
		//TearDown
		//drop table
		//clean up
	}
}
