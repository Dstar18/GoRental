package controller

import (
	"GoRental/model"
	"GoRental/services"
	"net/http"
	"strconv"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

type CarsController struct {
	carsService *services.CarsService
}

func NewCarsController(carsService *services.CarsService) *CarsController {
	return &CarsController{carsService: carsService}
}

func (h *CarsController) CarsGets(c echo.Context) error {
	// gets to service
	cars, err := h.carsService.GetsCars()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"code":    http.StatusInternalServerError,
			"message": err.Error(),
		})
	}

	// return success
	return c.JSON(http.StatusOK, map[string]interface{}{
		"code":    http.StatusOK,
		"message": "Cars successfully",
		"data":    cars,
	})
}

func (h *CarsController) CarsGetID(c echo.Context) error {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"code":    http.StatusBadRequest,
			"message": "Invalid id",
		})
	}

	//  get to service
	cars, err := h.carsService.GetIdCars(uint(id))
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]interface{}{
			"code":    http.StatusNotFound,
			"message": "Data not found",
		})
	}

	// return success
	return c.JSON(http.StatusOK, map[string]interface{}{
		"code":    http.StatusOK,
		"message": "Cars By ID successfully",
		"data":    cars,
	})
}

type CarsValidate struct {
	CarName   string `json:"car_name" validate:"required"`
	DayRate   string `json:"day_rate" validate:"required,numeric"`
	MonthRate string `json:"month_rate" validate:"required,numeric"`
	Image     string `json:"image" validate:"required"`
}

func (h *CarsController) CarsStore(c echo.Context) error {
	var carVal CarsValidate

	// Check request body
	if err := c.Bind(&carVal); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"code":    http.StatusBadRequest,
			"message": "Invalid request body",
		})
	}

	// Validation struct
	validate := validator.New()
	if err := validate.Struct(carVal); err != nil {
		errors := make(map[string]string)
		for _, err := range err.(validator.ValidationErrors) {
			errors[err.Field()] = "This field is" + " " + err.Tag() + " " + err.Param()
		}
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"code":    http.StatusBadRequest,
			"message": errors,
		})
	}

	param := model.Cars{
		CarName:   carVal.CarName,
		DayRate:   carVal.DayRate,
		MonthRate: carVal.MonthRate,
		Image:     carVal.Image,
	}

	// create to service
	if err := h.carsService.CreateCars(&param); err != nil {
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

func (h *CarsController) CarsUpdate(c echo.Context) error {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"code":    http.StatusBadRequest,
			"message": "Invalid id",
		})
	}

	// check data by ID
	_, errCheck := h.carsService.GetIdCars(uint(id))
	if errCheck != nil {
		return c.JSON(http.StatusNotFound, map[string]interface{}{
			"code":    http.StatusNotFound,
			"message": "Data not found",
		})
	}

	var carVal CarsValidate

	// Check request body
	if err := c.Bind(&carVal); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"code":    http.StatusBadRequest,
			"message": "Invalid request body",
		})
	}

	// Validation struct
	validate := validator.New()
	if err := validate.Struct(carVal); err != nil {
		errors := make(map[string]string)
		for _, err := range err.(validator.ValidationErrors) {
			errors[err.Field()] = "This field is" + " " + err.Tag() + " " + err.Param()
		}
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"code":    http.StatusBadRequest,
			"message": errors,
		})
	}

	// Update to service
	param := model.Cars{
		CarID:     uint(id),
		CarName:   carVal.CarName,
		DayRate:   carVal.DayRate,
		MonthRate: carVal.MonthRate,
		Image:     carVal.Image,
	}
	if err := h.carsService.UpdateCars(&param); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"code":    http.StatusInternalServerError,
			"message": err.Error(),
		})
	}

	// return success
	return c.JSON(http.StatusOK, map[string]interface{}{
		"code":    http.StatusOK,
		"message": "Update successfully",
		"data":    nil,
	})
}

func (h *CarsController) CarsDelete(c echo.Context) error {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"code":    http.StatusBadRequest,
			"message": "Invalid id",
		})
	}

	// check data by ID
	_, errCheck := h.carsService.GetIdCars(uint(id))
	if errCheck != nil {
		return c.JSON(http.StatusNotFound, map[string]interface{}{
			"code":    http.StatusNotFound,
			"message": "Data not found",
		})
	}

	// delete to service
	if err := h.carsService.DeleteCars(uint(id)); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"code":    http.StatusInternalServerError,
			"message": err.Error(),
		})
	}

	// return success
	return c.JSON(http.StatusOK, map[string]interface{}{
		"code":    http.StatusOK,
		"message": "Delete successfully",
		"data":    nil,
	})
}
