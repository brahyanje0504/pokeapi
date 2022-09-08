package handler

import (
	"github.com/gofiber/fiber/v2"
	"poo/usecases"
)

type HandlersExecuter interface {
	Pokemones(c *fiber.Ctx) error
	PokemonesBajos(c *fiber.Ctx) error
	PokemonesAltos(c *fiber.Ctx) error
}

type handlers struct {
	usecases usecases.UseCasesTraerPokemonesExecuter
}

func NewHandlers(usecases usecases.UseCasesTraerPokemonesExecuter) HandlersExecuter {
	return &handlers{
		usecases: usecases,
	}
}
