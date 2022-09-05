package handler

import (
	"fmt"
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

func (h *handlers) Pokemones(c *fiber.Ctx) error {
	pokemones, err := h.usecases.TraerPokemones()
	if err != nil {
		return c.SendString(fmt.Sprintf("error %v", err.Error()))
	}

	return c.JSON(pokemones)

}

func (h *handlers) PokemonesBajos(c *fiber.Ctx) error {
	pokemones, err := h.usecases.TraerPokemonesBajos()
	if err != nil {
		return c.SendString(fmt.Sprintf("error %v", err.Error()))
	}

	return c.JSON(pokemones)
}

func (h *handlers) PokemonesAltos(c *fiber.Ctx) error {
	pokemones, err := h.usecases.TraerPokemonesAltos()
	if err != nil {
		return c.SendString(fmt.Sprintf("error %v", err.Error()))
	}

	return c.JSON(pokemones)
}
