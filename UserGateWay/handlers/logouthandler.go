package handlers

import (
	"net/http"

	"github.com/Tabed23/UserGateWay/res"
	"github.com/Tabed23/UserGateWay/service"
	"github.com/Tabed23/UserGateWay/types"
	"github.com/Tabed23/UserGateWay/utils"
	"github.com/gin-gonic/gin"
)

// [Logout] - User logout Handler (Takes the User Logout Struct and Confirm that then Insert Data into AWS Red Shift)
func Logout(ctx *gin.Context) {
	userlogout := types.UserLogOut{}
	//Check if the JSON is valid with our Struct
	if err := ctx.ShouldBindJSON(&userlogout); err != nil { //will Bind JSON to user logout struct
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status": http.StatusBadRequest,
			"error":  utils.BadRequest.Error(),
		})

		return //exit the function
	}
	//  servie layer
	if err := service.UserLogout(userlogout); err != nil { //check if the err is not nil
		utils.Logger.Info("Logout service error: " + err.Error())

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
