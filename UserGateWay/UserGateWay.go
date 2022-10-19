package main

import (
	"log"
	"net/http"

	"github.com/Tabed23/UserGateWay/database"
	"github.com/Tabed23/UserGateWay/handlers"
	"github.com/apex/gateway"
	"github.com/gin-gonic/gin"
)

// Routing of the API endpoint
func routerEngine() *gin.Engine {
	// set server mode
	gin.SetMode(gin.DebugMode)

	r := gin.New() //new gin Engine

	// Global middleware
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	//lambda routes
	r.GET("/", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"status": http.StatusOK,
			"message":  "lambda with gin is running",
		})
	})
	r.POST("/user/registration", handlers.Registration)
	r.POST("/user/login", handlers.Login)
	r.POST("/user/logout", handlers.Logout)
	r.POST("/user/deposits", handlers.Deposit)
	r.POST("/user/withdraws", handlers.Withdrawals)

	return r // return the new gin Engine
}

func main() {
	database.RedShiftDatabaseConnection()                   // make a connection to the Database so the (DB.sql) can be initlize
	addr := ":" + "3000"                                    //giving the port
	log.Fatal(gateway.ListenAndServe(addr, routerEngine())) // this will listenAndServe to lambda with gin
}
