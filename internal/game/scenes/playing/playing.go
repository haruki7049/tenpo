package playing

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/haruki7049/tenpo/internal/scene"
)

type PlayingScene struct{}

var _ scene.Scene = (*PlayingScene)(nil)

func NewPlayingScene() *PlayingScene {
	return &PlayingScene{}
}

func (es *PlayingScene) Draw(screen *ebiten.Image) {
	ebitenutil.DebugPrint(screen, "In PlayingScene")
}

func (es *PlayingScene) Update() (scene.Scene, error) {
	return nil, nil
}
