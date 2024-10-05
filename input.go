package planetwarsbot

import (
	"bufio"
	"fmt"
	"io"
	"regexp"
	"strconv"
	"strings"
)

var planetLineRegexp = regexp.MustCompile("P" + strings.Repeat(`\s+([^\s]+)`, 5))
var fleetLineRegexp = regexp.MustCompile("F" + strings.Repeat(`\s+([^\s]+)`, 6))

func ParseInputMap(reader io.Reader) Map {
	rd := bufio.NewReader(reader)
	planets := make([]Planet, 0)
	fleets := make([]Fleet, 0)
	for {
		line, _, err := rd.ReadLine()
		if err != nil {
			panic(fmt.Errorf("input error before 'go' line: %w", err))
		}
		if string(line) == "go" {
			break
		}
		if len(line) == 0 {
			continue
		}
		if line[0] == 'P' {
			planets = append(planets, parsePlanet(string(line), len(planets)))
		}
		if line[0] == 'F' {
			fleets = append(fleets, parseFleet(string(line)))
		}
	}
	return Map{Planets: planets, Fleets: fleets}
}

func parsePlanet(line string, ID int) Planet {
	submatches := planetLineRegexp.FindStringSubmatch(line)
	return Planet{
		ID:       PlanetID(ID),
		Position: Point{X: parseFloat(submatches[1]), Y: parseFloat(submatches[2])},
		Owner:    Owner(parseInt(submatches[3])),
		Ships:    Ships(parseInt(submatches[4])),
		Growth:   Growth(parseInt(submatches[5])),
	}
}

func parseFloat(text string) float64 {
	result, err := strconv.ParseFloat(text, 32)
	if err != nil {
		panic(err)
	}
	return result
}

func parseInt(text string) int {
	result, err := strconv.Atoi(text)
	if err != nil {
		panic(err)
	}
	return result

}

func parseFleet(line string) Fleet {
	submatches := fleetLineRegexp.FindStringSubmatch(line)
	return Fleet{
		Owner:         Owner(parseInt(submatches[1])),
		Ships:         Ships(parseInt(submatches[2])),
		Source:        PlanetID(Growth(parseInt(submatches[3]))),
		Dest:          PlanetID(Growth(parseInt(submatches[4]))),
		TotalTurn:     Trun(parseInt(submatches[5])),
		RemainingTurn: Trun(parseInt(submatches[6])),
	}
}
