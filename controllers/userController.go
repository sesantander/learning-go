package controllers

import (
	"fmt"
	"net/http"
	"test/initializers"
	"test/models"
	"test/services"

	"github.com/gin-gonic/gin"
)

func SignUp(c *gin.Context) {

	body := models.User{}

	err := c.ShouldBindJSON(&body)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	fmt.Printf("Received User: %+v\n", body)

	user, result := services.SignUpService(body)

	c.JSON(200, gin.H{
		"data":  user,
		"count": result.RowsAffected,
	})
}

func Login(c *gin.Context) {
	body := models.User{}

	err := c.ShouldBindJSON(&body)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	tokenString, err := services.Login(body)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token_string": tokenString})
}

func Validate(c *gin.Context) {
	user, _ := c.Get("user")

	c.JSON(http.StatusOK, gin.H{
		"message": user,
	})

	c.JSON(200, gin.H{
		"data": user,
	})
}

func GetAllUsers(c *gin.Context) {

	var users []models.User
	initializers.DB.Find(&users)

	c.JSON(200, gin.H{
		"data": users,
	})
}

func GetOneUser(c *gin.Context) {

	id := c.Param("id")

	var user models.User
	initializers.DB.First(&user, id)

	c.JSON(200, gin.H{
		"data": user,
	})
}

func UpdateUser(c *gin.Context) {

	id := c.Param("id")

	body := models.User{}

	err := c.ShouldBindJSON(&body)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var user models.User

	initializers.DB.First(&user, id)

	initializers.DB.Model(&user).Updates(models.User{
		Name:  body.Name,
		Email: body.Email,
		Age:   body.Age,
	})

	c.JSON(200, gin.H{
		"data": user,
	})
}
