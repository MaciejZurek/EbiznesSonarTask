package controllers

import (
	"net/http"
	"strconv"

	"ebiznes_go/models"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)



func CreateProduct(c echo.Context) error {
	db := c.Get("db").(*gorm.DB)
	
	p := new(models.Product)
	if err := c.Bind(p); err != nil {
		return c.JSON(http.StatusBadRequest, "Invalid request payload")
	}
	db.Create(p)

	return c.JSON(http.StatusCreated, p)
}

func GetProduct(c echo.Context) error {
	db := c.Get("db").(*gorm.DB)
	
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, "Invalid product ID")
	}

	p := new(models.Product)
	res := db.First(p, id)
	if res.Error != nil {
		return c.JSON(http.StatusNotFound, "Product not found")
	}

	return c.JSON(http.StatusOK, p)
}

func UpdateProductPrice(c echo.Context) error {
	db := c.Get("db").(*gorm.DB)
	
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, "Invalid product ID")
	}

	updatingProd := new(models.Product)
	if err := c.Bind(updatingProd); err != nil {
		return c.JSON(http.StatusBadRequest, "Invalid request payload")
	}

	pToUpdate := new(models.Product)
	res := db.First(pToUpdate, id)
	if res.Error != nil {
		return c.JSON(http.StatusNotFound, "Product not found")
	}
	db.Model(pToUpdate).Update("price", updatingProd.Price)

	return c.JSON(http.StatusOK, pToUpdate)
}

func DeleteProduct(c echo.Context) error {
	db := c.Get("db").(*gorm.DB)
	
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, "Invalid product ID")
	}
	p := new(models.Product)
	db.Delete(p, id)

	return c.NoContent(http.StatusNoContent)
}

func GetAllProducts(c echo.Context) error {
	db := c.Get("db").(*gorm.DB)
	
	var products []*models.Product
	db.Find(&products)
	return c.JSON(http.StatusOK, products)
}

func InitProductController(e *echo.Echo) {
	e.GET("/products", GetAllProducts)
	e.GET("/products/:id", GetProduct)
	e.POST("/products", CreateProduct)
	e.PUT("/products/:id", UpdateProductPrice)
	e.DELETE("/products/:id", DeleteProduct)
}