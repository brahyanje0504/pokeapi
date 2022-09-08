package handler

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
)

func (h *handlers) PokemonesAltos(c *fiber.Ctx) error {
	pokemones, err := h.usecases.TraerPokemonesAltos()
	if err != nil {
		return c.SendString(fmt.Sprintf("error %v", err.Error()))
	}

	return c.JSON(pokemones)
}
