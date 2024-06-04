package Middlewares

import (
	"net/http"

	Encryptiontoken "github.com/Lucasmartinsn/grocery-api/Services/EncryptionToken"
	"github.com/gin-gonic/gin"
)

func Auth_default() gin.HandlerFunc {
	return func(c *gin.Context) {
		const BearerSchema = "Bearer "
		header := c.GetHeader("Authorization")
		if header == "" {
			c.JSON(http.StatusUnauthorized, gin.H{
				"mensage": "token authentication not provided",
			})
			c.Abort()
			return
		}

		token := header[len(BearerSchema):]
		if valid, err := Encryptiontoken.NewJWTService_Default().ValidateToken_Default(token); err != nil || !valid {
			c.JSON(http.StatusUnauthorized, gin.H{
				"mensage": "invalid token authentication",
				"error":   err.Error(),
			})
			c.Abort()
			return
		}
		c.Next()
	}
}

func Auth_customer() gin.HandlerFunc {
	return func(c *gin.Context) {
		const BearerSchema = "Bearer "
		header := c.GetHeader("Authorization")
		if header == "" {
			c.JSON(http.StatusUnauthorized, gin.H{
				"mensage": "token authentication not provided",
			})
			c.Abort()
			return
		}

		token := header[len(BearerSchema):]
		if valid, err := Encryptiontoken.NewJWTService_Customer().ValidateToken_Customer(token); err != nil || !valid {
			c.JSON(http.StatusUnauthorized, gin.H{
				"mensage": "invalid token authentication",
				"error":   err.Error(),
			})
			c.Abort()
			return
		}
		c.Next()
	}
}
