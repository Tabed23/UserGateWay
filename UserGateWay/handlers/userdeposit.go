package handlers

import (
	"net/http"

	"github.com/Tabed23/UserGateWay/res"
	"github.com/Tabed23/UserGateWay/utils"
	"github.com/Tabed23/UserGateWay/service"
	"github.com/Tabed23/UserGateWay/types"
	"github.com/gin-gonic/gin"
)

// [Deposit] - User Deposit Handler (Takes the User Deposit Struct and Confirm that then Insert Data into AWS Red Shift)
func Deposit(ctx *gin.Context) {
	userdeposit := types.UserDeposit{}
	//Check if the JSON is valid with our Struct
	if err := ctx.ShouldBindJSON(&userdeposit); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status": http.StatusBadRequest,
			"error":  utils.BadRequest.Error(),
		})

		return //exit the function
	}

	//  servie layer
	if err := service.UserDeposit(userdeposit); err != nil { //check if the err is not nil
		utils.Logger.Info("Deposit service error: " + err.Error())
		ctx.JSON(http.StatusOK, gin.H{
			"response": res.Response{
				Result:  "fail",
				Code:    http.StatusNotAcceptable,
				Success: false,
				Message: err.Error(),
			},
		})
		return //exit the function
	}

	ctx.JSON(http.StatusOK, gin.H{
		"response": res.Response{
			Result:  "success",
			Code:    http.StatusOK,
			Success: true,
			Message: "ok",
		},
	})
}
