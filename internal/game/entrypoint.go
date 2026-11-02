package game

import (
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/haruki7049/tenpo/internal/game/configs"
)

func Run() {
	ebiten.SetWindowSize(configs.WINDOW_WIDTH, configs.WINDOW_HEIGHT)
	ebiten.SetWindowTitle(configs.WINDOW_TITLE)
	ebiten.SetTPS(configs.TPS)

	var game *Game
	var err error
	game, err = newGame()
	if err != nil {
		log.Fatal(err)
	}

	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
