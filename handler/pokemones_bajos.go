package handler

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
)

func (h *handlers) PokemonesBajos(c *fiber.Ctx) error {
	pokemones, err := h.usecases.TraerPokemonesBajos()
	if err != nil {
		return c.SendString(fmt.Sprintf("error %v", err.Error()))
	}

	return c.JSON(pokemones)
}
