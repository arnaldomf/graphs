package graphs_test

import (
	"fmt"
	"github.com/arnaldomf/graphs"
)

func Example() {
	gp := &graphs.GraphProcessor{
		ProcessEdge: func(x, y int) {
			fmt.Printf("processed edge (%d, %d)\n", x, y)
		},
		ProcessVertexLate: func(x int) {
		},
		ProcessVertexEarly: func(x int) {
			fmt.Printf("Processed vertex %d\n", x)
		},
	}

	g, err := graphs.New(10, false)
	if err != nil {
		panic(err)
	}
	g.AddEdge(1, 8, 0)
	g.AddEdge(1, 9, 0)
	g.AddEdge(5, 3, 0)
	g.AddEdge(3, 9, 0)
	g.AddEdge(5, 1, 0)
	g.BreadthFirstSearch(1, gp)

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
