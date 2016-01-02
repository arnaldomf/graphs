package graphs

//GraphProcessor is used to process the graph
type GraphProcessor struct {
	ProcessEdge        func(int, int)
	ProcessVertexEarly func(int)
	ProcessVertexLate  func(int)
}

var (
	discovered = make(map[*Graph][]bool)
	processed  = make(map[*Graph][]bool)
	parent     = make(map[*Graph][]int)
)

func (g *Graph) initializeSearch() {
	if processed[g] == nil {
		processed[g] = make([]bool, g.MAXV)
	}
	if discovered[g] == nil {
		discovered[g] = make([]bool, g.MAXV)
	}
	if parent[g] == nil {
		parent[g] = make([]int, g.MAXV)
	}
	for i := 1; i < g.NVertices; i++ {
		processed[g][i] = false
		discovered[g][i] = false
		parent[g][i] = -1
	}
}

/*BreadthFirstSearch algorithm for graph traversal, where each level is explored
before getting to the next one. "start" must be part of Graph, otherise panics
*/
func (g *Graph) BreadthFirstSearch(start int, gp *GraphProcessor) {
	if g.Edges[start] == nil {
		panic("BreadthFirstSearch: Cant start from dead end")
	}
	var p *Edgenode
	var v int
	var y int
	g.initializeSearch()
	queue := make(chan int, g.NVertices+1)
	queue <- start
	discovered[g][start] = true
	for len(queue) > 0 {
		v = <-queue
		gp.ProcessVertexEarly(v)
		processed[g][v] = true
		p = g.Edges[v]
		for p != nil {
			y = p.Y
			if !processed[g][y] || g.Directed {
				gp.ProcessEdge(v, y)
			}
			if !discovered[g][y] {
				queue <- y
				discovered[g][y] = true
				parent[g][y] = v
			}
			p = p.next
		}
		gp.ProcessVertexLate(v)
	}
}
