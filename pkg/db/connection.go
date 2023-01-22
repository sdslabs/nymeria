package db

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"

	"github.com/sdslabs/nymeria/config"
	"github.com/sdslabs/nymeria/log"
)

func Connection() error {
	var connStr string
	if config.KratosClientConfig.DB.DSN != "" {
		connStr = config.KratosClientConfig.DB.DSN
	} esle {
		connStr = fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s",
		 	config.KratosClientConfig.DB.Host,
		  	config.KratosClientConfig.DB.Port,
		   	config.KratosClientConfig.DB.User,
		    config.KratosClientConfig.DB.Password,
			config.KratosClientConfig.DB.DBName)
	}

	db, err := sql.Open("postgres", connStr)

	if err != nil {
		log.ErrorLogger("DB connection failed", err)
		return err
	}

	
}
