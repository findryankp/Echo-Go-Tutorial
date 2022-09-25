package main

import (
	"fmt"
	"log"
	"os"
	"soccerApi/routes"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	e := echo.New()
	routes.Routes(e)

	e.Start(fmt.Sprintf(":%v", os.Getenv("PORT")))
}
