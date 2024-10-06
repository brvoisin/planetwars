package planetwars

type Point struct {
	X float64
	Y float64
}

type Owner int

const (
	Neutral  = Owner(0)
	Myself   = Owner(1)
	Opponent = Owner(2)
)

type Ships int

type Growth int

type PlanetID int

type Planet struct {
	ID       PlanetID
	Position Point
	Owner
	Ships
	Growth
}

type Trun int

type Fleet struct {
	Owner
	Ships
	Source        PlanetID
	Dest          PlanetID
	TotalTurn     Trun
	RemainingTurn Trun
}

type Map struct {
	Planets []Planet
	Fleets  []Fleet
}

func (m Map) PlanetByID(ID PlanetID) Planet {
	return m.Planets[ID]
}

func (m Map) MyPlanets() []Planet {
	result := make([]Planet, 0)
	for _, p := range m.Planets {
		if p.Owner == Myself {
			result = append(result, p)
		}
	}
	return result
}

type Order struct {
	Source PlanetID
	Dest   PlanetID
	Ships
}
