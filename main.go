package main

import (
	"github.com/gofiber/fiber/v2"
	"poo/handler"
	"poo/router"
)

func main() {
	app := fiber.New()         //
	h := handler.NewHandlers() //
	s := router.NewRouter(app, h)
	s.Start()
}
