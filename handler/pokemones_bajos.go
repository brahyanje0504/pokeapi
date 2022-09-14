package handler

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"poo/usecases"
)

type PokemonesBajos struct {
	usecases usecases.PokemonesBajosExecuter
}

type PokemonesBajosHandlerExecutor interface {
	PokemonesBajos(c *fiber.Ctx) error
}

func NewPokemonesBajosHandler(usecases usecases.PokemonesBajosExecuter) PokemonesBajosHandlerExecutor {
	return &PokemonesBajos{
		usecases: usecases,
	}

}

func (h *PokemonesBajos) PokemonesBajos(c *fiber.Ctx) error {
	pokemones, err := h.usecases.TraerPokemonesBajos()
	if err != nil {
		return c.SendString(fmt.Sprintf("error %v", err.Error()))
	}

	return c.JSON(pokemones)
}
