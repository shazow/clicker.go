package components

import (
	"fmt"
	"sync"

	"github.com/gopherjs/vecty"
	"github.com/gopherjs/vecty/elem"
)

type Buff interface {
	Tick() float32
}

type buff struct {
}

type Counter interface {
	vecty.Component
	Increment()
}

func NewCounterView(label string) *CounterView {
	return &CounterView{
		rate:  0.1,
		Label: label,
	}
}

type CounterView struct {
	vecty.Core
	mu     sync.RWMutex
	count  float64
	rate   float64
	button vecty.Component

	Label string
}

func (c *CounterView) Button(label string) {
	c.button = &IncrementButton{
		Label: label,
		OnClick: func(event *vecty.Event) {
			c.rate += 0.2
		},
	}
}

func (c *CounterView) Increment() {
	c.mu.Lock()
	c.count += c.rate
	c.mu.Unlock()
}

// Render implements the vecty.Component interface.
func (c *CounterView) Render() *vecty.HTML {
	c.mu.RLock()
	reqText := fmt.Sprintf("%d %s", int(c.count), c.Label)
	c.mu.RUnlock()

	if c.button != nil {
		return elem.Div(
			vecty.Text(reqText),
			c.button.Render(),
		)
	}
	return elem.Div(
		vecty.Text(reqText),
	)
}
