package main

import (
	"github.com/brvoisin/planetwarsbot"
)

type starterBot struct{}

func NewStarterBot() planetwarsbot.Player {
	return &starterBot{}
}

// DoTurn implements Player.
func (b *starterBot) DoTurn(planetMap planetwarsbot.Map) []planetwarsbot.Order {
	// TODO
	return nil
}
