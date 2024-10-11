package main

import (
	"sort"

	"github.com/brvoisin/planetwars"
)

func computeCandidates(planetMap planetwars.Map) []Candidate {
	candidates := make([]Candidate, 0, pow2Int(len(planetMap.Planets)))
	for _, myPlanet := range planetMap.MyPlanets() {
		for _, planet := range planetMap.Planets {
			candidates = append(
				candidates,
				Candidate{myPlanet.ID, planet.ID, computeScore(planetMap, myPlanet, planet)},
			)
		}
	}
	sort.SliceStable(candidates, func(i, j int) bool {
		return candidates[i].Score > candidates[j].Score
	})
	return candidates
}

func pow2Int(x int) int {
	return x * x
}

func computeScore(_ planetwars.Map, srcPlanet planetwars.Planet, destPlanet planetwars.Planet) Score {
	if srcPlanet.Owner != planetwars.Myself {
		return 0
	}
	if destPlanet.Owner != planetwars.Myself && destPlanet.Ships < srcPlanet.Ships {
		return Score(1 / planetwars.Distance(srcPlanet, destPlanet))
	}
	return 0
}
