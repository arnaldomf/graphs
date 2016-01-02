/*
Package graphs defines a graph type and algorithms to traverse it
*/
package graphs

import "fmt"

// Edgenode represents an edge on an adjacency list
type Edgenode struct {
	Y      int
	Weight int
	next   *Edgenode
}

// Graph is an adjacency list representing a graph
type Graph struct {
	Edges     []*Edgenode
	Degree    []int
	MAXV      int
	NVertices int
	NEdges    int
	Directed  bool
}

// New creates a new Graph with maxvertices alocated
func New(maxvertices int, direct bool) (*Graph, error) {
	if maxvertices <= 0 {
		return nil, fmt.Errorf("graphs.New: maxvertices less than 0")
	}
	g := &Graph{}
	g.Directed = direct
	g.MAXV = maxvertices
	for i := 0; i <= maxvertices; i++ {
		g.Edges = append(g.Edges, nil)
		g.Degree = append(g.Degree, 0)
	}
	return g, nil
}

func (g *Graph) auxiliaryAddEdge(x, y, weight int) {
	auxEdge := new(Edgenode)
	if g.Edges[x] == nil {
		g.Edges[x] = auxEdge
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
}

// AddEdge adds a new edge to Graph structure, if weight = 0 it will be ignored
func (g *Graph) AddEdge(x, y, weight int) {
	if x >= g.MAXV || y >= g.MAXV {
		errMessage := fmt.Sprintf("x and y must be < %d", g.MAXV)
		panic(errMessage)
	}
	g.auxiliaryAddEdge(x, y, weight)
	if !g.Directed {
		g.auxiliaryAddEdge(y, x, weight)
	}
	g.Degree[x]++
	g.NVertices++
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
