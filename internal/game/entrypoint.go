package game

import (
	"log"

	"github.com/hajimehoshi/ebiten/v2"
)

func Run() {
	ebiten.SetWindowSize(WINDOW_WIDTH, WINDOW_HEIGHT)
	ebiten.SetWindowTitle(WINDOW_TITLE)

	game := NewGame()
	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
