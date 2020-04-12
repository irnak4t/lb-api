package main

import (
	"net/http"

	"github.com/irnak4t/lb-api/db/mysql"
	"github.com/irnak4t/lb-api/models"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/:runner", getRuns)

	e.Logger.Fatal(e.Start(":3000"))
}

func getRuns(c echo.Context) error {
	db, _ := mysql.Open()
	defer db.Close()

	var records []models.Record
	db.Where(&models.Record{Runner: c.Param("runner")}).Find(&records)
	return c.JSON(http.StatusOK, records)
}
