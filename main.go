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
	orderService := services.NewOrdersService(dbConn)

	// Init Controller
	carsController := controller.NewCarsController(carsService)
	ordersController := controller.NewOrdersController(orderService)

	// Routes
	e := echo.New()
	api := e.Group("/api")

	// Route cars
	api.GET("/cars", carsController.CarsGets)
	api.GET("/cars/:id", carsController.CarsGetID)
	api.POST("/cars", carsController.CarsStore)
	api.POST("/cars/:id", carsController.CarsUpdate)
	api.DELETE("/cars/:id", carsController.CarsDelete)

	// Route orders
	api.GET("/orders", ordersController.OrdersGets)
	api.GET("/orders/:id", ordersController.OrdersGetID)

	e.Logger.Fatal(e.Start(":3000"))
}
