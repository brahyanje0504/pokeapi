package handler

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"poo/usecases"
)

type HandlersTraerPokemonesExecuter interface {
	Pokemones(c *fiber.Ctx) error
}

type handlersTraerPokemones struct {
	usecases usecases.UseCasesTraerPokemonesExecuter
}

func NewHandlersTraerPokemones(usecases usecases.UseCasesTraerPokemonesExecuter) HandlersTraerPokemonesExecuter {
	return &handlersTraerPokemones{
		usecases: usecases,
	}
}

func (h *handlersTraerPokemones) Pokemones(c *fiber.Ctx) error {
	pokemones, err := h.usecases.TraerPokemones()
	if err != nil {
		return c.SendString(fmt.Sprintf("error %v", err.Error()))
	}

	return c.JSON(pokemones)

}
