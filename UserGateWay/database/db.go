package database

import (
	"database/sql"
	"os"

	"github.com/Tabed23/UserGateWay/utils"
	_ "github.com/lib/pq"
)

var (
	DB *sql.DB
)

// RedShiftDatabaseConnection - Connect to AWS Red Shift Database (VPC)
func RedShiftDatabaseConnection() {

	var err error
	//get the database cerds from lambda eviroment
	host := os.Getenv("host")
	port := os.Getenv("port")
	user := os.Getenv("user")
	password := os.Getenv("password")
	dbname := os.Getenv("db")
	connStr := "postgres://" + user + ":" + password + "@" + host + ":" + port + "/" + dbname + ""

	DB, err = sql.Open("postgres", connStr) // open the database connection
	if err != nil {
		utils.Logger.Infof("Error creating database connection: %v", err)
	}

	utils.Logger.Info("connection has been made with AWS redshift")
}
