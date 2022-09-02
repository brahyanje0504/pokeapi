package router

import (
	"github.com/gofiber/fiber/v2"
	"poo/handler"
)

type fiberRouter struct {
	fiber    *fiber.App
	handlers handler.HandlersExecuter
}

func NewRouter(fiberA *fiber.App, handlerA handler.HandlersExecuter) fiberRouter {
	return fiberRouter{
		fiber:    fiberA,
		handlers: handlerA,
	}
}

func (f *fiberRouter) Start() {

	f.fiber.Post("/Pokemones", f.handlers.Pokemones)

	f.fiber.Listen(":3000")

}
