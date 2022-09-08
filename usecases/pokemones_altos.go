package usecases

import "poo/driven"

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
