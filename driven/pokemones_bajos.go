package driven

import (
	"encoding/json"
	"io"
	"net/http"
)

type DrivenTraerPokemonesBajos struct {
	TraerPokemones DrivenTraerPokemonesExecuter
}

type DrivenTraerPokemonesBajosExecuter interface {
	DrivenTraerPokemonesBajos() ([]Pokemon, error)
}

func NewDrivenTraerPokemonesBajos(TraerPokemones DrivenTraerPokemonesExecuter) DrivenTraerPokemonesBajosExecuter {
	return &DrivenTraerPokemonesBajos{
		TraerPokemones: TraerPokemones,
	}
}
func (d *DrivenTraerPokemonesBajos) DrivenTraerPokemonesBajos() ([]Pokemon, error) {

	var respuesta Pokemon
	var pokemonesResultado []Pokemon

	pokemones, err := d.TraerPokemones.DrivenTraerPokemones()
	if err != nil {
		return []Pokemon{}, err
	}

	for _, pokemon := range pokemones.Results {

		resp, err := http.Get(pokemon.Url)
		if err != nil {
			return []Pokemon{}, err
		}

		defer resp.Body.Close()
		body, err := io.ReadAll(resp.Body)
		if err != nil {
			return []Pokemon{}, err
		}

		err = json.Unmarshal(body, &respuesta)
		if err != nil {
			return []Pokemon{}, err
		}

		pokemonesResultado = append(pokemonesResultado, respuesta)

	}
	return pokemonesResultado, err

}
