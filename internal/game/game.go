package game

import (
	"errors"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/haruki7049/tenpo/internal/game/scenes/playing"
	"github.com/haruki7049/tenpo/internal/scene"
)

// Game is the root ebiten.Game implementation. It only owns the currently
// active Scene and delegates Update/Draw to it.
type Game struct {
	scene scene.Scene
}

func (g *Game) Update() error {
	if g.scene == nil {
		return errors.New("Failed to use g.scene because g.scene is null")
	}

	next, err := g.scene.Update()
	if err != nil {
		return err
	}
	if next != nil {
		g.scene = next
	}

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	if g.scene == nil {
		return
	}

	g.scene.Draw(screen)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return outsideWidth, outsideHeight
}

// Initialize a value typed Game struct
func newGame() (*Game, error) {
	g := &Game{}
	scene, err := playing.NewPlayingScene()
	if err != nil {
		return nil, err
	}

	g.scene = scene
	return g, err
}
