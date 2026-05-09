package expense

import (
	"net/http"

	"github.com/expense-tracker-api/models"
	"github.com/expense-tracker-api/services"
	"github.com/gin-gonic/gin"
)

func ImportExpenses(c *gin.Context) {
	file, err := c.FormFile("file")
	if err != nil {
		c.JSON(http.StatusBadRequest, models.CustomResponseModel{
			Status:  "error",
			Message: "CSV file is required",
		})
		return
	}

	f, err := file.Open()
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.CustomResponseModel{
			Status:  "error",
			Message: "Failed to open file",
		})
		return
	}
	defer f.Close()

	userId := c.GetString("id")

	result, err := services.ImportExpenses(userId, f)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.CustomResponseModel{
			Status:  "error",
			Message: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, models.CustomResponseModel{
		Status:  "success",
		Message: "CSV processed successfully",
		Data:    result,
	})
}
