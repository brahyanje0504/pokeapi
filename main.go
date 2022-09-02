package main

import (
	"github.com/gofiber/fiber/v2"
	"poo/driven"
	"poo/handler"
	"poo/router"
	"poo/usecases"
)

func main() {
	app := fiber.New()
	d := driven.NewDriven()
	u := usecases.NewUseCasesTraerPokemones(d)
	h := handler.NewHandlers(u) //
	s := router.NewRouter(app, h)
	s.Start()
}
