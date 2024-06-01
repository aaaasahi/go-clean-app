package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"time"

	"go-clean-app/config"

	_ "github.com/go-sql-driver/mysql"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
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

	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	log.Println("Server starting on port 8080")

	if err := e.Start(":8080"); err != nil {
		log.Fatal(err)
	}
}