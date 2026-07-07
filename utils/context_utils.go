package utils

import (
	"github.com/expense-tracker-api/constants"
	"github.com/gin-gonic/gin"
)

func GetUserID(c *gin.Context) string {
	return c.MustGet(constants.ContextUserID).(string)
}

func GetUserEmail(c *gin.Context) string {
	return c.MustGet(constants.ContextEmail).(string)
}

func GetUserRole(c *gin.Context) string {
	return c.MustGet(constants.ContextRole).(string)
}
