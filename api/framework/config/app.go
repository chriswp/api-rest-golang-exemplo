package config

import (
	"api/framework/database"
	"github.com/jinzhu/gorm"
)

type App struct {
	DB *gorm.DB
}

func (a *App) Initialize() {
	db := database.NewDb()
	a.DB = db
}
