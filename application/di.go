package application

import (
	"github.com/samber/do"
	"github.com/tadasi/portfolio/application/usecases"
	"github.com/tadasi/portfolio/domain/models"
	"github.com/tadasi/portfolio/infrastructure/repositories"
	"github.com/tadasi/portfolio/presentation/controllers"
)

func RegisterInjector() *do.Injector {
	injector := do.New()
	//
	// Infrastructure
	//
	// Repositories
	do.Provide(injector, repositories.NewTodoRepository)

	//
	// Presentation
	//
	// Controllers
	do.Provide(injector, controllers.NewController)
	do.Provide(injector, controllers.NewTodoController)

	//
	// Application
	//
	// Usecases
	do.Provide(injector, usecases.NewTodoInteractor)

	//
	// Domain
	//
	// Models
	do.Provide(injector, models.NewTodoFactory)
	return injector
}
