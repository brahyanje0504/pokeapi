package usecases

import (
	"poo/driven"
)

type UseCasesTraerPokemonesExecuter interface {
	TraerPokemones() (driven.Pokemones, error)
	TraerPokemonesBajos() ([]driven.Pokemon, error)
	TraerPokemonesAltos() ([]driven.Pokemon, error)
}

type useCasesTraerPokemones struct {
	driven driven.DrivenExecuter
}

func NewUseCasesTraerPokemones(driven driven.DrivenExecuter) UseCasesTraerPokemonesExecuter {
	return &useCasesTraerPokemones{
		driven: driven,
	}
}
