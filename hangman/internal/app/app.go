package app

import (
	"github.com/go-openapi/runtime/logger"
	"honnef.co/go/tools/config"
)

type Game struct {
	Log    logger.Logger
	Config *config.Config
}

func NewGame(log logger.Logger, cfg *config.Config) *Game {
	return &Game{
		Log:    log,
		Config: cfg,
	}
}
func (g *Game) Run() {

}
