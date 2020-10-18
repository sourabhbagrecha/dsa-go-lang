package graph

//Graph Data Structure
type Graph struct {
	adjacenceyList map[string][]string
	visited        map[string]bool
	results        []string
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

	// Different Approach
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

func (g *Graph) dfsHelper(vertex string) {
	currentVertex := g.adjacenceyList[vertex]
	g.results = append(g.results, vertex)
	if g.visited == nil {
		g.visited = make(map[string]bool)
	}
	g.visited[vertex] = true
	if len(currentVertex) == 0 {
		return
	}
	for i := 0; i < len(currentVertex); i++ {
		if !g.visited[currentVertex[i]] {
			g.dfsHelper(currentVertex[i])
		}
	}
	return
}

//DFS depth first traversal of the complete graph
func (g *Graph) DFS(vertex string) []string {
	g.dfsHelper(vertex)

	//Store the results in temp variable
	results := g.results

	//To reset the visited and results properties of the graph
	g.results = []string{}
	g.visited = make(map[string]bool)

	return results
}

//DFSIterative traverse in a iteravtive rather than the recursive way
func (g *Graph) DFSIterative(vertex string) []string {
	stack := []string{vertex}
	g.visited = make(map[string]bool)

	for len(stack) > 0 {
		currentVertex := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		if !g.visited[currentVertex] {
			g.visited[currentVertex] = true
			g.results = append(g.results, currentVertex)

			//Looping in reverse fashion so that the immediate(array order) next neighbor would stay on top after pushing in the array
			for i := len(g.adjacenceyList[currentVertex]) - 1; i >= 0; i-- {
				stack = append(stack, g.adjacenceyList[currentVertex][i])
			}
		}
	}
	return g.results
}
