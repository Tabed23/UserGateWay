package helper

import (
	"github.com/Tabed23/UserGateWay/database"
	"github.com/Tabed23/UserGateWay/utils"
)

//[IsExist] - check if the member_id , email or mobile is already
func IsExist(member_id int, mobile, email,  tablename string )(bool, error){
	utils.Logger.Info("Check if meber , email or mobile IsExist")

	row, err := database.DB.Query("SELECT * FROM "+ tablename + " WHERE member_id = $1 or mobile = $2 or email=$3", member_id, mobile, email)

	utils.Logger.Info(row)

	if err != nil {

		utils.Logger.Errorf("Error checking IsExist: %v", err)

		return true, err
	}
	if row.Next(){

		return true, nil
	}else {

		return false, nil
	}
}