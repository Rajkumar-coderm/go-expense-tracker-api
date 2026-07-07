package budgets

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/expense-tracker-api/models"
	"github.com/expense-tracker-api/services"
	"github.com/expense-tracker-api/utils"
	"github.com/gin-gonic/gin"
)

func CreateNewBudget(ctx *gin.Context) {
	var req models.CreateBudgetRequestModel

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

	userId := utils.GetUserID(ctx)
	if strings.TrimSpace(userId) == "" {
		ctx.JSON(http.StatusUnauthorized, models.CustomResponseModel{
			Status:  "error",
			Message: "Unouthenticated: User ID is missing",
		})
		return
	}

	err := services.CreateNewBudget(req, userId)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, models.CustomResponseModel{
			Status:  "error",
			Message: "Failed to create budget",
		})
		return
	}

	ctx.JSON(http.StatusCreated, models.CustomResponseModel{
		Status:  "success",
		Message: "Budget created successfully",
	})
}
