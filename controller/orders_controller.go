package controller

import (
	"GoRental/services"
	"net/http"
	"strconv"

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
