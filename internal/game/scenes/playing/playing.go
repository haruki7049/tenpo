package playing

import (
	"fmt"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/haruki7049/tenpo/internal/scene"
)

type PlayingScene struct {
	count int
}

var _ scene.Scene = (*PlayingScene)(nil)

func NewPlayingScene() *PlayingScene {
	return &PlayingScene{}
}

func (es *PlayingScene) Draw(screen *ebiten.Image) {
	message := fmt.Sprintf("es.count: %d", es.count)

	ebitenutil.DebugPrint(screen, "In PlayingScene\n"+message)
}

func (es *PlayingScene) Update() (scene.Scene, error) {
	if inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) {
		es.count += 1
	}

	return nil, nil
}
