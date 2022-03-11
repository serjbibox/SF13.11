package graf

import "math"

func GetPath(start *Vertex, dest *Vertex) (route []int) {
	start.Cost = 0

	min := math.MaxInt
	minKey := 0
	for key, val := range start.Vertices {
		if val.EdgeWeight < min {
			min = val.EdgeWeight
			minKey = key
		}
	}
	return route
}
