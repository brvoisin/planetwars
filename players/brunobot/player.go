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
		totalTurn := planetwars.Trun(planetwars.Distance(myPlanet, pwMap.PlanetByID(candidate.Dest)))
		pwMap.Fleets = append(pwMap.Fleets, planetwars.Fleet{
			Owner:         myPlanet.Owner,
			Ships:         fleetShips,
			Source:        candidate.Source,
			Dest:          candidate.Dest,
			TotalTurn:     totalTurn,
			RemainingTurn: totalTurn,
		})
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
	result := -shipSign(pSrc.Owner, pDest.Owner) * pDest.Ships
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
		futurePlanet.Ships += shipSign(futurePlanet.Owner, f.Owner) * f.Ships
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

func shipSign(o1, o2 planetwars.Owner) planetwars.Ships {
	if o1 != o2 {
		return -1
	} else {
		return 1
	}
}
