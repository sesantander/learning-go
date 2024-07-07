package hooks

import (
	"os"
	"test/models"
	"test/tools"
)

func UserParser(user models.User) models.User {
	var secretKey = os.Getenv("PRIVATE_KEY")
	var encryptedPassword, err = tools.EncryptMessage(user.Password, secretKey)

	if err != nil {
		panic(err)
	}

	return models.User{
		Name:     user.Name,
		Email:    user.Email,
		Age:      user.Age,
		Username: user.Username,
		Password: encryptedPassword,
	}
}
