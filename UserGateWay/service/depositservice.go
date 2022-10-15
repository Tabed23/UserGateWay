package service

import (
	"github.com/Tabed23/UserGateWay/database"
	"github.com/Tabed23/UserGateWay/repository"
	"github.com/Tabed23/UserGateWay/types"
	"github.com/Tabed23/UserGateWay/utils"
)

// [Deposit]- Insert thr Request into database: returns the error
func UserDeposit(dep types.UserDeposit) error {
	utils.Logger.Info("invoke login service for inserting logout data")

	utils.Logger.Info(dep)

	stat, err := database.DB.Prepare(repository.Statdeposit) //get the Query
	if err != nil {
		utils.Logger.Error(err)
		//return error if the query failed
		return err
	}
	//execute the query and insert into database
	_, err = stat.Exec(dep.ID, dep.Coin, dep.ChainDetail, dep.MemberID, dep.Amount, dep.Status, dep.RequestTime, dep.RequestDeadline, dep.DepositLabel, dep.UpdatedTime)
	if err != nil {
		utils.Logger.Error("Error inserting register into database: ", err)

		return err // return the error if it is not inserted into database
	}

	return nil //return nil if the query is inserted into database
}
