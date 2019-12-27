package main

import (
	"fmt"
	"math"
)

// Node is the base unit the model, representing a point on a plane
type Node struct {
	X int
	Y int
}

// An Edge represents a vector going from Start to End
type Edge struct {
	Start, End *Node
	Cost       float64
}

// A Graph represents a set of Edges
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

func (n *Node) In(m map[*Node]float64) bool {
	_, ok := m[n]
	return ok
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

// InitMap returns a map with all Nodes in the graph and set the distance from
// source to the other Nodes as math.MaxFloat64, or 0 to itself
func InitMap(g *Graph, source, dest *Node) map[*Node]float64 {
	unvisited := make(map[*Node]float64)
	for _, e := range g.Edges {
		if !e.Start.In(unvisited) {
			if e.Start.Equal(source) {
				unvisited[e.Start] = 0
			} else {
				unvisited[e.Start] = math.MaxFloat64
			}
		}
		if !e.End.In(unvisited) {
			if e.End.Equal(source) {
				unvisited[e.End] = 0
			} else {
				unvisited[e.End] = math.MaxFloat64
			}
		}
	}
	return unvisited
}

func MinUnvisited(m map[*Node]float64) *Node {
	min := math.MaxFloat64
	var minNode *Node
	for k, v := range m {
		if min > v {
			min = v
			minNode = k
		}
	}
	return minNode
}

func Dijkstra(g *Graph, source, dest *Node) float64 {
	unvisited := InitMap(g, source, dest)

	if source.Equal(dest) {
		return 0.0
	}
	current := source
	for dest.In(unvisited) {
		for _, e := range g.Edges {
			if e.Start.Equal(current) {
				tmp := unvisited[current] + Distance(current, e.End)
				if tmp < unvisited[e.End] {
					unvisited[e.End] = tmp
				}
			}
		}

		delete(unvisited, current)
		current = MinUnvisited(unvisited)

		if current == dest {
			return unvisited[dest]
		}
		if unvisited[current] == math.MaxFloat64 {
			return -1
		}
	}

	return -1
}

func main() {
	A := NewNode(1, 2)
	B := NewNode(1, 3)
	C := NewNode(-2, 9)
	D := NewNode(1, 1)
	E := NewNode(7, 6)
	F := NewNode(5, 5)

	g := new(Graph)
	g.Add(NewEdge(A, B))
	g.Add(NewEdge(A, C))
	g.Add(NewEdge(A, F))
	g.Add(NewEdge(B, C))
	g.Add(NewEdge(B, D))
	g.Add(NewEdge(C, D))
	g.Add(NewEdge(C, F))
	g.Add(NewEdge(D, E))
	g.Add(NewEdge(F, E))

	fmt.Println("A", "B", Distance(A, B))
	fmt.Println("A", "C", Distance(A, C))
	fmt.Println("A", "F", Distance(A, F))
	fmt.Println("B", "C", Distance(B, C))
	fmt.Println("B", "D", Distance(B, D))
	fmt.Println("C", "D", Distance(C, D))
	fmt.Println("C", "F", Distance(C, F))
	fmt.Println("D", "E", Distance(D, E))
	fmt.Println("F", "E", Distance(F, E))
	fmt.Println(Dijkstra(g, A, E))
}
