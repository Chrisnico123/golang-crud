package main

import (
	"golang_api/app"
	"golang_api/controller"
	"golang_api/exception"
	"golang_api/helper"
	"golang_api/repository"
	"golang_api/service"
	"net/http"

	_ "github.com/go-sql-driver/mysql"

	"github.com/go-playground/validator/v10"
	"github.com/julienschmidt/httprouter"
)

func main() {
	db := app.NewDb()
	validate := validator.New()
	categoryRepository := repository.NewCategoryRepository()
	categoryService := service.NewCategoryService(categoryRepository , db , validate)
	categoryController := controller.NewCategoryController(categoryService)
	
	router := httprouter.New()
	router.GET("/api/categories" , categoryController.FindAll)
	router.GET("/api/categories/:categoryId" , categoryController.FindById)
	router.POST("/api/categories" , categoryController.Create)
	router.DELETE("/api/categories/:categoryId" , categoryController.Delete)
	router.PUT("/api/categories/:categoryId" , categoryController.Update)

	router.PanicHandler = exception.ErrorHandler

	server := http.Server{
		Addr: "localhost:3000",
		Handler: router,
	}

	err := server.ListenAndServe()
	helper.PanicIfError(err)

}