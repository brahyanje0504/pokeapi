package usecases

import (
	"fmt"
	"poo/driven"
)

type UseCasesTraerPokemonesExecuter interface {
	TraerPokemones() (driven.Pokemones, error)
}

type useCasesTraerPokemones struct {
	driven driven.DrivenExecuter
}

func NewUseCasesTraerPokemones(driven driven.DrivenExecuter) UseCasesTraerPokemonesExecuter {
	return &useCasesTraerPokemones{
		driven: driven,
	}
}

func (u *useCasesTraerPokemones) TraerPokemones() (driven.Pokemones, error) {
	pokemones, err := u.driven.DrivenTraerPoemones()
	if err != nil {
		return driven.Pokemones{}, err
	}

	for i := 0; i < len(pokemones.Results); i++ {
		pokemones.Results[i].Name = fmt.Sprintf("%v-%v", pokemones.Results[i].Name, 1)

	}
	return pokemones, err

}
