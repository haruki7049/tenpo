package playing

import (
	"errors"
	"fmt"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
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
		pack.Note{Button: ebiten.MouseButtonLeft, Timing: 60},
		pack.Note{Button: ebiten.MouseButtonLeft, Timing: 120},
		pack.Note{Button: ebiten.MouseButtonLeft, Timing: 180},
		pack.Note{Button: ebiten.MouseButtonLeft, Timing: 240},
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
	var isJustTiming bool = false
	tick := es.tick.Count()

	for _, n := range p.Sheets {
		mouseButton := n.Button
		maxTiming := n.Timing + TimingRange
		minTiming := n.Timing - TimingRange

		if maxTiming >= tick && minTiming <= tick && inpututil.IsMouseButtonJustPressed(mouseButton) {
			isJustTiming = true
		}
	}

	if isJustTiming {
		es.score += 1
	}
}

type tickCounter struct {
	count int
}

func newTickCounter() *tickCounter {
	tc := &tickCounter{}
	tc.count = 0

	return tc
}

func (tc *tickCounter) Count() int {
	return tc.count
}

func (tc *tickCounter) Update() {
	tc.count += 1
}

func (tc *tickCounter) UpdateWith(c int) {
	tc.count += c
}
