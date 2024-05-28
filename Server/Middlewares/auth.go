package Middlewares

import (
	"net/http"

	"github.com/Lucasmartinsn/grocery-api/Services"
	"github.com/gin-gonic/gin"
)

func Auth() gin.HandlerFunc {
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
		if valid, err := Services.NewJWTService().ValidateToken(token); err != nil || !valid {
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