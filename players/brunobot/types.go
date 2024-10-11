package main

import "github.com/brvoisin/planetwars"

type Score float64

type Candidate struct {
	Source planetwars.PlanetID
	Dest   planetwars.PlanetID
	Score
}
