package handler

import (
	"github.com/gueronlj/gowolf/model"
	"github.com/gueronlj/gowolf/view/player"
	"github.com/labstack/echo/v4"
)

type PlayerHandler struct{}

func (h PlayerHandler) HandlePlayerView(c echo.Context) error {
	p := model.Player{
		Name: "Luis",
		Role: "Villager",
	}
	return render(c, player.PlayerView(p))
}
