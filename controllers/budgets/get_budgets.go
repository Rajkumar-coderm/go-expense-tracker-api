package budgets

import (
	"strconv"

	"github.com/expense-tracker-api/models"
	"github.com/expense-tracker-api/services"
	"github.com/gin-gonic/gin"
)

func GetBudgets(ctx *gin.Context) {
	userId := ctx.GetString("id")
	if userId == "" {
		ctx.JSON(401, gin.H{"error": "Unauthorized: User ID is missing"})
		return
	}

	page, _ := strconv.Atoi(ctx.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(ctx.DefaultQuery("limit", "10"))

	budgets, err := services.GetBudget(models.GetBudgetsRequestModel{
		UserId:   userId,
		Category: ctx.Query("category"),
		Page:     page,
		Limit:    limit,
	})
	if err != nil {
		ctx.JSON(500, models.CustomResponseModel{
			Status:  "error",
			Message: "Failed to fetch budgets: " + err.Error(),
		})
		return
	}
	if len(budgets) == 0 {
		ctx.JSON(204, models.CustomResponseModel{
			Status:  "success",
			Message: "No budgets found",
		})
		return
	}
	ctx.JSON(200, models.CustomResponseModel{
		Status:  "success",
		Message: "Budgets fetched successfully",
		Data:    budgets,
	})
}
