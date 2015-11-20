package graphs

//GraphProcessor is used to process the graph
type GraphProcessor interface {
	ProcessEdge(int, int)
	ProcessVertexEarly(int)
	ProcessVertexLate(int)
}

var (
	discovered [MAXV + 1]bool
	processed  [MAXV + 1]bool
	parent     [MAXV + 1]int
)

func (g *Graph) initializeSearch() {
	for i := 1; i < g.NVertices; i++ {
		processed[i] = false
		discovered[i] = false
		parent[i] = -1
	}
}

/*BreadthFirstSearch algorithm for graph traverlas, where each level is explored
before getting to the next one. "start" must be part of Graph, otherise panics
*/
func (g *Graph) BreadthFirstSearch(start int, gp GraphProcessor) {
	if g.Edges[start] == nil {
		panic("BreadthFirstSearch: Cant start from dead end")
	}
	var p *Edgenode
	var v int
	var y int
	g.initializeSearch()
	queue := make(chan int, g.NEdges+1)
	queue <- start
	discovered[start] = true
	for len(queue) > 0 {
		v = <-queue
		gp.ProcessVertexEarly(v)
		processed[v] = true
		p = g.Edges[v]
		for p != nil {
			y = p.Y
			if !processed[y] || g.Directed {
				gp.ProcessEdge(v, y)
			}
			if !discovered[y] {
				queue <- y
				discovered[y] = true
				parent[y] = v
			}
			p = p.next
		}
		gp.ProcessVertexLate(v)
	}
}
