package handlers

import "fmt"

type HandlersSayGoodbye struct{}

func NewHandlerSayGoodbye() *HandlersSayGoodbye {
	return &HandlersSayGoodbye{}
}

func (h *HandlersSayGoodbye) Process() error {
	fmt.Println("Adios")
	return nil
}
