package pack

import "github.com/hajimehoshi/ebiten/v2"

type ScorePack struct {
	// Score data
	Scores []Note

	// File path
	Source string
}

func (sp ScorePack) MaxTick() int {
	var maxTick int

	for _, n := range sp.Scores {
		maxTick = n.Timing
	}

	return maxTick
}

type Note struct {
	// What mouse button user use to input
	Button ebiten.MouseButton

	// A ticks to determine timing
	Timing int
}
