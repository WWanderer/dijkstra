package dijkstra

import (
	"math"
)

type Node struct {
	X int
	Y int
}

type Edge struct {
	Start, End *Node
	Cost       float64
}

type Graph struct {
	Edges []*Edge
}

// nodes functions

func Distance(a, b *Node) float64 {
	xDist := math.Pow(float64(b.X-a.X), 2)
	yDist := math.Pow(float64(b.Y-a.Y), 2)
	return math.Sqrt(xDist + yDist)
}

func (n *Node) Equal(a *Node) bool {
	if n.X == a.X && n.Y == a.Y {
		return true
	}
	return false
}

func NewNode(x, y int) *Node {
	n := new(Node)
	n.X = x
	n.Y = y
	return n
}

// edges functions

func NewEdge(S, E *Node) *Edge {
	e := new(Edge)
	e.Start = S
	e.End = E
	e.Cost = Distance(S, E)
	return e
}

func (e *Edge) Equal(a *Edge) bool {
	if e.Start.Equal(a.Start) && e.End.Equal(a.End) {
		return true
	}
	return false
}

// graph functions

func (g *Graph) Add(e *Edge) {
	g.Edges = append(g.Edges, e)
}

func (g *Graph) Remove(e *Edge) {
	index := -1
	for i, edge := range g.Edges {
		if edge.Equal(e) {
			index = i
			break
		}
	}
	if index >= 0 {
		g.Edges = append(g.Edges[:index], g.Edges[index+1:]...)
	}

}

func Dijkstra(g *Graph, source, dest *Node) ([]*Node, float64) {
	if source.Equal(dest) {
		return []*Node{source}, 0.0
	}
	dist := make(map[*Node]int)
	dist[source] = 0
	for i, e := range g.Edges {
		if e.Start.Equal(source) {
			dist[e.End] = e.Cost
		} else {
			dist[e.End] = -1
		}
	}

	return nil, 0.0
}

func main() {
	A := NewNode(1, 2)
	B := NewNode(1, 3)
	C := NewNode(-2, 9)
	D := NewNode(1, 1)
	E := NewNode(7, 6)
	F := NewNode(5, 5)

	g := new(Graph)

}
