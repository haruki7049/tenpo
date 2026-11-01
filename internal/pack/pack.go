package pack

import "github.com/hajimehoshi/ebiten/v2"

type SheetPack struct {
	// Sheet data
	Sheets []Note

	// File path
	Source string
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
