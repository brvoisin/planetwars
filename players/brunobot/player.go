package main

import (
	"github.com/brvoisin/planetwars"
)

type brunoBot struct{}

func NewBrunoBot() planetwars.Player {
	return &brunoBot{}
}

// DoTurn implements Player.
func (b *brunoBot) DoTurn(pwMap planetwars.Map) []planetwars.Order {
	candidates := computeCandidates(pwMap)
	remainingShips := make(map[planetwars.PlanetID]planetwars.Ships, len(pwMap.Planets))
	orders := make([]planetwars.Order, 0)
	for _, candidate := range candidates {
		_, ok := remainingShips[candidate.Source]
		if !ok {
			remainingShips[candidate.Source] = pwMap.PlanetByID(candidate.Source).Ships
		}
		fleetShips := computeNeededFleetShips(pwMap, candidate.Source, candidate.Dest)
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

func computeNeededFleetShips(
	pwMap planetwars.Map,
	source planetwars.PlanetID,
	dest planetwars.PlanetID,
) planetwars.Ships {
	pSrc := pwMap.PlanetByID(source)
	pDest := pwMap.PlanetByID(dest)
	var growthAtArrival int
	if pDest.Owner == planetwars.Opponent {
		growthAtArrival = int(planetwars.Distance(pSrc, pDest) * float64(pDest.Growth))
	}
	return pDest.Ships + planetwars.Ships(growthAtArrival) + 1
}
