package handler

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
)

func (h *handlers) Pokemones(c *fiber.Ctx) error {
	pokemones, err := h.usecases.TraerPokemones()
	if err != nil {
		return c.SendString(fmt.Sprintf("error %v", err.Error()))
	}

	return c.JSON(pokemones)

}
