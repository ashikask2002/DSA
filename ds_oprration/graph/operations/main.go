package main

import (
	"fmt"
)

type graph struct {
	list []*vertex
}
type vertex struct {
	data     int
	adjacent []*vertex
}

func main() {
	g := graph{}
	g.addVertex(1)
	g.addVertex(2)
	g.addVertex(3)
	g.addVertex(4)
	g.addVertex(5)

	g.addEdge(1, 2)
	g.addEdge(1, 3)
	g.addEdge(3, 4)
	g.addEdge(5, 3)
	g.addEdge(2, 4)

	g.print()
	fmt.Println("\nAfterrrrrrr")
	g.deleteVetex(5)
	g.deleteEdge(1, 3)
	g.print()

}
func (g *graph) addVertex(data int) {
	if !contains(g.list, data) {
		g.list = append(g.list, &vertex{data: data})
	}
}
func contains(vertex []*vertex, data int) bool {
	for _, v := range vertex {
		if v.data == data {
			return true
		}
	}
	return false
}
func (g *graph) print() {
	for _, v := range g.list {
		fmt.Print("\n vertex", v.data, ":")
		for _, v := range v.adjacent {
			fmt.Print(" ", v.data)
		}

	}
}
func (g *graph) getVertex(data int) *vertex {
	for _, v := range g.list {
		if v.data == data {
			return v
		}
	}
	return nil
}
func (g *graph) addEdge(from, to int) {
	fromVertex := g.getVertex(from)
	toVertex := g.getVertex(to)
	if fromVertex == nil || toVertex == nil {
		fmt.Println("Error: Vertex not found.")
		return
	}

	fromVertex.adjacent = append(fromVertex.adjacent, toVertex)
	toVertex.adjacent = append(toVertex.adjacent, fromVertex)
}
func (v *vertex) removeAdjacent(data int) {
	for i, adj := range v.adjacent {
		if adj.data == data {
			v.adjacent = append(v.adjacent[:i], v.adjacent[i+1:]...)
			return
		}
	}
}
func (g *graph) deleteVetex(data int) {
	for i, v := range g.list {
		if v.data == data {
			g.list = append(g.list[:i], g.list[i+1:]...)
			for _, vertex := range g.list {
				vertex.removeAdjacent(data)
			}
		}
	}
}
func (g *graph) deleteEdge(from, to int) {
	vertexFrom := g.getVertex(from)
	vertexTo := g.getVertex(to)
	if vertexFrom == nil || vertexTo == nil {
		fmt.Println("Error")
		return
	}
	vertexFrom.removeAdjacent(to)
	vertexTo.removeAdjacent(from)
}
