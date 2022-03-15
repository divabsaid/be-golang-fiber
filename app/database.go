package app

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/spf13/viper"
)

func InitDatabase(config *viper.Viper) *sql.DB {
	// init database
	dbHost := config.GetString(`database.host`)
	dbPort := config.GetString(`database.port`)
	dbUser := config.GetString(`database.username`)
	dbPass := config.GetString(`database.password`)
	dbName := config.GetString(`database.name`)
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true", dbUser, dbPass, dbHost, dbPort, dbName)
	dbConn, err := sql.Open(`mysql`, dsn)
	if err != nil {
		panic(fmt.Errorf("Fatal error database connection: %s \n", err))
	}
	return dbConn
}
