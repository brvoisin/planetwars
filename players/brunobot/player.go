package main

import (
	"github.com/brvoisin/planetwars"
)

type brunoBot struct{}

func NewBrunoBot() planetwars.Player {
	return &brunoBot{}
}

// DoTurn implements Player.
func (b *brunoBot) DoTurn(planetMap planetwars.Map) []planetwars.Order {
	candidates := computeCandidates(planetMap)
	remainingShips := make(map[planetwars.PlanetID]planetwars.Ships, len(planetMap.Planets))
	orders := make([]planetwars.Order, 0)
	for _, candidate := range candidates {
		_, ok := remainingShips[candidate.Source]
		if !ok {
			remainingShips[candidate.Source] = planetMap.PlanetByID(candidate.Source).Ships
		}
		fleetShips := planetMap.PlanetByID(candidate.Dest).Ships + 1
		remainingShips[candidate.Source] -= fleetShips
		if remainingShips[candidate.Source] >= 10 {
			orders = append(
				orders,
				planetwars.Order{Source: candidate.Source, Dest: candidate.Dest, Ships: fleetShips},
			)
		}
	}
	return orders
}
