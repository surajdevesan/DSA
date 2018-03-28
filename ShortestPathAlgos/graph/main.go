package graph

type Edge struct {
	Src, Dest, Weight int
}

//Graph struct is used for constructing graphs of V vertices and edges
type Graph struct {
	Edges []Edge
	V, E  int
}

func (graphes *Graph) createGraph(V int, E int) {
	graphes.V = V
	graphes.E = E
	graphes.Edges = make([]Edge, E, E)
}

//CreateGraph is used to create graph of V vertices and E edges
func CreateGraph(V int, E int) Graph {
	var graphData Graph
	graphData.createGraph(V, E)
	return graphData
}
