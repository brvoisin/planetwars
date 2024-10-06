package planetwars

import "math"

func Distance(p1, p2 Planet) float64 {
	dx := p1.Position.X - p2.Position.X
	dy := p1.Position.Y - p2.Position.Y
	return math.Ceil(math.Sqrt(dx*dx + dy*dy))
}

func Filter[T any](elements []T, test func(element T) bool) []T {
	result := make([]T, 0, len(elements))
	for _, el := range elements {
		if test(el) {
			result = append(result, el)
		}
	}
	return result
}
