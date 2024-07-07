package services

import (
	"errors"
	"os"
	"test/hooks"
	"test/initializers"
	"test/models"
	"test/tools"
	"time"

	"github.com/golang-jwt/jwt/v5"

	"gorm.io/gorm"
)

func SignUpService(user models.User) (models.User, *gorm.DB) {

	user = hooks.UserParser(user)

	result := initializers.DB.Create(&user)

	return user, result
}

func Login(user models.User) (string, error) {
	var searchUser models.User

	initializers.DB.Where("username = ?", user.Username).First(&searchUser)

	decryptedPassword, err := tools.DecryptMessage(searchUser.Password, os.Getenv("PRIVATE_KEY"))

	if err != nil {
		panic(err)
	}

	if decryptedPassword != user.Password {
		return "", errors.New("invalid Username/Password")
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": searchUser.ID,
		"exp": time.Now().Add(time.Hour * 24 * 30).Unix(),
	})

	tokenString, err := token.SignedString([]byte(os.Getenv("PRIVATE_KEY")))

	if err != nil {
		return "", errors.New("error generating token")
	}

	return tokenString, err
}
