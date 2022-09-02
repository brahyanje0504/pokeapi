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
}

func NewDriven() DrivenExecuter {
	return &driven{}
}

type Pokemones struct {
	Count    int    `json:"count"`
	Next     string `json:"next"`
	Previous string `json:"previous"`
	Results  []pokemon
}
type pokemon struct {
	Name string `json:"name"`
	Url  string `json:"url"`
}

func (d *driven) DrivenTraerPoemones() (Pokemones, error) {
	url := "https://pokeapi.co/api/v2/pokemon?limit=10&offset=0"
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
