package pack

import (
	"io/fs"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/haruki7049/tenpo/assets"
)

type SheetPack struct {
	// File path
	AudioSource string

	// Sheet data
	Sheets []Note
}

func (sp SheetPack) OpenAudioSource() (fs.File, error) {
	return assets.Assets.Open(sp.AudioSource)
}

func (sp SheetPack) MaxTick() int {
	var maxTick int

	for _, n := range sp.Sheets {
		maxTick = n.Timing
	}

	return maxTick
}

type Note struct {
	// What mouse button user use to input
	Button ebiten.MouseButton

	// A ticks to determine timing
	Timing int

	// Whether it is already judged or not
	IsAlreadyJudged bool
}

func (n *Note) MarkAsJudged() {
	n.IsAlreadyJudged = true
}
