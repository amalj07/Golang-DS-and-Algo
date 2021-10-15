package main

import "fmt"

type Graph struct {
	vertices []*Vertex
}

type Vertex struct {
	data     int
	adjacent []*Vertex
}

func (graph *Graph) addVertex(value int) {
	if graph.contains(value) {
		fmt.Println(value, "already exist in graph")
		return
	}
	graph.vertices = append(graph.vertices, &Vertex{data: value})
}

func (graph *Graph) addEdge(from, to int) {
	if !graph.contains(from) || !graph.contains(to) {
		fmt.Println("Failed to add edge: provided data not in graph")
		return
	}
	var fromVertex *Vertex
	var toVertex *Vertex
	for _, vertex := range graph.vertices {
		if from == vertex.data {
			fromVertex = vertex
		}
		if to == vertex.data {
			toVertex = vertex
		}
	}
	fromVertex.adjacent = append(fromVertex.adjacent, toVertex)
}

func (graph *Graph) contains(value int) bool {
	for _, vertext := range graph.vertices {
		if value == vertext.data {
			return true
		}
	}
	return false
}

// Print the graph in adjacency list
func (graph *Graph) print() {
	for _, vertex := range graph.vertices {
		fmt.Print(vertex.data, ": ")
		for _, adjacent := range vertex.adjacent {
			fmt.Print(adjacent.data, " ")
		}
		fmt.Println()
	}
}

func main() {
	g := Graph{}
	for i := 0; i < 5; i++ {
		g.addVertex(i)
	}
	g.addEdge(1, 3)
	g.addEdge(2, 1)
	g.addEdge(1, 4)
	g.addEdge(0, 2)
	g.addEdge(2, 5)
	g.print()
}
