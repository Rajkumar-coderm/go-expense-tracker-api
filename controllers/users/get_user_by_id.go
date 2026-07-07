package users

import (
	"net/http"
	"strings"

	"github.com/expense-tracker-api/models"
	"github.com/expense-tracker-api/services"
	"github.com/expense-tracker-api/utils"
	"github.com/gin-gonic/gin"
)

func GetUserById(c *gin.Context) {
	id := c.Query("id")

	if strings.TrimSpace(id) == "" {
		id = utils.GetUserID(c)
	}

	if strings.TrimSpace(id) == "" {
		c.JSON(http.StatusBadRequest, models.CustomResponseModel{
			Status:  "error",
			Message: "ID field is missing. Please provide the ID field.",
		})
		return
	}
	res, err := services.GetUserById(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.CustomResponseModel{
			Status:  "error",
			Message: err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, models.CustomResponseModel{
		Status:  "success",
		Message: "",
		Data:    res,
	})
}
