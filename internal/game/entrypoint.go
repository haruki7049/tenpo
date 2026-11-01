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

	game := newGame()
	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
