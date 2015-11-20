package graphs_test

import (
	"fmt"
	"github.com/arnaldomf/graphs"
)

type myprocessor int

func (m myprocessor) ProcessEdge(x, y int) {
	fmt.Printf("processed edge (%d, %d)\n", x, y)
}
func (m myprocessor) ProcessVertexLate(x int) {

}
func (m myprocessor) ProcessVertexEarly(x int) {
	fmt.Printf("Processed vertex %d\n", x)
}

func Example() {
	var mp myprocessor
	g := new(graphs.Graph)
	g.AddEdge(1, 8, 0)
	g.AddEdge(1, 9, 0)
	g.AddEdge(5, 3, 0)
	g.AddEdge(3, 9, 0)
	g.AddEdge(5, 1, 0)
	g.BreadthFirstSearch(1, mp)

	// Output:
	// Processed vertex 1
	// processed edge (1, 8)
	// processed edge (1, 9)
	// processed edge (1, 5)
	// Processed vertex 8
	// Processed vertex 9
	// processed edge (9, 3)
	// Processed vertex 5
	// processed edge (5, 3)
	// Processed vertex 3
}
