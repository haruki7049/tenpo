package example

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/haruki7049/tenpo/internal/scene"
)

type ExampleScene struct{}

func NewExampleScene() *ExampleScene {
	return &ExampleScene{}
}

func (es *ExampleScene) Draw(screen *ebiten.Image) {
	ebitenutil.DebugPrint(screen, "Hello, tenpo game!!")
}

func (es *ExampleScene) Update() (scene.Scene, error) {
	return nil, nil
}

var _ scene.Scene = (*ExampleScene)(nil)
