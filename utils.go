package planetwars

import "math"

func Distance(p1, p2 Planet) float64 {
	dx := p1.Position.X - p2.Position.X
	dy := p1.Position.Y - p2.Position.Y
	return math.Ceil(math.Sqrt(dx*dx + dy*dy))
}
