package db

import (
	"database/sql"
	"fmt"
	"go-clean-app/config"
	"log"
	"time"
)

func Connect(c config.Conf) (*sql.DB, error) {
	dataSourceName := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true", c.DB.User, c.DB.Password, c.DB.Host, c.DB.Port, c.DB.Name)

	var db *sql.DB
	var err error
	for i := 0; i < 10; i++ {
		db, err = sql.Open("mysql", dataSourceName)
		if err == nil {
			err = db.Ping()
			if err == nil {
				break
			}
		}
		log.Printf("Waiting for database to be ready... (%d/10)\n", i+1)
		time.Sleep(5 * time.Second)
	}
	if err != nil {
		return nil, fmt.Errorf("failed to connect to the database: %w", err)
	}

	return db, nil
}