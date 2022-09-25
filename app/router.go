package app

import (
	"go-restful-api/controller"
	"go-restful-api/exception"

	"github.com/julienschmidt/httprouter"
)

func NewRouter(categoryController controller.CategoryControllerContract) *httprouter.Router {
	router := httprouter.New()

	router.PanicHandler = exception.ErrorHandler

	router.GET("/api/categories", categoryController.FindAll)
	router.GET("/api/categories/:categoryId", categoryController.FindById)
	router.POST("/api/categories", categoryController.Create)
	router.PUT("/api/categories/:categoryId", categoryController.Update)
	router.DELETE("/api/categories/:categoryId", categoryController.Delete)

	return router
}
