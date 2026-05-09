package users

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/expense-tracker-api/models"
	"github.com/expense-tracker-api/services"
	"github.com/expense-tracker-api/utils"
	"github.com/gin-gonic/gin"
)

func UserLogin(ctx *gin.Context) {

	var req models.LoginReq
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

	err, user := services.UserLogin(req)
	if err != nil {
		ctx.JSON(
			http.StatusBadRequest,
			models.CustomResponseModel{
				Status:  "error",
				Message: err.Error(),
			},
		)
		return
	}
	ctx.JSON(http.StatusOK, models.CustomResponseModel{
		Status: "success",
		Data:   user,
	})
}
