package usecases

import (
	"fmt"
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

func (u *useCasesTraerPokemones) TraerPokemonesBajos() ([]driven.Pokemon, error) {

	var p []driven.Pokemon

	pokemonesResultado, err := u.driven.DrivenTraerPokemonesBajos()
	if err != nil {
		return []driven.Pokemon{}, err
	}

	for i := 0; i < len(pokemonesResultado); i++ {
		if pokemonesResultado[i].Height <= 15 {
			p = append(p, pokemonesResultado[i])
		}

	}

	return p, err
}

func (u *useCasesTraerPokemones) TraerPokemonesAltos() ([]driven.Pokemon, error) {
	var p []driven.Pokemon

	pokemonesResultado, err := u.driven.DrivenTraerPokemonesBajos()
	if err != nil {
		return []driven.Pokemon{}, err
	}

	for i := 0; i < len(pokemonesResultado); i++ {
		if pokemonesResultado[i].Height >= 15 {
			p = append(p, pokemonesResultado[i])
		}

	}

	return p, err
}
