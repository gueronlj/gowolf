package handler

import (
	player "github.com/gueronlj/gowolf/view"
	"github.com/labstack/echo/v4"
)

type PlayerHandler struct{}

func (h PlayerHandler) HandlePlayerView(c echo.Context) error {
	return render(c, player.PlayerView("Player 1"))
}
