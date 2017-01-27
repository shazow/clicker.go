package main

import (
	"time"

	"github.com/gopherjs/vecty"
	"github.com/shazow/clicker.go/components"
)

func main() {
	vecty.SetTitle("clicker.go")
	g := components.NewGame()
	g.AddCounter(&components.CounterView{Label: "requests"})

	vecty.RenderBody(g)

	ticker := time.Tick(500 * time.Millisecond)
	for _ = range ticker {
		g.Tick()
	}
}
