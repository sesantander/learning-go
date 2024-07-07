package main

import (
	"os"
	"test/controllers"
	"test/initializers"
	"test/middlewares"

	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
}

func main() {
	r := gin.Default()

	v1 := r.Group("/api/v1")
	{
		users := v1.Group("/users")
		{
			users.POST("/", controllers.SignUp)
			users.GET("/", middlewares.RequireAuth, controllers.GetAllUsers)
			users.GET("/:id", controllers.GetOneUser)
			users.PATCH("/:id", controllers.UpdateUser)
			users.POST("/login", controllers.Login)
		}
	}

	r.Run(":" + os.Getenv("PORT"))
}
