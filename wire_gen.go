// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

import (
	"github.com/go-playground/validator/v10"
	"github.com/google/wire"
	"go-restful-api/app"
	"go-restful-api/controller"
	"go-restful-api/middleware"
	"go-restful-api/repository"
	"go-restful-api/service"
	"net/http"
)

// Injectors from injector.go:

func InitializedServer() *http.Server {
	categoryRepositoryConcrete := repository.NewCategoryRepository()
	db := app.NewDB()
	validate := validator.New()
	categoryServiceConcrete := service.NewCategoryService(categoryRepositoryConcrete, db, validate)
	categoryControllerConcrete := controller.NewCategoryController(categoryServiceConcrete)
	router := app.NewRouter(categoryControllerConcrete)
	authMiddleware := middleware.NewAuthMiddleware(router)
	server := NewServer(authMiddleware)
	return server
}

// injector.go:

var categorySet = wire.NewSet(repository.NewCategoryRepository, wire.Bind(new(repository.CategoryRepositoryContract), new(*repository.CategoryRepositoryConcrete)), service.NewCategoryService, wire.Bind(new(service.CategoryServiceContract), new(*service.CategoryServiceConcrete)), controller.NewCategoryController, wire.Bind(new(controller.CategoryControllerContract), new(*controller.CategoryControllerConcrete)))