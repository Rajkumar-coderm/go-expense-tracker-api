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

func CreateNewExpense(ctx *gin.Context) {
	var req models.CreateExpenseRequestModel

	if err := ctx.ShouldBindJSON(&req); err != nil {
		requestError := utils.FormatValidationError(err)
		var message strings.Builder
		for _, v := range requestError {
			fmt.Fprint(&message, v.Msg+". ")
		}

		ctx.JSON(http.StatusBadRequest, models.CustomResponseModel{
			Status:  "error",
			Message: message.String(),
		})
		return
	}
	userId := ctx.GetString("id")
	if strings.TrimSpace(userId) == "" {
		ctx.JSON(http.StatusUnauthorized, models.CustomResponseModel{
			Status:  "error",
			Message: "Unouthenticated: User ID is missing",
		})
		return
	}

	if !utils.ValidatePaymentMethod(strings.ToLower(req.Type), strings.ToLower(req.PaymentMethod)) {
		ctx.JSON(http.StatusBadRequest, models.CustomResponseModel{
			Status:  "error",
			Message: "Invalid payment method for the given transaction type",
		})
		return
	}

	err := services.CreateNewExpense(req, userId)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, models.CustomResponseModel{
			Status:  "error",
			Message: err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusCreated, models.CustomResponseModel{
		Status:  "success",
		Message: "Expense Created Successfully",
	})
}
