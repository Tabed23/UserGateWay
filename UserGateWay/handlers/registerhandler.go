package handlers

import (
	"net/http"

	"github.com/Tabed23/UserGateWay/res"
	"github.com/Tabed23/UserGateWay/service"
	"github.com/Tabed23/UserGateWay/types"
	"github.com/Tabed23/UserGateWay/utils"
	"github.com/gin-gonic/gin"
)

// [Registration] - User Registration Handler (Takes the User Registration Struct and Confirm that then Insert Data into AWS Red Shift)
func Registration(ctx *gin.Context) {
	userreg := types.UserRegistration{}
	//Check if the JSON is valid with our Struct
	if err := ctx.ShouldBindJSON(&userreg); err != nil { //check if the err is not nil
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status": http.StatusBadRequest,
			"error":  utils.BadRequest.Error(),
		})

		return //exit the function
	}

	//  servie layer
	if err := service.UseraRegister(userreg); err != nil {
		utils.Logger.Info("register service error: " + err.Error())

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
