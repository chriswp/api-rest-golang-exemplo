package registry

import (
	"api/application/controllers"
	"api/application/repositories"
	"api/application/services"
)

func (r *registry) NewCategoriaController() controllers.CategoriaController {
	return controllers.NewCategoriaController(r.NewCategoriaService())
}

func (r *registry) NewCategoriaService() services.CategoriaService {
	return services.NewCategoriaService(r.NewCategoriaRepository(), r.NewCategoriaPresenter())
}

func (r *registry) NewCategoriaRepository() repositories.CategoriaRepository {
	return ir.(r.db)
}

func (r *registry) NewCategoriaPresenter() cr.CategoriaPresenter {
	return ip.NewCategoriaPresenter()
}
