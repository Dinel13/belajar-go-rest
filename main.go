package main

import (
	"github.com/go-playground/validator/v10"
	"github.com/julienschmidt/httprouter"
	"go-rest-api/app"
	"go-rest-api/controller"
	"go-rest-api/repository"
	"go-rest-api/service"
)

func main() {

	validate := validator.New()
	db := app.NewDB()

	categoryRepository :=  repository.NewCategoryRepository()
	categoryService := service.NewCategoryService(categoryRepository, db, validate)
	categoryController := controller.NewCategoryController(categoryService)


	router := httprouter.New()

	router.GET("/api/categories", categoryController.FindAll)
	router.GET("/api/categories/:categoryId", categoryController.FindById)
	router.POST("/api/categories", categoryController.Create)
	router.PUT("/api/categories/:categoryId", categoryController.Update)
	router.DELETE("/api/categories/:categoryId", categoryController.Delete)
}

