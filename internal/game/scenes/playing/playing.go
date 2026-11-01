package playing

import (
	"errors"
	"fmt"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/haruki7049/tenpo/internal/game/configs"
	"github.com/haruki7049/tenpo/internal/pack"
	"github.com/haruki7049/tenpo/internal/scene"
)

const (
	TrailingTicks     = 60 * 2
	TimingRange   int = 8
)

type PlayingScene struct {
	score int
	tick  *tickCounter
}

var _ scene.Scene = (*PlayingScene)(nil)

// default score pack
var defaultSheet pack.SheetPack = pack.SheetPack{
	Sheets: []pack.Note{
		pack.Note{Button: ebiten.MouseButtonLeft, Timing: configs.TPS * 0},
		pack.Note{Button: ebiten.MouseButtonLeft, Timing: configs.TPS * 1},
		pack.Note{Button: ebiten.MouseButtonLeft, Timing: configs.TPS * 2},
		pack.Note{Button: ebiten.MouseButtonLeft, Timing: configs.TPS * 3},
		pack.Note{Button: ebiten.MouseButtonLeft, Timing: configs.TPS * 4},
	},
	Source: "",
}

func NewPlayingScene() *PlayingScene {
	g := &PlayingScene{}
	g.tick = newTickCounter()

	return g
}

func (es *PlayingScene) Draw(screen *ebiten.Image) {
	message := fmt.Sprintf("es.score: %d, es.tick: %d", es.score, es.tick.Count())

	ebitenutil.DebugPrint(screen, "In PlayingScene\n"+message)
}

func (es *PlayingScene) Update() (scene.Scene, error) {
	es.tick.Update()
	es.judge(defaultSheet)

	if defaultSheet.MaxTick()+TrailingTicks < es.tick.Count() {
		return nil, errors.New("Finished the defaultScore")
	}

	return nil, nil
}

func (es *PlayingScene) judge(p pack.SheetPack) {
	tick := es.tick.Count()

	for i, n := range p.Sheets {
		mouseButton := n.Button
		maxTiming := n.Timing + TimingRange
		minTiming := n.Timing - TimingRange

		if maxTiming >= tick && minTiming <= tick && inpututil.IsMouseButtonJustPressed(mouseButton) {
			es.score += 1
			p.Sheets[i].MarkAsJudged()
			break
		}
	}
}
