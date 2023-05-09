package controllers

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

func CreatePayment(c echo.Context) error {

	fmt.Println("PAYMENT DATA RECEIVED")

	return c.JSON(http.StatusCreated, "")
}

func InitPaymentController(e *echo.Echo) {
	e.POST("/payment", CreatePayment)
}