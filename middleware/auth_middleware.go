package middleware

import (
	"net/http"
	"os"
	"strings"

	"github.com/expense-tracker-api/models"
	"github.com/expense-tracker-api/utils"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")

		if strings.TrimSpace(authHeader) == "" {
			c.JSON(http.StatusUnauthorized, models.CustomResponseModel{
				Status:  "error",
				Message: "Authorization header is missing. Please provide Bearer token.",
			})
			c.Abort()
			return
		}

		headerParts := strings.Split(authHeader, " ")
		if len(headerParts) != 2 || headerParts[0] != "Bearer" {
			c.JSON(http.StatusUnauthorized, models.CustomResponseModel{
				Status:  "error",
				Message: "Authentication token is required. Please provide a valid Bearer token.",
			})
			c.Abort()
			return
		}

		tokenString := headerParts[1]

		secret := os.Getenv("JWT_SECRET")
		if secret == "" {
			c.JSON(http.StatusInternalServerError, models.CustomResponseModel{
				Status:  "error",
				Message: "Unexpected server error. Please try again later.",
			})
			c.Abort()
			return
		}

		claims := &utils.JWTClaims{}

		token, err := jwt.ParseWithClaims(
			tokenString,
			claims,
			func(token *jwt.Token) (interface{}, error) {
				_, ok := token.Method.(*jwt.SigningMethodHMAC)
				if !ok {
					return nil, jwt.ErrSignatureInvalid
				}
				return []byte(secret), nil
			},
		)

		if err != nil || !token.Valid {
			c.JSON(http.StatusUnauthorized, models.CustomResponseModel{
				Status:  "error",
				Message: "Unauthorized access. Please log in to continue.",
			})
			c.Abort()
			return
		}

		c.Set("id", claims.UserID)
		c.Set("email", claims.Email)
		c.Set("role", claims.Role)

		c.Next()
	}
}
