package driven

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
