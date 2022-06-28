package main

import (
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/julienschmidt/httprouter"
	"github.com/tomiprasetyo/belajar-restful-api-bunker-service/app"
	"github.com/tomiprasetyo/belajar-restful-api-bunker-service/controller"
	"github.com/tomiprasetyo/belajar-restful-api-bunker-service/exception"
	"github.com/tomiprasetyo/belajar-restful-api-bunker-service/helper"
	"github.com/tomiprasetyo/belajar-restful-api-bunker-service/middleware"
	"github.com/tomiprasetyo/belajar-restful-api-bunker-service/repository"
	"github.com/tomiprasetyo/belajar-restful-api-bunker-service/service"
)

func main() {

	// koneksi ke database
	db := app.NewDB()

	// validasi data
	validate := validator.New()

	bunkerServiceRepository := repository.NewBunkerServiceRepository()
	bunkerServiceServe := service.NewBunkerServiceServe(bunkerServiceRepository, db, validate)
	bunkerServiceController := controller.NewBunkerServiceController(bunkerServiceServe)

	// buat router dari httprouter
	router := httprouter.New()

	// API ambil semua data
	router.GET("/api/bunker-services", bunkerServiceController.FindAll)

	// API ambil data berdasarkan id
	router.GET("/api/bunker-services/:bunkerServiceId", bunkerServiceController.FindById)

	// API membuat data
	router.POST("/api/bunker-services", bunkerServiceController.Create)

	// API mengupdate/merubah data
	router.PUT("/api/bunker-services/:bunkerServiceId", bunkerServiceController.Update)

	// API menghapus data
	router.DELETE("/api/bunker-services/:bunkerServiceId", bunkerServiceController.Delete)

	// set standar error http
	router.PanicHandler = exception.ErrorHandler

	// buat server
	server := http.Server{
		// set address
		Addr: "localhost:8095",
		// set handler
		Handler: middleware.NewAuthMiddleware(router),
	}

	err := server.ListenAndServe()
	helper.PanicIfError(err)
}
