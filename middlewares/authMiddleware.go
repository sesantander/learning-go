package middlewares

import (
	"fmt"
	"net/http"
	"os"
	"strings"
	"test/initializers"
	"test/models"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func RequireAuth(c *gin.Context) {
	authHeader := c.GetHeader("Authorization")

	if authHeader == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header is missing"})
		return
	}

	authParts := strings.Split(authHeader, " ")
	if len(authParts) != 2 || authParts[0] != "Bearer" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid Authorization header format"})
		return
	}

	tokenString := authParts[1]
	token, errorT := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		return []byte(os.Getenv("PRIVATE_KEY")), nil
	})

	if errorT != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"status": false, "message": errorT.Error()})
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		if float64(time.Now().Unix()) > claims["exp"].(float64) {
			c.AbortWithStatus(http.StatusUnauthorized)
		}

		var user models.User
		initializers.DB.First(&user, claims["sub"])

		if user.ID == 0 {
			c.AbortWithStatus(http.StatusUnauthorized)
		}

		c.Set("user", user)

		c.Next()
	} else {
		c.AbortWithStatus(http.StatusUnauthorized)
	}
}
