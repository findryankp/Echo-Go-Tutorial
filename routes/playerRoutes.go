package routes

import (
	"soccerApi/controllers"

	"github.com/labstack/echo/v4"
)

func PlayerRoutes(e *echo.Echo) {
	p := e.Group("/players")
	p.GET("", controllers.GetAllPlayers)
	p.GET("/:id", controllers.GetPlayer)
	p.POST("", controllers.CreatePlayer)
	p.PUT("/:id", controllers.UpdatePlayer)
	p.DELETE("/:id", controllers.DeletePlayer)
}
