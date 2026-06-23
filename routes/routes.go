package routes

import (
	"github.com/expense-tracker-api/controllers/budgets"
	"github.com/expense-tracker-api/controllers/expense"
	"github.com/expense-tracker-api/controllers/users"
	"github.com/expense-tracker-api/middleware"
	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine) {
	api := r.Group("/api")

	api.POST("/signup", users.UserSignup)
	api.POST("/login", users.UserLogin)

	api.Use(middleware.AuthMiddleware())
	api.GET("/profile", users.GetUserById)

	// Expense Roues
	api.POST("/expense", expense.CreateNewExpense)
	api.GET("/expense", expense.GetExpense)
	api.DELETE("/expense/:id", expense.DeleteExpense)
	api.POST("/expense/import", expense.ImportExpenses)
	api.GET("/expense/export", expense.ExportExpenses)

	// Budget Routes
	api.POST("/budget", budgets.CreateNewBudget)

}
