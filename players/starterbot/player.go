package main

import (
	"github.com/brvoisin/planetwars"
)

type starterBot struct{}

func NewStarterBot() planetwars.Player {
	return &starterBot{}
}

// DoTurn implements Player.
func (b *starterBot) DoTurn(planetMap planetwars.Map) []planetwars.Order {
	// TODO
	return nil
}
