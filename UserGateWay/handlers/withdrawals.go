package handlers

import (
	"net/http"

	"github.com/Tabed23/UserGateWay/res"
	"github.com/Tabed23/UserGateWay/service"
	"github.com/Tabed23/UserGateWay/types"
	"github.com/Tabed23/UserGateWay/utils"
	"github.com/gin-gonic/gin"
)

//[Withdrawals] - User Withdrawals (Takes the User Withdrawals Struct and Confirm that then Insert Data into AWS Red Shift)
func  Withdrawals(ctx *gin.Context){
	userwithdraw := types.UserWithdrawals{}
		//Check if the JSON is valid with our Struct
	if err := ctx.ShouldBindJSON(&userwithdraw); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"error": utils.BadRequest.Error(),
		})

		return //exit the function
	}

	//  servie layer
	if err:= service.UserWithdrawl(userwithdraw); err != nil {//check if the err is not nil
		utils.Logger.Info("Withdrwal service error: " + err.Error())

		ctx.JSON(http.StatusOK, gin.H{
			"response": res.Response{
				Result: "fail",
				Code: http.StatusNotAcceptable,
				Success: false,
				Message: err.Error(),
			},
		})

		return //exit the function
	}

	 ctx.JSON(http.StatusOK, gin.H{
		"response": res.Response{
			Result: "success",
			Code:  http.StatusOK,
			Success: true,
			Message: "ok",
		},
	})
}