package state

import "github.com/Dhairya3124/PizzaTime-Backend-Golang/internal/database"

type State struct {
	DB *database.Queries
}

func New() State {
	return State{}
}
