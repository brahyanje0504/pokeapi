package usecases

import "poo/driven"

type PokemonesAltos struct {
	pokemonesBajos driven.DrivenTraerPokemonesBajosExecuter
}

type PokemonesAltosExecuter interface {
	TraerPokemonesAltos() ([]driven.Pokemon, error)
}

func NewUsecasesPokemonesAltos(PokemonesBajos driven.DrivenTraerPokemonesBajosExecuter) PokemonesAltosExecuter {
	return &PokemonesAltos{
		pokemonesBajos: PokemonesBajos,
	}

}

func (p *PokemonesAltos) TraerPokemonesAltos() ([]driven.Pokemon, error) {
	var dp []driven.Pokemon

	pokemonesResultado, err := p.pokemonesBajos.DrivenTraerPokemonesBajos()
	if err != nil {
		return []driven.Pokemon{}, err
	}

	for i := 0; i < len(pokemonesResultado); i++ {
		if pokemonesResultado[i].Height >= 15 {
			dp = append(dp, pokemonesResultado[i])
		}

	}

	return dp, err
}
