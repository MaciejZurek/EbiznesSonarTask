package controllers

import (
	"net/http"
	"strconv"

	"ebiznes_go/models"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func CreateCategory(c echo.Context) error {
	db := c.Get("db").(*gorm.DB)
	
	cat := new(models.Category)
	if err := c.Bind(cat); err != nil {
		return c.JSON(http.StatusBadRequest, "Invalid request payload")
	}
	db.Create(cat)

	return c.JSON(http.StatusCreated, cat)
}

func GetCategory(c echo.Context) error {
	db := c.Get("db").(*gorm.DB)
	
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, "Invalid category ID")
	}

	cat := new(models.Category)
	res := db.First(cat, id)
	if res.Error != nil {
		return c.JSON(http.StatusNotFound, "Category not found")
	}

	return c.JSON(http.StatusOK, cat)
}

func UpdateCategoryDescription(c echo.Context) error {
	db := c.Get("db").(*gorm.DB)
	
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, "Invalid category ID")
	}

	updatingCategory := new(models.Category)
	if err := c.Bind(updatingCategory); err != nil {
		return c.JSON(http.StatusBadRequest, "Invalid request payload")
	}

	categoryToUpdate := new(models.Category)
	res := db.First(categoryToUpdate, id)
	if res.Error != nil {
		return c.JSON(http.StatusNotFound, "Category not found")
	}
	db.Model(categoryToUpdate).Update("description", updatingCategory.Description)

	return c.JSON(http.StatusOK, categoryToUpdate)
}

func DeleteCategory(c echo.Context) error {
	db := c.Get("db").(*gorm.DB)
	
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, "Invalid category ID")
	}
	cat := new(models.Category)
	db.Delete(cat, id)

	return c.NoContent(http.StatusNoContent)
}

func GetAllCategories(c echo.Context) error {
	db := c.Get("db").(*gorm.DB)
	
	var categories []*models.Category
	db.Find(&categories)
	return c.JSON(http.StatusOK, categories)
}

func InitCategoryController(e *echo.Echo) {
	e.GET("/categories", GetAllCategories)
	e.GET("/categories/:id", GetCategory)
	e.POST("/categories", CreateCategory)
	e.PUT("/categories/:id", UpdateCategoryDescription)
	e.DELETE("/categories/:id", DeleteCategory)
}