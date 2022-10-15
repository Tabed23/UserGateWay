package service

import (
	"github.com/Tabed23/UserGateWay/database"
	"github.com/Tabed23/UserGateWay/repository"
	"github.com/Tabed23/UserGateWay/types"
	"github.com/Tabed23/UserGateWay/utils"
)

// [LogOut] - Insert thr Request into database: returns the error
func UserLogout(logout types.UserLogOut) error {
	utils.Logger.Info("invoke login service for inserting logout data")

	utils.Logger.Info(logout)

	stat, err := database.DB.Prepare(repository.Statlogout) //get the Query
	if err != nil {
		utils.Logger.Error(err)
		//return error if the query failed
		return err
	}
	//execute the query and insert into database
	_, err = stat.Exec(logout.MemberID, logout.LogoutTime, logout.LogoutIP, logout.Platform)
	if err != nil {
		utils.Logger.Error("Error inserting logout into database: ", err)

		return err //return the error if it is not inserted into database
	}

	return nil //return nil if the query is inserted into Database
}
