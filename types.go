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
	return Filter(m.Planets, func(p Planet) bool { return p.Owner == Myself })
}

func (m Map) NotMyPlanets() []Planet {
	return Filter(m.Planets, func(p Planet) bool { return p.Owner != Myself })
}

func (m Map) MyFleets() []Fleet {
	return Filter(m.Fleets, func(f Fleet) bool { return f.Owner == Myself })
}

func (m Map) FleetsTo(ID PlanetID) []Fleet {
	return Filter(m.Fleets, func(f Fleet) bool { return f.Dest == ID })
}

type Order struct {
	Source PlanetID
	Dest   PlanetID
	Ships
}
