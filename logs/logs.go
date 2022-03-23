package logs

import (
	"fmt"
	"github.com/nelsongp/decorator/decorator"
)

type LogRegistry struct {
	Handler decorator.Decorator
}

func NewLogRegistry(handler decorator.Decorator) *LogRegistry {
	return &LogRegistry{handler}
}

func (lr *LogRegistry) Process() error {
	defer fmt.Println("Peticion")
	return lr.Handler.Process()
}
