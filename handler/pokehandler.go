package handler

import "github.com/gofiber/fiber/v2"

type HandlersExecuter interface {
	Pokemones(c *fiber.Ctx) error
}

type handlers struct{}

func NewHandlers() HandlersExecuter {
	return &handlers{}
}

func (h *handlers) Pokemones(c *fiber.Ctx) error {
	return c.SendString("hola ejecutarion modificar nombre")
}
