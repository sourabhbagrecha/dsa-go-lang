package graph

//Graph Data Structure
type Graph struct {
	adjacenceyList map[string][]string
}

//AddVertexWithEdges add a new node into the graph along with edges
func (g *Graph) AddVertexWithEdges(vertex string, connections []string) *Graph {
	if g.adjacenceyList == nil {
		g.adjacenceyList = make(map[string][]string)
	}
	g.adjacenceyList[vertex] = connections
	for i := 0; i < len(connections); i++ {
		g.adjacenceyList[connections[i]] = append(g.adjacenceyList[connections[i]], vertex)
	}
	return g
}

//AddVertex add a new node into the graph
func (g *Graph) AddVertex(vertex string) *Graph {
	if g.adjacenceyList == nil {
		g.adjacenceyList = make(map[string][]string)
	}
	g.adjacenceyList[vertex] = []string{}
	return g
}

//AddEdge add a new connection into the graph between any two nodes
func (g *Graph) AddEdge(vertex1 string, vertex2 string) *Graph {
	if g.adjacenceyList == nil {
		g.adjacenceyList = make(map[string][]string)
	}

	//Error handling for adding more than one edge with same pair of node is pending

	g.adjacenceyList[vertex1] = append(g.adjacenceyList[vertex1], vertex2)
	g.adjacenceyList[vertex2] = append(g.adjacenceyList[vertex2], vertex1)
	return g
}

//RemoveEdge removes a new connection between any two nodes in the graph
func (g *Graph) RemoveEdge(vertex1 string, vertex2 string) *Graph {
	if g.adjacenceyList == nil {
		g.adjacenceyList = make(map[string][]string)
	}
	for i := 0; i < len(g.adjacenceyList[vertex1]); i++ {
		if g.adjacenceyList[vertex1][i] == vertex2 {
			g.adjacenceyList[vertex1] = append(g.adjacenceyList[vertex1][:i], g.adjacenceyList[vertex1][i+1:]...)
			break
		}
	}
	for i := 0; i < len(g.adjacenceyList[vertex2]); i++ {
		if g.adjacenceyList[vertex2][i] == vertex1 {
			g.adjacenceyList[vertex2] = append(g.adjacenceyList[vertex2][:i], g.adjacenceyList[vertex2][i+1:]...)
			break
		}
	}
	return g
}

//RemoveVertex removes a vertex from a graph and also removes all the connections associated with it
func (g *Graph) RemoveVertex(vertex string) *Graph {
	if g.adjacenceyList == nil {
		g.adjacenceyList = make(map[string][]string)
	}

	for len(g.adjacenceyList[vertex]) > 0 {
		g.RemoveEdge(vertex, g.adjacenceyList[vertex][0])
	}

	// Different Solution
	// for i := 0; i < len(g.adjacenceyList[vertex]); i++ {
	// 	vertexWithConnection := g.adjacenceyList[g.adjacenceyList[vertex][i]]
	// 	for j := 0; j < len(vertexWithConnection); j++ {
	// 		if vertexWithConnection[j] == vertex {
	// 			g.adjacenceyList[g.adjacenceyList[vertex][i]] = append(g.adjacenceyList[g.adjacenceyList[vertex][i]][:j], g.adjacenceyList[g.adjacenceyList[vertex][i]][j+1:]...)
	// 		}
	// 	}
	// }
	delete(g.adjacenceyList, vertex)
	return g
}
