package route

import (
	"fmt"
	"github.com/nelsongp/decorator/decorator"
)

type Route struct {
	decorators map[string]decorator.Decorator
}

func NewRoute() Route {
	return Route{make(map[string]decorator.Decorator)}
}

func (r *Route) Add(decorator decorator.Decorator, path string) {
	r.decorators[path] = decorator
}

func (r *Route) Exec(path string) {
	d, ok := r.decorators[path]
	if !ok {
		fmt.Println("no existe la ruta")
		return
	}

	err := d.Process()
	if err != nil {
		fmt.Println("ERROR", err)
	}
}
