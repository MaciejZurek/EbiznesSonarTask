package main

import (
	"fmt"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"

	"ebiznes_go/controllers"
	"ebiznes_go/models"
)

func initDatabase(db *gorm.DB, e *echo.Echo) {
	var err error
	db, err = gorm.Open(sqlite.Open("testEbiznes.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	db.AutoMigrate(&models.Product{}, &models.Cart{}, &models.Category{})
	e.Use(func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			c.Set("db", db)
			return next(c)
		}
	})
}

func main() {

	fmt.Println("APPLICATION STARTS")
	e := echo.New()
	db := new(gorm.DB)

	initDatabase(db, e)

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
        AllowOrigins: []string{"http://localhost:3000"},
        AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
    }))

	controllers.InitProductController(e)
	controllers.InitCartController(e)
	controllers.InitCategoryController(e)
	controllers.InitPaymentController(e)

	// Start server
	e.Logger.Fatal(e.Start(":1323"))
}
