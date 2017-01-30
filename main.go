package main

import (
	"fmt"
	"time"

	"github.com/gopherjs/vecty"
	"github.com/gopherjs/vecty/elem"
	"github.com/gopherjs/vecty/event"
)

func main() {
	vecty.SetTitle("clicker.go")

	count := 0.0
	rate := 0.1

	v := &View{
		OnClick: func(_ *vecty.Event) {
			rate += 0.1
		},
	}
	vecty.RenderBody(v)

	ticker := time.Tick(50 * time.Millisecond)
	for _ = range ticker {
		count += rate

		v.Count = int(count)
		vecty.Rerender(v)
	}
}

type View struct {
	vecty.Core

	Count   int
	OnClick func(event *vecty.Event)
}

func (v *View) Render() *vecty.HTML {
	return elem.Body(
		elem.Header1(
			vecty.Text("clicker.go"),
		),
		elem.Div(
			vecty.Text(fmt.Sprintf("%d %s", v.Count, "requests")),
		),
		elem.Div(
			elem.Button(
				vecty.Text("Add Goroutine"),
				event.Click(v.OnClick),
			),
		),
	)
}
