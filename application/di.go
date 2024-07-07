package application

import (
	"github.com/samber/do"
	"github.com/tadasi/portfolio/infrastructure/repositories"
	"github.com/tadasi/portfolio/presentation/controllers"
)

func RegisterInjector() *do.Injector {
	injector := do.New()
	// infrastructure
	do.Provide(injector, repositories.NewTodoRepository)
	// presentation
	do.Provide(injector, controllers.NewController)
	do.Provide(injector, controllers.NewTodoController)
	return injector
}
