package main

import (
	"finance/config"
	"finance/factory"
	"finance/infrastruktur/database/mysql"
	"fmt"
	"log"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	cfg := config.GetConfig()
	db := mysql.InitDB(cfg)
	mysql.MigrateData(db)
	mysql.SeedUsers(db)
	e := echo.New()

	factory.Initfactory(e, db)

	fmt.Println("Running program ....")
	dsn := fmt.Sprintf(":%d", config.SERVERPORT)
	e.Logger.Fatal(e.Start(dsn))
}
