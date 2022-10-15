package handlers

import (
	"net/http"

	"github.com/Tabed23/UserGateWay/res"
	"github.com/Tabed23/UserGateWay/utils"
	"github.com/Tabed23/UserGateWay/service"
	"github.com/Tabed23/UserGateWay/types"
	"github.com/gin-gonic/gin"
)

// [Login] - User login Handler (Takes the User Login Struct and Confirm that then Insert Data into AWS Red Shift)
func Login(ctx *gin.Context) {
	var usrlogin types.UserLogin
	//Check if the JSON is valid with our Struct
	if err := ctx.ShouldBindJSON(&usrlogin); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status": http.StatusBadRequest,
			"error":  utils.BadRequest.Error(),
		})

		return //exit the function
	}
	//servie layer to inserting data into redshift
	if err := service.UsersLogin(usrlogin); err != nil { //check if the err is not nil
		utils.Logger.Info("Login service error: " + err.Error())

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

	//User Respose if Service layer has inserted the data
	ctx.JSON(http.StatusOK, gin.H{
		"response": res.Response{
			Result:  "success",
			Code:    http.StatusOK,
			Success: true,
			Message: "ok",
		},
	})
}
