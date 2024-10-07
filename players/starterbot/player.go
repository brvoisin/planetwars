package main

import (
	"github.com/brvoisin/planetwars"
)

type starterBot struct{}

func NewStarterBot() planetwars.Player {
	return &starterBot{}
}

// DoTurn implements Player.
// It implements the same basic strategy as the starter bot from
// https://github.com/xtevenx/planet-wars-starterpackage/.
func (b *starterBot) DoTurn(pwMap planetwars.Map) []planetwars.Order {
	// (1) If we currently have a fleet in flight, just do nothing.
	if len(pwMap.MyFleets()) == 1 {
		return nil
	}

	// (2) Find my strongest planet.
	source := planetwars.PlanetID(-1)
	sourceScore := -1
	sourceShips := planetwars.Ships(0)
	for _, p := range pwMap.MyPlanets() {
		score := int(p.Ships)
		if score > sourceScore {
			sourceScore = score
			source = p.ID
			sourceShips = p.Ships
		}
	}

	// (3) Find the weakest enemy or neutral planet.
	dest := planetwars.PlanetID(-1)
	destScore := -1.
	for _, p := range pwMap.NotMyPlanets() {
		score := 1. / (1 + float64(p.Ships))
		if score > destScore {
			destScore = score
			dest = p.ID
		}
	}

	// (4) Send half the ships from my strongest planet to the weakest
	// planet that I do not own.
	if source >= 0 && dest >= 0 {
		return []planetwars.Order{{Source: source, Dest: dest, Ships: sourceShips / 2}}
	}
	return nil
}
