package main

import (
	"fmt"
	"math/rand"
	"time"

	"SF13.11/bst"
	"SF13.11/graf"
)

func init() {
	rand.Seed(time.Now().UnixNano()) // необходимо для того, чтобы рандом был похож на рандомный
}

func main() {
	g := graf.NewDirectedGraph()
	for i := 1; i < 11; i++ {
		g.AddVertex(i)
	}
	g.AddEdge(1, 2, 8)
	g.AddEdge(1, 5, 15)
	g.AddEdge(1, 6, 12)
	g.AddEdge(2, 3, 11)
	g.AddEdge(2, 4, 63)
	g.AddEdge(3, 5, 24)
	g.AddEdge(3, 4, 38)
	g.AddEdge(4, 9, 9)
	g.AddEdge(4, 5, 12)
	g.AddEdge(5, 8, 14)
	g.AddEdge(6, 5, 15)
	g.AddEdge(6, 7, 15)
	g.AddEdge(7, 9, 5)
	g.AddEdge(8, 7, 8)
	g.AddEdge(8, 9, 21)
	g.AddEdge(8, 4, 16)
	g.AddEdge(9, 1, 56)
	start := 1
	dest := 9
	for i := 0; i < 5; i++ {
		graf.ClearCosts(g)
		fmt.Println(graf.GetRoute(g.Vertices[start], g.Vertices[dest], g))
		fmt.Println("длина пути: ", g.Vertices[dest].Cost)
	}
	//g.AddEdge(10, 5, 100)
	/*graf.ClearCosts(g)

	fmt.Println(graf.GetRoute(g.Vertices[start], g.Vertices[dest], g))
	fmt.Println("длина пути: ", g.Vertices[dest].Cost)
	graf.ClearCosts(g)
	start = 1
	dest = 10
	fmt.Println(graf.GetRoute(g.Vertices[start], g.Vertices[dest], g))
	graf.ClearCosts(g)
	//g.DeleteVertex(6)
	start = 1
	dest = 9
	fmt.Println(graf.GetRoute(g.Vertices[start], g.Vertices[dest], g))
	fmt.Println("длина пути: ", g.Vertices[dest].Cost)
	graf.ClearCosts(g)*/
	//graf.GetCostsFromKey(g.Vertices[1])
	/*do := 0
	fmt.Print("Выберите:\n1 - двоичное дерево;\n2 - неориентированный граф, поиск в ширину;\n")
	fmt.Scanln(&do)
	switch do {
	case 1:
		printBst()
	case 2:
		printBFS()
	}
	*/
}

func printBFS() {
	g := graf.NewUndirectedGraph()
	for i := 1; i < 10; i++ {
		g.AddVertex(i)
	}
	g.AddEdge(1, 2, 8)
	g.AddEdge(1, 5, 15)
	g.AddEdge(1, 6, 12)
	g.AddEdge(2, 3, 11)
	g.AddEdge(2, 4, 63)
	g.AddEdge(3, 5, 24)
	g.AddEdge(3, 4, 38)
	g.AddEdge(4, 9, 9)
	g.AddEdge(4, 5, 12)
	g.AddEdge(5, 8, 14)
	g.AddEdge(6, 5, 15)
	g.AddEdge(6, 7, 15)
	g.AddEdge(7, 9, 5)
	g.AddEdge(8, 7, 8)
	g.AddEdge(8, 9, 21)
	g.AddEdge(8, 4, 16)
	//g.AddEdge(9, 1, 56)
	/*g := graf.NewUndirectedGraph()
	for i := 0; i < 10; i++ {
		g.AddVertex(i)
	}
	g.AddEdge(1, 6, 1)
	g.AddEdge(2, 3, 1)
	g.AddEdge(4, 8, 1)
	g.AddEdge(3, 4, 1)
	g.AddEdge(0, 8, 1)
	g.AddEdge(8, 5, 1)
	g.AddEdge(5, 3, 1)
	g.AddEdge(3, 1, 1)
	*/
	//fmt.Println(g)
	route := []int{}
	rf := func(i int) {
		route = append(route, i)
	}
	startVertex := 1
	if g.Vertices[startVertex] == nil {
		fmt.Printf("Нет такого стартового ключа: %d", startVertex)
		return
	} else {
		fmt.Printf("Стартуем из вершины с ключом: %d\n", startVertex)
	}

	for s := 0; s < 10; s++ {
		route = nil
		fmt.Printf("Найдем ключ %d? ", s)
		if graf.BfsToKey(g, g.Vertices[startVertex], s, rf) {
			fmt.Printf("Ключ %d найден: ", s)
			fmt.Println(route)
		} else {
			fmt.Printf("Ключ %d не найден\n", s)
		}
	}
}

func printBst() {
	t := bst.Root(15)
	nodes := []int{4, 5, 100, 1, 65, 43, 89, 13, -1, 7}
	for _, node := range nodes {
		t.Insert(node)
	}
	//вывод дерева как есть в качесте эксперимента,
	//работает не на всех входных данных
	fmt.Println("дерево:")
	t.PrintTree()

	fmt.Print("сортированный вывод: ")
	t.PrintInorder()
	fmt.Print("\n")
	f := 5
	if n, ok := t.Find(f); ok {
		fmt.Println("найден элемент: ", n.Value())
	} else {
		fmt.Printf("элемент %d не найден", f)
	}
	f = 12
	if n, ok := t.Find(f); ok {
		fmt.Println("найден элемент: ", n.Value())
	} else {
		fmt.Printf("элемент %d не найден\n", f)
	}
	d := 5
	fmt.Printf("удаление элемента %d\n", d)
	t.Delete(d)
	d = 7
	fmt.Printf("удаление элемента %d\n", d)
	t.Delete(d)
	d = 94
	fmt.Printf("вставка элемента %d\n", d)
	t.Insert(d)
	fmt.Print("итог: ")
	t.PrintInorder()
	fmt.Print("\n")
	// пробуем случайный массив
	fmt.Println("случайный массив: ")
	ar := make([]int, 50)
	for i := range ar {
		ar[i] = rand.Intn(200) - 100
	}
	rt := bst.Root(15)
	for _, node := range ar {
		rt.Insert(node)
	}
	fmt.Println("итог: ")
	rt.PrintInorder()
	fmt.Print("\n")
	fmt.Printf("минимальное значение: %d\n", rt.FindMin())
	fmt.Printf("максимальное значение: %d\n", rt.FindMax())
}
