package main

import (
	"net/http"

	"github.com/irnak4t/lb-api/models"
	"github.com/irnak4t/leaderboards/db/mysql"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

var e *echo.Echo

func main() {
	echoInit()
	echoRouting()
	echoRun()
}

func echoInit() {
	e = echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
}

func echoRouting() {
	e.GET("/:runner", getRuns)
}

func echoRun() {
	e.Logger.Fatal(e.Start(":3000"))
}

func getRuns(c echo.Context) error {
	db := mysql.Open()
	defer db.Close()

	var records []models.Record
	db.Where(&models.Record{Runner: c.Param("runner")}).Find(&records)
	return c.JSON(http.StatusOK, records)
}
