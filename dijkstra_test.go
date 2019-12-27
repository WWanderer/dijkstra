package main

import (
	"math"
	"testing"
)

func TestDistance(t *testing.T) {
	cases := []struct {
		in   []*Node
		want float64
	}{
		{[]*Node{{X: 1, Y: 2}, {X: 3, Y: 4}}, 2.8284},
		{[]*Node{{X: 1, Y: 3}, {X: -2, Y: 9}}, 6.7082},
		{[]*Node{{X: 1, Y: 2}, {X: 5, Y: 5}}, 5.0},
		{[]*Node{{X: 1, Y: 2}, {X: 7, Y: 6}}, 7.2111},
		{[]*Node{{X: 1, Y: 1}, {X: 7, Y: -7}}, 10.0},
	}
	for _, c := range cases {
		got := Distance(c.in[0], c.in[1])
		if math.Abs(got-c.want) > 0.001 {
			t.Error("Distance(", c.in[0], c.in[1], ") ==", got, "want", c.want)
		}
	}
}

func TestDijkstra(t *testing.T) {
	// g := new(Graph)
	// g.Add(New(*Node{1,2}, *Node{3, 4}))
}
