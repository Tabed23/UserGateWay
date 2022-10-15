package service

import (
	"github.com/Tabed23/UserGateWay/database"
	"github.com/Tabed23/UserGateWay/repository"
	"github.com/Tabed23/UserGateWay/types"
	"github.com/Tabed23/UserGateWay/utils"
)

// [Login] - Insert thr Request into database: returns the error
func UsersLogin(login types.UserLogin) error {
	utils.Logger.Info("invoke login service for inserting login data")

	utils.Logger.Info(login)

	stat, err := database.DB.Prepare(repository.Statlogin) //get the Query
	if err != nil {
		utils.Logger.Error(err)
		//return error if the query failed
		return err
	}
	//execute the query and insert into database
	_, err = stat.Exec(login.MemberID, login.LoginTime, login.LoginIP, login.Platform)
	if err != nil {
		utils.Logger.Error("Error inserting login into database: ", err)

		return err //return the error if it is not inserted into database
	}

	return nil //return nil if the query is inserted into Database
}
