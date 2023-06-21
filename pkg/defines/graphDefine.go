package defines

type GraphNode struct {
	ID     int
	Weight int
}

type GraphEdge struct {
	End    int
	Weight int
}

type Graph struct {
	Nodes []*GraphNode
	Edges [][]*GraphEdge
}
