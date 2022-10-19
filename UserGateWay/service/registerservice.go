package service

import (
	"fmt"

	"github.com/Tabed23/UserGateWay/database"
	"github.com/Tabed23/UserGateWay/helper"
	"github.com/Tabed23/UserGateWay/repository"
	"github.com/Tabed23/UserGateWay/types"
	"github.com/Tabed23/UserGateWay/utils"
)

// [Register] - Insert thr Request into database: returns the error
func UseraRegister(reg types.UserRegistration) error {
	utils.Logger.Info("invoke login service for inserting logout data")

	utils.Logger.Info(reg)

	isExist, err := helper.IsExist(reg.MemberID, reg.Mobile, reg.Email, "registration")

	if err != nil {// check if there is a error
		utils.Logger.Error(err)
		return err
	}

	if isExist{
		//return error that this member_id is already in the database
		return fmt.Errorf("the user with member id %d or mobile %s or email %s is already registered", reg.MemberID,reg.Mobile, reg.Email)
	}else{
		stat, err := database.DB.Prepare(repository.Statreg) //get the Query
		if err != nil {
			utils.Logger.Error(err)
			//return error if the query failed
			return err
		}
		//execute the query and insert into database
		_, err = stat.Exec(reg.MemberID, reg.Mobile, reg.MobilePrefix, reg.Email, reg.Status, reg.RegisterIP, reg.AccountStatus, reg.CreatedAt, reg.RegisterPlatform)
		if err != nil {
			utils.Logger.Error("Error inserting register into database: ", err)
			//return the error if it is not inserted into database
			return err
		}

		return nil //return nil if the query is inserted into Database

	}
	
}
