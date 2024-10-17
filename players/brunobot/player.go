package main

import (
	"math"
	"sort"

	"github.com/brvoisin/planetwars"
)

type brunoBot struct{}

func NewBrunoBot() planetwars.Player {
	return &brunoBot{}
}

// DoTurn implements Player.
func (b *brunoBot) DoTurn(pwMap planetwars.Map) []planetwars.Order {
	candidates := computeCandidates(pwMap)
	orders := make([]planetwars.Order, 0)
	for _, candidate := range candidates {
		myPlanet := pwMap.PlanetByID(candidate.Source)
		fleetShips := computeNeededFleetShips(pwMap, candidate.Source, candidate.Dest)
		if fleetShips <= 0 || fleetShips >= myPlanet.Ships {
			continue
		}
		myPlanet.Ships -= fleetShips
		futurePlanet := planetStateAfterFleets(pwMap, myPlanet)
		if futurePlanet.Owner != myPlanet.Owner {
			continue
		}
		if myPlanet.Ships < 10 {
			continue
		}
		orders = append(
			orders,
			planetwars.Order{Source: candidate.Source, Dest: candidate.Dest, Ships: fleetShips},
		)
		pwMap.Planets[myPlanet.ID] = myPlanet
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
	result := pDest.Ships
	if pDest.Owner == planetwars.Opponent {
		result += planetwars.Ships(planetwars.Distance(pSrc, pDest) * float64(pDest.Growth))
	}
	for _, f := range pwMap.FleetsTo(dest) {
		if f.Owner == pDest.Owner {
			result += f.Ships
		} else {
			result -= f.Ships
		}
	}
	return result + 1
}

func planetStateAfterFleets(pwMap planetwars.Map, planet planetwars.Planet) planetwars.Planet {
	futurePlanet := planetwars.Planet(planet)
	fleetsTo := pwMap.FleetsTo(planet.ID)
	sort.SliceStable(fleetsTo, func(i, j int) bool {
		return fleetsTo[i].RemainingTurn < fleetsTo[j].RemainingTurn
	})
	turn := planetwars.Trun(0)
	for _, f := range fleetsTo {
		futurePlanet.Ships += planetwars.Ships(futurePlanet.Growth) * planetwars.Ships(f.RemainingTurn-turn)
		shipSign := 1
		if futurePlanet.Owner != f.Owner {
			shipSign = -1
		}
		futurePlanet.Ships += planetwars.Ships(shipSign) * f.Ships
		if futurePlanet.Ships == 0 {
			futurePlanet.Owner = planetwars.Neutral
		}
		if futurePlanet.Ships < 0 {
			futurePlanet.Owner = f.Owner
			futurePlanet.Ships = planetwars.Ships(math.Abs(float64(futurePlanet.Ships)))
		}
		turn = f.RemainingTurn
	}
	return futurePlanet
}
