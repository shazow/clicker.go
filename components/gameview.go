package components

import (
	"github.com/gopherjs/vecty"
	"github.com/gopherjs/vecty/elem"
)

func NewGame() *GameView {
	return &GameView{
		Counters: []Counter{},
	}
}

type GameView struct {
	vecty.Core
	Counters []Counter
}

func (g *GameView) Render() *vecty.HTML {
	rendered := make(vecty.List, 0, len(g.Counters)) // TODO: Reuse
	for _, c := range g.Counters {
		rendered = append(rendered, c.Render())
	}
	return elem.Body(rendered...)
}

func (g *GameView) Tick() {
	for _, c := range g.Counters {
		c.Increment(1)
	}
	vecty.Rerender(g)
}

func (g *GameView) AddCounter(c Counter) {
	g.Counters = append(g.Counters, c)
}
