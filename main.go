package main

import (
	"fmt"
	"github.com/nelsongp/decorator/handlers"
	"github.com/nelsongp/decorator/logs"
	"github.com/nelsongp/decorator/route"
	"os"
)

func main() {
	route := route.NewRoute()
	start(&route)

	var path string
	_, err := fmt.Scanf("%s", &path)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	route.Exec(path)
}

func start(route *route.Route) {
	route.Add(logs.NewLogRegistry(handlers.NewHandlerSayHello()), "/saludar")
	route.Add(logs.NewLogRegistry(handlers.NewHandlerSayGoodbye()), "/despedirse")
}
