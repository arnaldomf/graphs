package graphs_test

import (
	"github.com/arnaldomf/graphs"
)

func ExampleGraph_AddEdge() {
	g, err := graphs.New(10, false)
	if err != nil {
		panic(err)
	}
	g.AddEdge(1, 8, 0)
	g.AddEdge(9, 1, 0)
	// Output:
}
