package main

import (
	"math"
	"sort"

	"github.com/brvoisin/planetwars"
)

const untilLastFleet = -1

type brunoBot struct{}

func NewBrunoBot() planetwars.Player {
	return &brunoBot{}
}

// DoTurn implements Player.
func (b *brunoBot) DoTurn(pwMap planetwars.Map) []planetwars.Order {
	myPlanets := pwMap.MyPlanets()
	if len(myPlanets) == 1 {
		pwMap = worstFictiveScenario(pwMap, myPlanets[0])
	}
	candidates := computeCandidates(pwMap)
	orders := make([]planetwars.Order, 0)
	for _, candidate := range candidates {
		myPlanet := pwMap.PlanetByID(candidate.Source)
		destPlanet := pwMap.PlanetByID(candidate.Dest)
		futureDestPlanet := computePlanetState(pwMap, destPlanet, untilLastFleet)
		if futureDestPlanet.Owner == myPlanet.Owner {
			continue
		}
		fleetShips := computeNeededFleetShips(pwMap, myPlanet, destPlanet)
		if fleetShips <= 0 || fleetShips >= myPlanet.Ships {
			continue
		}
		myPlanet.Ships -= fleetShips
		myFuturePlanet := computePlanetState(pwMap, myPlanet, untilLastFleet)
		if myFuturePlanet.Owner != myPlanet.Owner {
			continue
		}
		orders = append(
			orders,
			planetwars.Order{Source: candidate.Source, Dest: candidate.Dest, Ships: fleetShips},
		)
		pwMap.Planets[myPlanet.ID] = myPlanet
		totalTurn := planetwars.Trun(planetwars.Distance(myPlanet, destPlanet))
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

func worstFictiveScenario(pwMap planetwars.Map, myPlanet planetwars.Planet) planetwars.Map {
	opponentPlanets := planetwars.Filter(
		pwMap.Planets,
		func(p planetwars.Planet) bool { return p.Owner == planetwars.Opponent },
	)
	for _, opponentPlanet := range opponentPlanets {
		totalTurn := planetwars.Trun(planetwars.Distance(opponentPlanet, myPlanet))
		fictiveFleet := planetwars.Fleet{
			Owner:         opponentPlanet.Owner,
			Ships:         opponentPlanet.Ships,
			Source:        opponentPlanet.ID,
			Dest:          myPlanet.ID,
			TotalTurn:     totalTurn,
			RemainingTurn: totalTurn,
		}
		pwMap.Fleets = append(pwMap.Fleets, fictiveFleet)
	}
	return pwMap
}

func computeNeededFleetShips(
	pwMap planetwars.Map,
	source planetwars.Planet,
	dest planetwars.Planet,
) planetwars.Ships {
	totalTurn := planetwars.Trun(planetwars.Distance(source, dest))
	futurePlanet := computePlanetState(pwMap, dest, totalTurn)
	if futurePlanet.Owner != source.Owner {
		return futurePlanet.Ships + 1
	} else {
		return -futurePlanet.Ships
	}
}

func computePlanetState(pwMap planetwars.Map, planet planetwars.Planet, maxTurn planetwars.Trun) planetwars.Planet {
	futurePlanet := planet
	fleetsTo := pwMap.FleetsTo(planet.ID)
	sort.SliceStable(fleetsTo, func(i, j int) bool {
		return fleetsTo[i].RemainingTurn < fleetsTo[j].RemainingTurn
	})
	turn := planetwars.Trun(0)
	for _, f := range fleetsTo {
		turnJump := f.RemainingTurn - turn
		turn += turnJump
		if maxTurn != untilLastFleet && turn > maxTurn {
			break
		}
		if futurePlanet.Owner != planetwars.Neutral {
			futurePlanet.Ships += planetwars.Ships(futurePlanet.Growth) * planetwars.Ships(turnJump)
		}
		futurePlanet.Ships += shipSign(futurePlanet.Owner, f.Owner) * f.Ships
		if futurePlanet.Ships == 0 {
			futurePlanet.Owner = planetwars.Neutral
		}
		if futurePlanet.Ships < 0 {
			futurePlanet.Owner = f.Owner
			futurePlanet.Ships = planetwars.Ships(math.Abs(float64(futurePlanet.Ships)))
		}
	}
	if maxTurn != untilLastFleet && turn < maxTurn && futurePlanet.Owner != planetwars.Neutral {
		futurePlanet.Ships += planetwars.Ships(futurePlanet.Growth) * planetwars.Ships(maxTurn-turn)
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
