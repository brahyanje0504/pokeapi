package main

import (
	"github.com/gofiber/fiber/v2"
	"poo/driven"
	"poo/handler"
	"poo/router"
	"poo/usecases"
)

func Executer() {
	app := fiber.New()
	drivenTraerPokemon := driven.NewDrivenTraerPokemones()
	drivenPokemonesBajos := driven.NewDrivenTraerPokemonesBajos(drivenTraerPokemon)
	drivenPokemonsAltos := driven.NewDrivenTraerPokemonesBajos(drivenTraerPokemon)
	usecasesPokemonesAltos := usecases.NewUsecasesPokemonesAltos(drivenPokemonsAltos)
	usecasesPokemonesBajos := usecases.NewPokemonesBajos(drivenPokemonesBajos)
	usecasesTraerPokemon := usecases.NewUseCasesTraerPokemones(drivenTraerPokemon)
	handlerTraerPokemones := handler.NewHandlersTraerPokemones(usecasesTraerPokemon)
	handlerTraerPokemonesBajos := handler.NewPokemonesBajosHandler(usecasesPokemonesBajos)
	handlerTraerPokemonesAltos := handler.NewHandlerPokemonesAltos(usecasesPokemonesAltos)
	s := router.NewRouter(app, handlerTraerPokemones, handlerTraerPokemonesBajos, handlerTraerPokemonesAltos)
	s.Start()
}
