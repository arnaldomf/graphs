package graphs_test

import (
	"github.com/arnaldomf/graphs"
)

func ExampleGraph_AddEdge() {
	g := new(graphs.Graph)
	g.AddEdge(1, 8, 0)
	g.AddEdge(9, 1, 0)
	// Output:
}
