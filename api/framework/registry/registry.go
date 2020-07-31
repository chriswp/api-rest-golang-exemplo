package registry

import (
	"api/application/controllers"
	"github.com/jinzhu/gorm"
)

type registry struct {
	db *gorm.DB
}

type Registry interface {
	NewAppController() controllers.AppController
}

func NewRegistry(db *gorm.DB) Registry {
	return &registry{db}
}

func (r *registry) NewAppController() controllers.AppController {
	return r.NewCategoriaController()
}
