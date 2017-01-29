package main

import (
	"time"

	"github.com/gopherjs/vecty"
	"github.com/shazow/clicker.go/components"
)

func main() {
	vecty.SetTitle("clicker.go")
	g := components.NewGame()

	reqCounter := components.NewCounterView("requests")
	reqCounter.Button("Add Goroutine")
	g.AddCounter(reqCounter)

	vecty.RenderBody(g)

	ticker := time.Tick(50 * time.Millisecond)
	for _ = range ticker {
		g.Tick()
	}
}
