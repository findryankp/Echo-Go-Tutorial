package controllers

import (
	"net/http"
	"soccerApi/models"
	"strconv"

	"github.com/labstack/echo/v4"
)

var players []models.Player

func getIndexPlayer(id int) (int, bool) {
	for i, v := range teams {
		if id == v.ID {
			return i, true
		}
	}
	return 0, false
}

// GetAllPlayers godoc
// @Summary Get All Players.
// @Description Get a list of Players.
// @Tags Players
// @Produce json
// @Success 200 {object} []models.Player
// @Router /players [get]
func GetAllPlayers(c echo.Context) error {
	return c.JSON(http.StatusOK, map[string]interface{}{
		"status":  true,
		"message": "-",
		"data":    players,
	})
}

// GetPlayer godoc
// @Summary Get Player by ID.
// @Description Get a Player from id.
// @Tags Players
// @Produce json
// @Param id path string true "id player"
// @Success 200 {object} models.Player
// @Router /players/{id} [get]
func GetPlayer(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	index, flag := getIndexPlayer(id)
	if !flag {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"status":  false,
			"message": "data not found",
			"data":    nil,
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"status":  true,
		"message": "-",
		"data":    players[index],
	})
}

// CreatePlayer godoc
// @Summary Create Player .
// @Description Create Player .
// @Tags Players
// @Param Body body models.Player true "the body to create a new Player"
// @Produce json
// @Success 201 {object} models.Player
// @Router /Players [post]
func CreatePlayer(c echo.Context) error {
	player := models.Player{}
	c.Bind(&player)
	player.ID = len(players) + 1

	players = append(players, player)

	return c.JSON(http.StatusCreated, map[string]interface{}{
		"status":  true,
		"message": "-",
		"data":    player,
	})
}

// UpdatePlayer godoc
// @Summary Update players .
// @Description Update players .
// @Tags Players
// @Param Body body models.Player true "the body to Update a players"
// @Param id path string true "id players"
// @Produce json
// @Success 200 {object} []models.Player
// @Router /players/{id} [put]
func UpdatePlayer(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	index, flag := getIndexTeam(id)
	if !flag {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"status":  false,
			"message": "data not found",
			"data":    nil,
		})
	}

	updatePlayer := models.Player{}
	c.Bind(&updatePlayer)
	updatePlayer.ID = id

	players[index] = updatePlayer
	return c.JSON(http.StatusOK, map[string]interface{}{
		"status":  true,
		"message": "Data Updated Successfully",
		"data":    players[index],
	})
}

// DeletePlayer godoc
// @Summary Delete players.
// @Description Delete a players by id.
// @Tags Players
// @Produce json
// @Param id path string true "id players"
// @Success 200 {object} map[string]boolean
// @Router /players/{id} [delete]
func DeletePlayer(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	index, flag := getIndexTeam(id)
	if !flag {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"status":  false,
			"message": "data not found",
			"data":    nil,
		})
	}
	players = append(players[:index], players[index+1:]...)

	return c.JSON(http.StatusOK, map[string]interface{}{
		"status":  true,
		"message": "Data Deleted Successfully",
		"data":    nil,
	})
}
