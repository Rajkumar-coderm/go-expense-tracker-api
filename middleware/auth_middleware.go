package middleware

import (
	"net/http"
	"strings"

	"github.com/expense-tracker-api/constants"
	"github.com/expense-tracker-api/models"
	"github.com/expense-tracker-api/repository"
	"github.com/expense-tracker-api/utils"
	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {

	return func(c *gin.Context) {

		authHeader := strings.TrimSpace(c.GetHeader("Authorization"))

		if authHeader == "" {

			c.JSON(http.StatusUnauthorized, models.CustomResponseModel{
				Status:  "error",
				Message: "Authorization header is missing. Please provide a Bearer token.",
			})
			c.Abort()
			return
		}

		headerParts := strings.Fields(authHeader)

		if len(headerParts) != 2 || !strings.EqualFold(headerParts[0], "Bearer") {

			c.JSON(http.StatusUnauthorized, models.CustomResponseModel{
				Status:  "error",
				Message: "Authentication failed. Please provide a valid Bearer token.",
			})

			c.Abort()
			return
		}

		claims, err := utils.ValidateAccessToken(headerParts[1])

		if err != nil {

			c.JSON(http.StatusUnauthorized, models.CustomResponseModel{
				Status:  "error",
				Message: "Authentication failed. Please log in again.",
			})

			c.Abort()
			return
		}

		_, err = repository.GetUserById(claims.UserID)

		if err != nil {

			c.JSON(http.StatusUnauthorized, models.CustomResponseModel{
				Status:  "error",
				Message: "Authentication failed. Please log in again.",
			})

			c.Abort()
			return
		}

		c.Set(constants.ContextUserID, claims.UserID)
		c.Set(constants.ContextEmail, claims.Email)
		c.Set(constants.ContextRole, claims.Role)

		c.Next()
	}
}
