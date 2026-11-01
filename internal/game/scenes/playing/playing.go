package playing

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/audio"
	"github.com/hajimehoshi/ebiten/v2/audio/wav"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/haruki7049/tenpo/assets"
	"github.com/haruki7049/tenpo/internal/game/configs"
	"github.com/haruki7049/tenpo/internal/pack"
	"github.com/haruki7049/tenpo/internal/scene"
)

const (
	TrailingTicks     = 60 * 2
	TimingRange   int = 8
)

type PlayingScene struct {
	pack  pack.SheetPack
	score int
	tick  *tickCounter
}

type player struct {
	audioContext *audio.Context
	audioPlayer  *audio.Player
	current      time.Duration
	total        time.Duration
}

func newPlayer(audioContext *audio.Context) (*player, error) {
	type audioStream interface {
		io.ReadSeeker
		Length() int64
	}

	const sampleRate = 48000

	// bytesPerSample is the byte size for one sample (8 [bytes] = 2 [channels] * 4 [bytes] (32bit float)).
	const bytesPerSample = 8

	f, err := assets.Assets.Open("audio/defaultSource.wav")
	if err != nil {
		return nil, err
	}

	// Read defaultSource.wav bytes
	data, err := io.ReadAll(f)
	if err != nil {
		return nil, err
	}

	s, err := wav.DecodeF32(bytes.NewReader(data))
	if err != nil {
		return nil, err
	}

	// Create *audio.Player
	p, err := audioContext.NewPlayerF32(s)
	if err != nil {
		return nil, err
	}

	// Define player
	player := &player{
		audioContext: audioContext,
		audioPlayer:  p,
		total:        time.Second * time.Duration(s.Length()) / bytesPerSample / sampleRate,
	}

	player.audioPlayer.Play()
	return player, nil
}

var _ scene.Scene = (*PlayingScene)(nil)

// default score pack
var defaultSheet pack.SheetPack = pack.SheetPack{
	AudioSource: "audio/defaultSource.wav",
	Sheets: []pack.Note{
		pack.Note{Button: ebiten.MouseButtonLeft, Timing: configs.TPS * 0},
		pack.Note{Button: ebiten.MouseButtonLeft, Timing: configs.TPS * 1},
		pack.Note{Button: ebiten.MouseButtonLeft, Timing: configs.TPS * 2},
		pack.Note{Button: ebiten.MouseButtonLeft, Timing: configs.TPS * 3},
		pack.Note{Button: ebiten.MouseButtonLeft, Timing: configs.TPS * 4},
	},
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
