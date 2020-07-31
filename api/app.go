package main

import (
	"api/framework/database"
	"github.com/jinzhu/gorm"
	"log"
	"os"
)

type App struct {
	DB *gorm.DB
}

func (a *App) Initialize() {
	dbInstace := database.NewDb()
	dbInstace.Dsn = os.Getenv("DSN")
	dbInstace.DbType = os.Getenv("DB_TYPE")
	dbInstace.Debug = false
	dbInstace.AutoMigrateDb = false
	dbInstace.Env = os.Getenv("ENV")

	connection, err := dbInstace.Connect()

	if err != nil {
		log.Fatalf("Problema ao conectar no banco. erro: %v", err)
	}

	a.DB = connection
}
