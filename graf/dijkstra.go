package graf

import (
	"fmt"
	"math"
	"sort"
)

func GetRoute(start *Vertex, dest *Vertex, g *Graph) (route []int) {
	if start == nil {
		fmt.Println("стартовой вершины не существует")
		return
	}
	if dest == nil {
		fmt.Println("целевой вершины не существует")
		return
	}
	start.Cost = 0
	vertexQueue := &queue{}
	visitedVertices := map[int]bool{}
	current := start
	for {
		visitedVertices[current.Key] = true
		costArray := []int{}
		costMap := map[int]int{}
		for key, v := range current.Edges {
			costMap[v] = key
			costArray = append(costArray, v)
		}
		sort.Ints(costArray)
		for i := 0; i < len(costArray); i++ {
			if !visitedVertices[costMap[costArray[i]]] && !current.Vertices[costMap[costArray[i]]].Visited {
				vertexQueue.enqueue(current.Vertices[costMap[costArray[i]]])
			}
		}
		setCost(current)
		current = vertexQueue.dequeue()
		if current == nil {
			break
		}

	}
	if dest.Cost == math.MaxInt {
		fmt.Println("вершина недостижима!")
		return nil
	}
	//восстанавливает маршрут
	v := g.Vertices[dest.Key]
	route = append(route, dest.Key)
	for {
		if len(v.CostRoute) == 0 {
			break
		}
		route = append(route, v.CostRoute[len(v.CostRoute)-1])
		if g.Vertices[v.CostRoute[len(v.CostRoute)-1]] == nil {
			break
		}
		v = g.Vertices[v.CostRoute[len(v.CostRoute)-1]]
	}
	//разворачивает маршрут
	for i, j := 0, len(route)-1; i < j; i, j = i+1, j-1 {
		route[i], route[j] = route[j], route[i]
	}
	fmt.Println("вершина достигнута по пути:")
	return route
}

func GetCostsFromKey(start *Vertex) {
	start.Cost = 0
	vertexQueue := &queue{}
	visitedVertices := map[int]bool{}
	current := start
	for {
		visitedVertices[current.Key] = true
		for _, v := range current.Vertices {
			if !visitedVertices[v.Key] && !v.Visited {
				vertexQueue.enqueue(v)
			}
		}
		setCost(current)
		current = vertexQueue.dequeue()
		if current == nil {
			break
		}
	}
}

func setCost(v *Vertex) (endFlag bool) {
	for key, val := range v.Vertices {
		endFlag = true
		if val.Visited {
			continue
		}
		endFlag = false
		//prevCost := val.Cost
		tempCost := v.Edges[key] + v.Cost

		if tempCost < val.Cost {
			//fmt.Println("key", val.Key, "cost", tempCost, "prevCost", prevCost)
			val.Cost = tempCost
			val.CostRoute = append(val.CostRoute, v.Key)
		}

	}
	v.Visited = true
	return endFlag
}

func ClearCosts(g *Graph) {
	for _, v := range g.Vertices {
		if v != nil {
			v.Cost = math.MaxInt
			v.Visited = false
			v.CostRoute = []int{}
		}
	}
}
