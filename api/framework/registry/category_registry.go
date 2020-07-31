package registry

import (
	"api/application/controllers"
	cp "api/application/presenters"
	cr "api/application/repositories"
	"api/application/services"
	ip "api/application/usecase/presenters"
	ir "api/application/usecase/repositories"
)

func (r *registry) NewCategoriaController() controllers.CategoriaController {
	return controllers.NewCategoriaController(r.NewCategoriaService())
}

func (r *registry) NewCategoriaService() services.CategoriaService {
	return services.NewCategoriaService(r.NewCategoriaRepository(), r.NewCategoriaPresenter())
}

func (r *registry) NewCategoriaRepository() ir.CategoriaRepository {
	return cr.NewCategoriaRepository(r.db)
}

func (r *registry) NewCategoriaPresenter() ip.CategoriaPresenter {
	return cp.NewCategoriaPresenter()
}
