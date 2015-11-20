/*
Package graphs defines a graph type and algorithms to traverse it
*/
package graphs

import "fmt"

// MAXV is the maximum number of vertices on a Graph
const MAXV = 10

// Edgenode represents an edge on an adjacency list
type Edgenode struct {
	Y      int
	Weight int
	next   *Edgenode
}

// Graph is an adjacency list representing a graph
type Graph struct {
	Edges       [MAXV + 1]*Edgenode
	Degree      [MAXV + 1]int
	NVertices   int
	NEdges      int
	Directed    bool
	initialized bool
}

func (g *Graph) auxiliaryAddEdge(x, y, weight int) {
	auxEdge := new(Edgenode)
	if g.Edges[x] == nil {
		g.Edges[x] = auxEdge
		g.NVertices++
	} else {
		ptr := g.Edges[x]
		for {
			if ptr.next == nil {
				ptr.next = auxEdge
				break
			}
			ptr = ptr.next
		}
	}
	auxEdge.Y = y
	auxEdge.Weight = weight
	g.Degree[x]++
}

// AddEdge adds a new edge to Graph structure, if weight = 0 it will be ignored
func (g *Graph) AddEdge(x, y, weight int) {
	if x >= MAXV || y >= MAXV {
		errMessage := fmt.Sprintf("x and y must be < %d", MAXV)
		panic(errMessage)
	}
	g.auxiliaryAddEdge(x, y, weight)
	if !g.Directed {
		g.auxiliaryAddEdge(y, x, weight)
	}
	g.NEdges++
}

// Print show infos from Graph and run through Edges printing the entire list
func (g *Graph) Print() {
	fmt.Println("Number of Vertices: ", g.NVertices)
	fmt.Println("Number of Edges: ", g.NEdges)
	fmt.Println("Directed: ", g.Directed)
	var ptr *Edgenode
	for i, v := range g.Edges {
		if v == nil {
			continue
		}
		fmt.Println(i, ": ")
		fmt.Println("\tDegree: ", g.Degree[i])
		fmt.Print("\t")
		ptr = v
		for ptr != nil {
			fmt.Printf("(%d, %d) ", i, ptr.Y)
			ptr = ptr.next
		}
		fmt.Print("\n")
	}
}
