package driven

type driven struct {
}

type DrivenExecuter interface {
	DrivenTraerPoemones() (Pokemones, error)
	DrivenTraerPokemonesBajos() ([]Pokemon, error)
}

func NewDriven() DrivenExecuter {
	return &driven{}
}
