package graf

import (
	"fmt"
	"math"
	"sort"
)

func GetRoute(start *Vertex, dest *Vertex, g *Graph) (route []int) {

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
		for _, v := range costArray {
			if !visitedVertices[costMap[v]] && !current.Vertices[costMap[v]].Visited {
				vertexQueue.enqueue(current.Vertices[costMap[v]])
			}
		}
		//fmt.Println("visited", current.Key)

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

/*
func GetPath(start *Vertex, dest *Vertex) (route []int) {
	start.Cost = 0
	min := math.MaxInt
	minKey := 0
	for key, val := range start.Edges {
		if val < min {
			min = val
			minKey = key
		}
	}
	fmt.Println("1", start.Vertices[minKey].Key)
	cnt = 0
	rec(false, start, dest)
	//fmt.Print(minKey)
	return route
}

var cnt int

func rec(flag bool, start *Vertex, dest *Vertex) {
	cnt++
	min := math.MaxInt
	minKey := 0
	if flag {
		return
	}

	for key, val := range start.Vertices {
		flag = true
		if !val.Visited {
			flag = false
			tempCost := start.Edges[key] + start.Cost
			//prevCost := val.Cost
			if tempCost < val.Cost {
				start.Vertices[key].Cost = tempCost
			}
			//fmt.Println("key", start.Vertices[key].Key, "cost", start.Vertices[key].Cost, "prevCost", prevCost, "cnt", cnt)
			if start.Vertices[key].Cost < min {
				min = start.Vertices[key].Cost
				minKey = key
			}
		}
	}
	start.Visited = true
	fmt.Println("start.Key", start.Key, "start.Cost", start.Cost)
	if minKey == 0 {
		fmt.Println("flag", flag)
		return
	}

	//route = append(route)
	for _, val := range start.Vertices {
		if start.Vertices[minKey] == nil {
			fmt.Println("key", start.Key, "cost", start.Cost, "dest", "end")
		} else {
			fmt.Println("key", start.Key, "cost", start.Cost, "dest", start.Vertices[minKey].Key)
		}
		rec(false, val, dest)
	}
}
*/
