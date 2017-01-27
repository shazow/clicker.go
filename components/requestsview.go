package components

import (
	"fmt"
	"sync"

	"github.com/gopherjs/vecty"
	"github.com/gopherjs/vecty/elem"
)

type Counter interface {
	vecty.Component
	Increment(int)
}

type CounterView struct {
	vecty.Core
	mu           sync.RWMutex
	requestCount int

	Label string
}

func (c *CounterView) Increment(amount int) {
	c.mu.Lock()
	c.requestCount += amount
	c.mu.Unlock()
}

// Render implements the vecty.Component interface.
func (c *CounterView) Render() *vecty.HTML {
	c.mu.RLock()
	reqText := fmt.Sprintf("%d %s", c.requestCount, c.Label)
	c.mu.RUnlock()

	return elem.Div(
		vecty.Text(reqText),
	)
}
