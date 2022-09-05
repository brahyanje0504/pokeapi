package driven

import (
	"encoding/json"
	"io"
	"net/http"
)

type driven struct {
}

type DrivenExecuter interface {
	DrivenTraerPoemones() (Pokemones, error)
	DrivenTraerPokemonesBajos() ([]Pokemon, error)
}

func NewDriven() DrivenExecuter {
	return &driven{}
}

type Pokemones struct {
	Count    int    `json:"count"`
	Next     string `json:"next"`
	Previous string `json:"previous"`
	Results  []Pokemon
}
type Pokemon struct {
	Name   string `json:"name"`
	Url    string `json:"url"`
	Height int    `json:"height"`
}

func (d *driven) DrivenTraerPoemones() (Pokemones, error) {
	url := "https://pokeapi.co/api/v2/pokemon?limit=20&offset=0"
	resp, err := http.Get(url)
	if err != nil {
		return Pokemones{}, err
	}

	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return Pokemones{}, err
	}

	var respuesta Pokemones
	err = json.Unmarshal(body, &respuesta)
	if err != nil {
		return Pokemones{}, err
	}

	return respuesta, err

}

func (d *driven) DrivenTraerPokemonesBajos() ([]Pokemon, error) {

	var respuesta Pokemon
	var pokemonesResultado []Pokemon

	pokemones, err := d.DrivenTraerPoemones()
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
