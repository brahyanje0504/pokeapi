package handler

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"poo/usecases"
)

type PokemonesAltosHandler struct {
	usecases usecases.PokemonesAltosExecuter
}

type PokemonesAltosExecuter interface {
	PokemonesAltos(c *fiber.Ctx) error
}

func NewHandlerPokemonesAltos(usecases usecases.PokemonesAltosExecuter) PokemonesAltosExecuter {
	return &PokemonesAltosHandler{
		usecases: usecases,
	}
}

func (p *PokemonesAltosHandler) PokemonesAltos(c *fiber.Ctx) error {
	pokemones, err := p.usecases.TraerPokemonesAltos()
	if err != nil {
		return c.SendString(fmt.Sprintf("error %v", err.Error()))
	}

	return c.JSON(pokemones)
}
