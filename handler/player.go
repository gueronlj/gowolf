package handler

import (
	"github.com/gueronlj/gowolf/view/player"
	"github.com/labstack/echo/v4"
)

type PlayerHandler struct{}

func (h PlayerHandler) HandlePlayerView(c echo.Context) error {
	return render(c, player.PlayerView("Luis"))
}
