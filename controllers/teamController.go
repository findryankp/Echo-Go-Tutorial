package controllers

import (
	"net/http"
	"soccerApi/models"
	"strconv"

	"github.com/labstack/echo/v4"
)

var teams []models.Team
var transactionPlayers []models.TransactionPlayer

func getIndexTeam(id int) (int, bool) {
	for i, v := range teams {
		if id == v.ID {
			return i, true
		}
	}
	return 0, false
}

// GetAllTeams godoc
// @Summary Get All Teams.
// @Description Get a list of Teams.
// @Tags Teams
// @Produce json
// @Success 200 {object} []models.Team
// @Router /teams [get]
func GetAllTeams(c echo.Context) error {
	return c.JSON(http.StatusOK, map[string]interface{}{
		"status":  true,
		"message": "-",
		"data":    teams,
	})
}

// GetTeam godoc
// @Summary Get Team by ID.
// @Description Get a Team from id.
// @Tags Teams
// @Produce json
// @Param id path string true "id team"
// @Success 200 {object} models.Team
// @Router /teams/{id} [get]
func GetTeam(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	index, flag := getIndexTeam(id)
	if !flag {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"status":  false,
			"message": "Team not found",
			"data":    nil,
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"status":  true,
		"message": "-",
		"data":    teams[index],
	})
}

// CreateTeam godoc
// @Summary Create Team .
// @Description Create Team .
// @Tags Teams
// @Param Body body models.Team true "the body to create a new team"
// @Produce json
// @Success 201 {object} models.Team
// @Router /teams [post]
func CreateTeam(c echo.Context) error {
	team := models.Team{}
	c.Bind(&team)
	team.ID = len(teams) + 1

	teams = append(teams, team)

	return c.JSON(http.StatusCreated, map[string]interface{}{
		"status":  true,
		"message": "-",
		"data":    team,
	})
}

// UpdateTeam godoc
// @Summary Update teams .
// @Description Update teams .
// @Tags Teams
// @Param Body body models.Team true "the body to Update a teams"
// @Param id path string true "id teams"
// @Produce json
// @Success 200 {object} []models.Team
// @Router /teams/{id} [put]
func UpdateTeam(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	index, flag := getIndexTeam(id)
	if !flag {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"status":  false,
			"message": "data not found",
			"data":    nil,
		})
	}

	updateTeam := models.Team{}
	c.Bind(&updateTeam)
	updateTeam.ID = id

	teams[index] = updateTeam
	return c.JSON(http.StatusOK, map[string]interface{}{
		"status":  true,
		"message": "Data Updated Successfully",
		"data":    teams[index],
	})
}

// DeleteTeam godoc
// @Summary Delete teams.
// @Description Delete a teams by id.
// @Tags Teams
// @Produce json
// @Param id path string true "id team"
// @Success 200 {object} map[string]boolean
// @Router /teams/{id} [delete]
func DeleteTeam(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	index, flag := getIndexTeam(id)
	if !flag {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"status":  false,
			"message": "data not found",
			"data":    nil,
		})
	}
	teams = append(teams[:index], teams[index+1:]...)

	return c.JSON(http.StatusOK, map[string]interface{}{
		"status":  true,
		"message": "Data Deleted Successfully",
		"data":    nil,
	})
}

// GetPlayerFromTeam godoc
// @Summary Lits Player From a Team.
// @Description Get a Team from id.
// @Tags Teams
// @Produce json
// @Param id path string true "id team"
// @Success 200 {object} []models.TransactionPlayer
// @Router /teams/{id}/player [get]
func GetPlayerFromTeam(c echo.Context) error {
	idTeam, _ := strconv.Atoi(c.Param("id"))
	indexTeam, flagTeam := getIndexTeam(idTeam)

	if !flagTeam {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"status":  false,
			"message": "Team not found",
			"data":    nil,
		})
	}

	data := map[string]interface{}{
		"team":    teams[indexTeam],
		"players": transactionPlayers,
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"status":  true,
		"message": "-",
		"data":    data,
	})
}

// AddPlayerToTeam godoc
// @Summary Add Player To A Team.
// @Description Add Player To A Team.
// @Tags Teams
// @Produce json
// @Param id path string true "id Team"
// @Param idPlayer path string true "id Player"
// @Success 200 {object} models.Team
// @Router /teams/{id}/player/{idPlayer} [put]
func AddPlayerToTeam(c echo.Context) error {
	idTeam, _ := strconv.Atoi(c.Param("id"))
	idPlayer, _ := strconv.Atoi(c.Param("idPlayer"))

	_, flagTeam := getIndexTeam(idTeam)
	indexPlayer, flagPlayer := getIndexPlayer(idPlayer)

	if !flagTeam {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"status":  false,
			"message": "Team not found",
			"data":    nil,
		})
	} else if !flagPlayer {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"status":  false,
			"message": "Player not found",
			"data":    nil,
		})
	}

	newPlayer := models.TransactionPlayer{
		ID:       len(transactionPlayers) + 1,
		IdTeam:   idTeam,
		IdPlayer: idPlayer,
		Player:   players[indexPlayer],
	}

	transactionPlayers = append(transactionPlayers, newPlayer)

	return c.JSON(http.StatusOK, map[string]interface{}{
		"status":  true,
		"message": "Player Add Successfully",
		"data":    players[indexPlayer],
	})
}

// DeletePlayerFromTeam godoc
// @Summary Delete Player To A Team.
// @Description Delete Player To A Team.
// @Tags Teams
// @Produce json
// @Param id path string true "id Team"
// @Param idPlayer path string true "id Player"
// @Success 200 {object} models.Team
// @Router /teams/{id}/player/{idPlayer} [delete]
func DeletePlayerFromTeam(c echo.Context) error {
	idTeam, _ := strconv.Atoi(c.Param("id"))
	idPlayer, _ := strconv.Atoi(c.Param("idPlayer"))

	_, flagTeam := getIndexTeam(idTeam)
	_, flagPlayer := getIndexPlayer(idPlayer)

	if !flagTeam {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"status":  false,
			"message": "Team not found",
			"data":    nil,
		})
	} else if !flagPlayer {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"status":  false,
			"message": "Player not found",
			"data":    nil,
		})
	}

	for i, v := range transactionPlayers {
		if v.IdTeam == idTeam && v.IdPlayer == idPlayer {
			transactionPlayers = append(transactionPlayers[:i], transactionPlayers[i+1:]...)
		}
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"status":  true,
		"message": "Data Deleted Successfully",
		"data":    nil,
	})
}
