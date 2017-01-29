package components

import (
	"github.com/gopherjs/vecty"
	"github.com/gopherjs/vecty/elem"
	"github.com/gopherjs/vecty/event"
)

type IncrementButton struct {
	vecty.Core

	Label   string
	OnClick func(event *vecty.Event)
}

func (inc *IncrementButton) Render() *vecty.HTML {
	parts := []vecty.MarkupOrComponentOrHTML{
		vecty.Text(inc.Label),
	}
	if inc.OnClick != nil {
		parts = append(parts, event.Click(inc.OnClick))
	}
	return elem.Button(parts...)
}
