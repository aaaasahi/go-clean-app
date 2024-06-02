package main

import (
	"log"

	"go-clean-app/config"
	"go-clean-app/server"
	"go-clean-app/di"
	"go-clean-app/infra/mysql/db"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	var c config.Conf
	if err := c.Init(); err != nil {
		log.Fatal(err)
	} 

	dbConn, err := db.Connect(c)
	if err != nil {
		log.Fatalf("Failed to connect to the database: %v", err)
	}
	defer dbConn.Close()

	container := di.BuildContainer(dbConn)

	if err := container.Invoke(func(s *server.Server) {
		s.Run(&c)
	}); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}