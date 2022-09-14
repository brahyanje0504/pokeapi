package usecases

import "poo/driven"

type PokemonesBajos struct {
	pokemonesBajos driven.DrivenTraerPokemonesBajosExecuter
}

type PokemonesBajosExecuter interface {
	TraerPokemonesBajos() ([]driven.Pokemon, error)
}

func NewPokemonesBajos(traerPokemones driven.DrivenTraerPokemonesBajosExecuter) PokemonesBajosExecuter {
	return &PokemonesBajos{
		pokemonesBajos: traerPokemones,
	}
}

func (u *PokemonesBajos) TraerPokemonesBajos() ([]driven.Pokemon, error) {

	var p []driven.Pokemon

	pokemonesResultado, err := u.pokemonesBajos.DrivenTraerPokemonesBajos()
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
