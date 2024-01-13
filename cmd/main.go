package main

import (
	"github.com/gueronlj/gowolf/handler"
	"github.com/labstack/echo/v4"
)

func main() {
	app := echo.New()

	playerHandler := handler.PlayerHandler{ /*insert DB struct here*/ }
	app.GET("/", playerHandler.HandlePlayerView)

	app.Logger.Fatal(
		app.Start(":8080"),
	)
}
