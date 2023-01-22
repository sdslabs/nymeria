package db

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"

	"github.com/sdslabs/nymeria/config"
	"github.com/sdslabs/nymeria/log"
)

func Connection() (*sql.DB, error) {
	var connStr string
	if config.NymeriaConfig.DB.DSN != "" {
		connStr = config.NymeriaConfig.DB.DSN
	} else {
		connStr = fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
			config.NymeriaConfig.DB.Host,
			config.NymeriaConfig.DB.Port,
			config.NymeriaConfig.DB.User,
			config.NymeriaConfig.DB.Password,
			config.NymeriaConfig.DB.DBName)
	}

	db, err := sql.Open("postgres", connStr)

	if err != nil {
		log.ErrorLogger("DB connection failed", err)
		return nil, err
	}

	return db, nil

}
