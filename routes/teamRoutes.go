package routes

import (
	"soccerApi/controllers"

	"github.com/labstack/echo/v4"
)

func TeamRoutes(e *echo.Echo) {
	g := e.Group("/teams")
	g.GET("", controllers.GetAllTeams)
	g.GET("/:id", controllers.GetTeam)
	g.POST("", controllers.CreateTeam)
	g.PUT("/:id", controllers.UpdateTeam)
	g.DELETE("/:id", controllers.DeleteTeam)

	g.GET("/:id/player", controllers.GetPlayerFromTeam)
	g.PUT("/:id/player/:idPlayer", controllers.AddPlayerToTeam)
	g.DELETE("/:id/player/:idPlayer", controllers.DeletePlayerFromTeam)
}
