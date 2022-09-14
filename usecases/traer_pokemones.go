package usecases

import (
	"fmt"
	"poo/driven"
)

type useCasesTraerPokemones struct {
	driven driven.DrivenTraerPokemonesExecuter
}

type UseCasesTraerPokemonesExecuter interface {
	TraerPokemones() (driven.Pokemones, error)
}

func NewUseCasesTraerPokemones(driven driven.DrivenTraerPokemonesExecuter) UseCasesTraerPokemonesExecuter {
	return &useCasesTraerPokemones{
		driven: driven,
	}
}

func (u *useCasesTraerPokemones) TraerPokemones() (driven.Pokemones, error) {
	pokemones, err := u.driven.DrivenTraerPokemones()
	if err != nil {
		return driven.Pokemones{}, err
	}

	for i := 0; i < len(pokemones.Results); i++ {
		pokemones.Results[i].Name = fmt.Sprintf("%v-%v", pokemones.Results[i].Name, 1)

	}
	return pokemones, err

}
