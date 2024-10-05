package planetwarsbot

import "math"

func Distance(p1, p2 Planet) float64 {
	dx := p1.Position.X - p2.Position.X
	dy := p1.Position.Y - p2.Position.Y
	return math.Ceil(math.Sqrt(Pow2(dx) + Pow2(dy)))
}

func Pow2(x float64) float64 {
	return x * x
}

func Pow2Int(x int) int {
	return x * x
}
