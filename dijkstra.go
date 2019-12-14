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

type Path []Node

func Distance(a, b *Node) float64 {
	xDist := math.Pow(float64(b.X-a.X), 2)
	yDist := math.Pow(float64(b.Y-a.Y), 2)
	return math.Sqrt(xDist + yDist)
}

func Dijkstra(g *Graph, source, dest *Node) (p Path, d float64) {
	return nil, 0.0
}
