package driven

import (
	"encoding/json"
	"io"
	"net/http"
)

type DrivenTraerPokemones struct {
}

type DrivenTraerPokemonesExecuter interface {
	DrivenTraerPokemones() (Pokemones, error)
}

func NewDrivenTraerPokemones() DrivenTraerPokemonesExecuter {
	return &DrivenTraerPokemones{}
}

func (d *DrivenTraerPokemones) DrivenTraerPokemones() (Pokemones, error) {
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
