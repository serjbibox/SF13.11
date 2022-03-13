package graf

import (
	"fmt"
	"math"
	"strconv"
)

type Vertex struct {
	Key       int
	Vertices  map[int]*Vertex
	Edges     map[int]int
	Cost      int
	Visited   bool
	CostRoute []int
}

func NewVertex(key int) *Vertex {
	return &Vertex{
		Key:       key,
		Vertices:  map[int]*Vertex{},
		Cost:      math.MaxInt,
		Edges:     map[int]int{},
		CostRoute: []int{},
	}
}

type Graph struct {
	Vertices map[int]*Vertex
	directed bool
}

func NewDirectedGraph() *Graph {
	return &Graph{
		Vertices: map[int]*Vertex{},
		directed: true,
	}
}

func NewUndirectedGraph() *Graph {
	return &Graph{
		Vertices: map[int]*Vertex{},
	}
}

func (g *Graph) AddVertex(key int) {
	v := NewVertex(key)
	g.Vertices[key] = v
}
func (g *Graph) DeleteVertex(key int) {
	delete(g.Vertices, key)
}

func (g *Graph) AddEdge(k1, k2 int, ew int) {
	v1 := g.Vertices[k1]
	v2 := g.Vertices[k2]
	if v1 == nil {
		panic(fmt.Sprintf("Отсутствует вершина: %v", k1))
	}
	if v2 == nil {
		panic(fmt.Sprintf("Отсутствует вершина: %v", k2))
	}

	if _, ok := v1.Vertices[v2.Key]; ok {
		return
	}

	v1.Vertices[v2.Key] = v2
	if !g.directed && v1.Key != v2.Key {
		v2.Vertices[v1.Key] = v1
	}
	g.Vertices[v1.Key] = v1
	g.Vertices[v2.Key] = v2
	g.Vertices[v1.Key].Edges[v2.Key] = ew
}

func (g *Graph) String() string {
	s := ""
	i := 0
	for _, v := range g.Vertices {
		if i != 0 {
			s += "\n"
		}
		s += v.String()
		i++
	}
	return s
}

func (v *Vertex) String() string {
	s := strconv.Itoa(v.Key) + ":"
	for _, neighbor := range v.Vertices {
		s += " " + strconv.Itoa(neighbor.Key)
	}
	return s
}
