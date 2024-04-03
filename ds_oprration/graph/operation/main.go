package main

import "fmt"

type graph struct {
	list []*vertex
}

type vertex struct {
	data     int
	adjacent []*vertex
}

func (g *graph) addvertex(data int) {
	g.list = append(g.list, &vertex{data: data})
}

func contains(vertex []*vertex, data int) bool {
	for _, v := range vertex {
		if v.data == data {
			return true
		}
	}
	return false
}

func (g *graph) Print() {
	for _, v := range g.list {	
		fmt.Print("\n vertex ", v.data, ":")
		for _, v := range v.adjacent {
			fmt.Print(" ", v.data)
		}
	}
}

func (g *graph) getvertex(data int) *vertex {
	for _, val := range g.list {
		if val.data == data {
			return val
		}
	}
	return nil
}

func (g *graph) addedge(from, to int) {
	fromvertex := g.getvertex(from)
	tovertex := g.getvertex(to)

	if fromvertex == nil || tovertex == nil {
		fmt.Println("vertex not found")
		return
	}
	fromvertex.adjacent = append(fromvertex.adjacent, tovertex)
	tovertex.adjacent = append(tovertex.adjacent, fromvertex)

}

func main() {
  g := &graph{}
  g.addvertex(4)
  g.addvertex(5)
  g.addvertex(6)
  g.addvertex(7)
  g.addvertex(8)
  g.addvertex(9)

  g.addedge(4,6)
  g.addedge(7,9)
  g.addedge(6,8)

  g.Print()
}
