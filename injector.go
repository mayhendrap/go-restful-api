//go:build wireinject
// +build wireinject

package main

import (
	"go-restful-api/app"
	"go-restful-api/controller"
	"go-restful-api/middleware"
	"go-restful-api/repository"
	"go-restful-api/service"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/google/wire"
	"github.com/julienschmidt/httprouter"
)

var categorySet = wire.NewSet(
	repository.NewCategoryRepository,
	wire.Bind(new(repository.CategoryRepositoryContract), new(*repository.CategoryRepositoryConcrete)),
	service.NewCategoryService,
	wire.Bind(new(service.CategoryServiceContract), new(*service.CategoryServiceConcrete)),
	controller.NewCategoryController,
	wire.Bind(new(controller.CategoryControllerContract), new(*controller.CategoryControllerConcrete)),
)

func InitializedServer() *http.Server {
	wire.Build(
		app.NewDB,
		validator.New,
		categorySet,
		app.NewRouter,
		wire.Bind(new(http.Handler), new(*httprouter.Router)),
		middleware.NewAuthMiddleware,
		NewServer,
	)
	return nil
}
