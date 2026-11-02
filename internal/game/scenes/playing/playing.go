package playing

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"io/fs"

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

var (
	defaultSourceFile fs.File
)

const (
	TrailingTicks     = 60 * 2
	TimingRange   int = 8
	mouseButton       = ebiten.MouseButtonLeft
)

type PlayingScene struct {
	audioContext *audio.Context
	audioBytes   []byte

	pack         pack.SheetPack
	score        int
	tick         *tickCounter
	inputHistory inputHistory
}

var _ scene.Scene = (*PlayingScene)(nil)

// default score pack
var defaultSheet pack.SheetPack = pack.SheetPack{
	AudioSource: "audio/defaultSource.wav",
	Sheets: []pack.Note{
		{Button: ebiten.MouseButtonLeft, Timing: configs.TPS * 0},
		{Button: ebiten.MouseButtonLeft, Timing: configs.TPS * 1},
		{Button: ebiten.MouseButtonLeft, Timing: configs.TPS * 2},
		{Button: ebiten.MouseButtonLeft, Timing: configs.TPS * 3},
		{Button: ebiten.MouseButtonLeft, Timing: configs.TPS * 4},
	},
}

func NewPlayingScene() (*PlayingScene, error) {
	const sampleRate = 48000

	ps := &PlayingScene{}
	ps.tick = newTickCounter()
	ps.audioContext = audio.NewContext(sampleRate)

	type audioStream interface {
		io.ReadSeeker
		Length() int64
	}

	// bytesPerSample is the byte size for one sample (8 [bytes] = 2 [channels] * 4 [bytes] (32bit float)).
	const bytesPerSample = 8

	var err error
	var s audioStream

	defaultSourceFile, err = assets.Assets.Open("audio/defaultSource.wav")
	if err != nil {
		return nil, err
	}

	// Read defaultSource.wav bytes
	data, err := io.ReadAll(defaultSourceFile)
	if err != nil {
		return nil, err
	}

	s, err = wav.DecodeWithoutResampling(bytes.NewReader(data))
	if err != nil {
		return nil, err
	}

	ps.audioBytes, err = io.ReadAll(s)
	if err != nil {
		return nil, err
	}

	player := ps.audioContext.NewPlayerFromBytes(ps.audioBytes)
	player.Play()

	return ps, nil
}

func (ps *PlayingScene) Draw(screen *ebiten.Image) {
	message := fmt.Sprintf("ps.score: %d, ps.tick: %d\nps.inputHistory: %v", ps.score, ps.tick.Count(), ps.inputHistory.inner)

	ebitenutil.DebugPrint(screen, "In PlayingScene\n"+message)
}

func (ps *PlayingScene) Update() (scene.Scene, error) {
	ps.tick.Update()
	ps.judge(defaultSheet)

	if defaultSheet.MaxTick()+TrailingTicks < ps.tick.Count() {
		return nil, errors.New("Finished the defaultScore")
	}

	return nil, nil
}

func (ps *PlayingScene) judge(p pack.SheetPack) {
	tick := ps.tick.Count()

	if inpututil.IsMouseButtonJustPressed(mouseButton) {
		// Records the now tick
		ps.inputHistory.Append(tick)

		for i, n := range p.Sheets {
			maxTiming := n.Timing + TimingRange
			minTiming := n.Timing - TimingRange

			if maxTiming >= tick && minTiming <= tick {
				ps.score += 1
				p.Sheets[i].MarkAsJudged()
				break
			}
		}
	}
}
