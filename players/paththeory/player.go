package main

import (
	"sort"

	"github.com/brvoisin/planetwars"
)

type starterBot struct {
	source planetwars.Planet
	target planetwars.Planet
	path   Path
}

type Path []planetwars.Planet

func NewStarterBot() planetwars.Player {
	return &starterBot{target: planetwars.Planet{ID: -1}}
}

// DoTurn implements Player.
// It implements the same basic strategy as the starter bot from
// https://github.com/xtevenx/planet-wars-starterpackage/.
func (b *starterBot) DoTurn(pwMap planetwars.Map) []planetwars.Order {
	if b.target.ID == -1 || b.target.Owner == planetwars.Myself {
		b.source = b.FindSource(pwMap)
		b.target = b.FindTarget(pwMap)
		b.path = b.FindPath(pwMap, b.source, b.target)
	}
	var orders []planetwars.Order
	for i := range len(b.path) - 1 {
		planet := b.path[i]
		if planet.Owner == planetwars.Myself {
			orders = append(orders, planetwars.Order{
				Source: planet.ID,
				Dest:   b.path[i+1].ID,
				Ships:  planet.Ships / 2,
			})
		}
	}
	return orders
}

func (b *starterBot) FindSource(pwMap planetwars.Map) planetwars.Planet {
	planets := pwMap.MyPlanets()
	if len(planets) != 0 {
		return planets[0]
	}
	return planetwars.Planet{}
}

func (b *starterBot) FindTarget(pwMap planetwars.Map) planetwars.Planet {
	planets := pwMap.NotMyPlanets()
	for i := range planets {
		if planets[i].Owner == planetwars.Opponent {
			return planets[i]
		}
	}
	return planetwars.Planet{}
}

func (b *starterBot) FindPath(pwMap planetwars.Map, source, target planetwars.Planet) Path {
	path := Path{source}
	planets := pwMap.Planets
	distance := planetwars.Distance(source, target)
	for _, planet := range planets {
		if planet.ID == source.ID || planet.ID == target.ID {
			continue
		}
		relDist := planetwars.Distance(source, planet) + planetwars.Distance(planet, target)
		if relDist-distance < 10 {
			path = append(path, planet)
		}
	}
	sort.Slice(path, func(i, j int) bool {
		return planetwars.Distance(source, path[i]) < planetwars.Distance(source, path[j])
	})
	path = append(path, target)
	return path
}
