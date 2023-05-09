package controllers

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"

	"ebiznes_go/models"
)

func CreateProductInCart(c echo.Context) error {
	db := c.Get("db").(*gorm.DB)

	cartItem := new(models.Cart)
	if err := c.Bind(cartItem); err != nil {
		return c.JSON(http.StatusBadRequest, "Invalid request payload")
	}
	db.Create(cartItem)

	return c.JSON(http.StatusCreated, cartItem)
}

func GetProductFromCart(c echo.Context) error {
	db := c.Get("db").(*gorm.DB)

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, "Invalid cart record ID")
	}

	cartItem := new(models.Cart)
	res := db.First(cartItem, id)
	if res.Error != nil {
		return c.JSON(http.StatusNotFound, "Product not found in cart")
	}

	return c.JSON(http.StatusOK, cartItem)
}

func UpdateProductQuantityInCart(c echo.Context) error {
	db := c.Get("db").(*gorm.DB)

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, "Invalid cart record ID")
	}

	updatingProd := new(models.Cart)
	if err := c.Bind(updatingProd); err != nil {
		return c.JSON(http.StatusBadRequest, "Invalid request payload")
	}

	cartItemToUpdate := new(models.Cart)
	res := db.First(cartItemToUpdate, id)
	if res.Error != nil {
		return c.JSON(http.StatusNotFound, "Product not found in cart")
	}
	db.Model(cartItemToUpdate).Update("quantity", updatingProd.Quantity)

	return c.JSON(http.StatusOK, cartItemToUpdate)
}

func DeleteProductFromCart(c echo.Context) error {
	db := c.Get("db").(*gorm.DB)

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, "Invalid cart record ID")
	}
	productToBeDeleted := new(models.Cart)
	db.Delete(productToBeDeleted, id)

	return c.NoContent(http.StatusNoContent)
}

func GetAllProductsFromCart(c echo.Context) error {
	db := c.Get("db").(*gorm.DB)

	var productsInCart []*models.Cart
	db.Find(&productsInCart)
	return c.JSON(http.StatusOK, productsInCart)
}

func InitCartController(e *echo.Echo) {
	e.GET("/cart", GetAllProductsFromCart)
	e.GET("/cart/:id", GetProductFromCart)
	e.POST("/cart", CreateProductInCart)
	e.PUT("/cart/:id", UpdateProductQuantityInCart)
	e.DELETE("/cart/:id", DeleteProductFromCart)
}