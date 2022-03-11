package graf

type node struct {
	v    *Vertex
	next *node
}

type queue struct {
	head *node
	tail *node
}

func (q *queue) enqueue(v *Vertex) {
	n := &node{v: v}
	if q.tail == nil {
		q.head = n
		q.tail = n
		return
	}
	q.tail.next = n
	q.tail = n
}

func (q *queue) dequeue() *Vertex {
	n := q.head
	if n == nil {
		return nil
	}

	q.head = q.head.next
	if q.head == nil {
		q.tail = nil
	}

	return n.v
}

//Обход графа, начиная со стартовой вершины startVertex
func BFS(g *Graph, startVertex *Vertex, visitCb func(int)) {
	vertexQueue := &queue{}
	visitedVertices := map[int]bool{}
	currentVertex := startVertex
	for {
		visitCb(currentVertex.Key)
		visitedVertices[currentVertex.Key] = true
		for _, v := range currentVertex.Vertices {
			if !visitedVertices[v.Key] {
				vertexQueue.enqueue(v)
			}
		}
		currentVertex = vertexQueue.dequeue()
		if currentVertex == nil {
			break
		}
	}
}

//Поиск заданного ключа key, начиная со стартовой вершины startVertex
func BfsToKey(g *Graph, startVertex *Vertex, key int, route func(int)) bool {
	vertexQueue := &queue{}
	visitedVertices := map[int]bool{}
	currentVertex := startVertex
	for {
		route(currentVertex.Key)
		if currentVertex.Key == key {
			return true
		}
		visitedVertices[currentVertex.Key] = true
		for _, v := range currentVertex.Vertices {
			if !visitedVertices[v.Key] {
				vertexQueue.enqueue(v)
			}
		}
		currentVertex = vertexQueue.dequeue()
		if currentVertex == nil {
			break
		}
	}
	return false
}
