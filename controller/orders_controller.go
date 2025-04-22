package controller

import (
	"GoRental/model"
	"GoRental/services"
	"net/http"
	"strconv"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

type OrdersController struct {
	ordersService *services.OrdersService
}

func NewOrdersController(ordersService *services.OrdersService) *OrdersController {
	return &OrdersController{ordersService: ordersService}
}

func (h *OrdersController) OrdersGets(c echo.Context) error {
	// gets to service
	orders, err := h.ordersService.GetsOrders()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"code":    http.StatusInternalServerError,
			"message": err.Error(),
		})
	}

	// return success
	return c.JSON(http.StatusOK, map[string]interface{}{
		"code":    http.StatusOK,
		"message": "Orders successfully",
		"data":    orders,
	})
}

func (h *OrdersController) OrdersGetID(c echo.Context) error {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"code":    http.StatusBadRequest,
			"message": "Invalid id",
		})
	}

	//  get to service
	orders, err := h.ordersService.GetIdOrders(uint(id))
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]interface{}{
			"code":    http.StatusNotFound,
			"message": "Data not found",
		})
	}

	// return success
	return c.JSON(http.StatusOK, map[string]interface{}{
		"code":    http.StatusOK,
		"message": "Order By ID successfully",
		"data":    orders,
	})
}

type OrdersValidate struct {
	IDCar           int       `json:"id_car" validate:"required"`
	OrderDate       time.Time `json:"order_date" validate:"required"`
	PickupDate      time.Time `json:"pickup_date" validate:"required"`
	DropOffDate     time.Time `json:"dropoff_date" validate:"required"`
	PickupLocation  string    `json:"pickup_location" validate:"required"`
	DropOffLocation string    `json:"dropoff_location" validate:"required"`
}

func (h *OrdersController) OrdersStore(c echo.Context) error {
	var orderVal OrdersValidate

	// Check request body
	if err := c.Bind(&orderVal); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"code":    http.StatusBadRequest,
			"message": "Invalid request body",
		})
	}

	// Validation struct
	validate := validator.New()
	if err := validate.Struct(orderVal); err != nil {
		errors := make(map[string]string)
		for _, err := range err.(validator.ValidationErrors) {
			errors[err.Field()] = "This field is" + " " + err.Tag() + " " + err.Param()
		}
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"code":    http.StatusBadRequest,
			"message": errors,
		})
	}

	param := model.Orders{
		IDCar:           orderVal.IDCar,
		OrderDate:       orderVal.OrderDate,
		PickupDate:      orderVal.PickupDate,
		DropOffDate:     orderVal.DropOffDate,
		PickupLocation:  orderVal.PickupLocation,
		DropOffLocation: orderVal.DropOffLocation,
	}
	// create to service
	if err := h.ordersService.CreateOrders(&param); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"code":    http.StatusInternalServerError,
			"message": err.Error(),
		})
	}

	// return success
	return c.JSON(http.StatusOK, map[string]interface{}{
		"code":    http.StatusOK,
		"message": "Created successfully",
		"data":    nil,
	})
}
