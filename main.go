package main

import (
	"GoRental/config"
	"GoRental/controller"
	"GoRental/services"

	"github.com/labstack/echo/v4"
)

func main() {
	config.LoadEnv()
	dbConn := config.ConnectDB()

	// Init Service
	carsService := services.NewCarsService(dbConn)
	// Init Controller
	carsController := controller.NewCarsController(carsService)

	// Routes
	e := echo.New()
	api := e.Group("/api")
	api.GET("/cars", carsController.CarsGets)
	api.GET("/cars/:id", carsController.CarsGetID)
	api.POST("/cars", carsController.CarsStore)
	api.POST("/cars/:id", carsController.CarsUpdate)
	api.DELETE("/cars/:id", carsController.CarsDelete)

	e.Logger.Fatal(e.Start(":3000"))
}
