package router

import (
	"github.com/gofiber/fiber/v2"
	"poo/handler"
)

type fiberRouter struct {
	fiber         *fiber.App
	handlersTraer handler.HandlersTraerPokemonesExecuter
	handlerBajos  handler.PokemonesBajosHandlerExecutor
	handlerAltos  handler.PokemonesAltosExecuter
}

func NewRouter(fiberA *fiber.App, handlersTraer handler.HandlersTraerPokemonesExecuter, handlerBajos handler.PokemonesBajosHandlerExecutor, handlerAltos handler.PokemonesAltosExecuter) fiberRouter {
	return fiberRouter{
		fiber:         fiberA,
		handlersTraer: handlersTraer,
		handlerBajos:  handlerBajos,
		handlerAltos:  handlerAltos,
	}
}

func (f *fiberRouter) Start() {

	f.fiber.Post("/Pokemones", f.handlersTraer.Pokemones)
	f.fiber.Post("/Bajos", f.handlerBajos.PokemonesBajos)
	f.fiber.Post("/Altos", f.handlerAltos.PokemonesAltos)

	f.fiber.Listen(":3000")

}
