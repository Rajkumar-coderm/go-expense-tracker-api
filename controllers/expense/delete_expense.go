package expense

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/expense-tracker-api/models"
	"github.com/expense-tracker-api/services"
	"github.com/expense-tracker-api/utils"
	"github.com/gin-gonic/gin"
)

func DeleteExpense(ctx *gin.Context) {
	userId := ctx.GetString("id")
	expenseId := ctx.Param("id")
	if strings.TrimSpace(expenseId) == "" {
		ctx.JSON(http.StatusBadRequest, models.CustomResponseModel{
			Status:  "error",
			Message: "Expense ID parameter is required and cannot be empty",
		})
		return
	}
	err := services.DeleteExpense(expenseId, userId)
	if err != nil {
		statusCode, e1 := utils.FormatDBError(err)
		var message strings.Builder
		for _, v := range e1 {
			fmt.Fprint(&message, v+". ")
		}
		ctx.JSON(statusCode, models.CustomResponseModel{
			Status:  "error",
			Message: message.String(),
		})
		return
	}
	ctx.JSON(http.StatusOK, models.CustomResponseModel{
		Status:  "success",
		Message: "Expense deleted successfully",
	})
}
