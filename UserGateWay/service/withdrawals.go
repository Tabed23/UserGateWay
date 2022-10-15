package service

import (
	"github.com/Tabed23/UserGateWay/database"
	"github.com/Tabed23/UserGateWay/repository"
	"github.com/Tabed23/UserGateWay/types"
	"github.com/Tabed23/UserGateWay/utils"
)

// [Withdrawals] - Insert thr Request into database: returns the error
func UserWithdrawl(width types.UserWithdrawals) error {
	utils.Logger.Info("invoke login service for inserting logout data")

	utils.Logger.Info(width)

	stat, err := database.DB.Prepare(repository.Statwidthdrawl) //get the Query
	if err != nil {
		utils.Logger.Error(err)
		//return error if the query failed
		return err
	}

	_, err = stat.Exec(width.ID, width.MemberID, width.Coin, width.Status, width.Amount, width.Fee, width.Address, width.Txid, width.SubmitedAt, width.SuccessedAt, width.UpdatedAt, width.ChainType, width.HashInfo)
	if err != nil {
		utils.Logger.Error("Error inserting register into database: ", err)
		//return the error if it is not inserted into database
		return err
	}

	return nil //return nil if the query is inserted into Database
}
