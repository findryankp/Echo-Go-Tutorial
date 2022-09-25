package routes

import (
	"fmt"
	"net/http"
	"os"
	"soccerApi/docs"

	"github.com/labstack/echo/v4"
	echoSwagger "github.com/swaggo/echo-swagger"
)

func Routes(e *echo.Echo) {
	docs.SwaggerInfo.Title = "Soccer API - Problem2"
	docs.SwaggerInfo.Description = "Kitabisa.com"
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.Host = fmt.Sprintf("localhost:%v", os.Getenv("PORT"))
	docs.SwaggerInfo.Schemes = []string{"http", "https"}

	e.GET("/docs/*", echoSwagger.WrapHandler)
	e.GET("/", index)
	TeamRoutes(e)
	PlayerRoutes(e)
}

func index(c echo.Context) error {
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message":       "Selamat datang di Soccer API",
		"documentation": fmt.Sprintf("localhost:%v/docs/index.html", os.Getenv("PORT")),
	})
}
