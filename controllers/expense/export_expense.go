package expense

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/expense-tracker-api/models"
	"github.com/expense-tracker-api/services"
	"github.com/expense-tracker-api/utils"
	"github.com/gin-gonic/gin"
)

func ExportExpenses(c *gin.Context) {

	var query models.ExpenseFilterQueryModel

	if err := c.ShouldBindQuery(&query); err != nil {
		requestError := utils.FormatValidationError(err)
		var message strings.Builder
		for _, v := range requestError {
			fmt.Fprint(&message, v.Msg+". ")
		}

		c.JSON(http.StatusBadRequest, models.CustomResponseModel{
			Status:  "error",
			Message: message.String(),
		})
		return
	}

	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "10"))
	typeFilter := c.DefaultQuery("type", "Expense")
	userId := utils.GetUserID(c)

	if typeFilter != "" && typeFilter != "Expense" && typeFilter != "Income" {
		c.JSON(http.StatusBadRequest, models.CustomResponseModel{
			Status:  "error",
			Message: "Invalid type filter. Allowed values are 'Expense' or 'Income'.",
		})
		return
	}

	req := models.ExpenseGetQueryModel{
		DateFilter: query,
		Page:       page,
		Limit:      limit,
		UserId:     userId,
		Type:       typeFilter,
	}

	format := strings.ToLower(c.DefaultQuery("format", "csv"))

	data, fileName, contentType, err :=
		services.ExportExpenses(format, req)

	if err != nil {
		c.JSON(http.StatusBadRequest, models.CustomResponseModel{
			Status:  "error",
			Message: err.Error(),
		})
		return
	}

	c.Header(
		"Content-Disposition",
		"attachment; filename="+fileName,
	)

	c.Data(
		http.StatusOK,
		contentType,
		data,
	)
}
