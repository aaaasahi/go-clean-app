package main

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	"go-clean-app/config"
	"go-clean-app/server"
	"go-clean-app/di"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	var c config.Conf
	if err := c.Init(); err != nil {
		log.Fatal(err)
	} 

	dataSourceName := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", c.DB.User, c.DB.Password, c.DB.Host, c.DB.Port, c.DB.Name)

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
		log.Fatal("Failed to connect to the database:", err)
	}
	defer db.Close()

	container := di.BuildContainer(db)

	if err := container.Invoke(func(s *server.Server) {
		s.Run(&c)
	}); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}