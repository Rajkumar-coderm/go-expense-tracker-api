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

func GetExpense(c *gin.Context) {
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
	userId := c.GetString("id")

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

	expenses, err := services.GetExpense(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.CustomResponseModel{
			Status:  "error",
			Message: "Failed to fetch expenses: " + err.Error(),
		})
		return
	}

	if len(expenses) == 0 {
		c.JSON(http.StatusNoContent, models.CustomResponseModel{
			Status:  "success",
			Message: "No expenses found",
		})
		return
	}

	c.JSON(http.StatusOK, models.CustomResponseModel{
		Status:  "success",
		Message: "Expenses fetched successfully",
		Data:    expenses,
	})
}
