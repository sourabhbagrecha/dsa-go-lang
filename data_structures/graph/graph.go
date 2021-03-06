package graph

//Graph Data Structure
type Graph struct {
	adjacencyList map[string][]string
	visited       map[string]bool
	results       []string
}

//AddVertexWithEdges add a new node into the graph along with edges
func (g *Graph) AddVertexWithEdges(vertex string, connections []string) *Graph {
	if g.adjacencyList == nil {
		g.adjacencyList = make(map[string][]string)
	}
	g.adjacencyList[vertex] = connections
	for i := 0; i < len(connections); i++ {
		g.adjacencyList[connections[i]] = append(g.adjacencyList[connections[i]], vertex)
	}
	return g
}

//AddVertex add a new node into the graph
func (g *Graph) AddVertex(vertex string) *Graph {
	if g.adjacencyList == nil {
		g.adjacencyList = make(map[string][]string)
	}
	g.adjacencyList[vertex] = []string{}
	return g
}

//AddEdge add a new connection into the graph between any two nodes
func (g *Graph) AddEdge(vertex1 string, vertex2 string) *Graph {
	if g.adjacencyList == nil {
		g.adjacencyList = make(map[string][]string)
	}

	//Error handling for adding more than one edge with same pair of node is pending

	g.adjacencyList[vertex1] = append(g.adjacencyList[vertex1], vertex2)
	g.adjacencyList[vertex2] = append(g.adjacencyList[vertex2], vertex1)
	return g
}

//RemoveEdge removes a new connection between any two nodes in the graph
func (g *Graph) RemoveEdge(vertex1 string, vertex2 string) *Graph {
	if g.adjacencyList == nil {
		g.adjacencyList = make(map[string][]string)
	}
	for i := 0; i < len(g.adjacencyList[vertex1]); i++ {
		if g.adjacencyList[vertex1][i] == vertex2 {
			g.adjacencyList[vertex1] = append(g.adjacencyList[vertex1][:i], g.adjacencyList[vertex1][i+1:]...)
			break
		}
	}
	for i := 0; i < len(g.adjacencyList[vertex2]); i++ {
		if g.adjacencyList[vertex2][i] == vertex1 {
			g.adjacencyList[vertex2] = append(g.adjacencyList[vertex2][:i], g.adjacencyList[vertex2][i+1:]...)
			break
		}
	}
	return g
}

//RemoveVertex removes a vertex from a graph and also removes all the connections associated with it
func (g *Graph) RemoveVertex(vertex string) *Graph {
	if g.adjacencyList == nil {
		g.adjacencyList = make(map[string][]string)
	}

	for len(g.adjacencyList[vertex]) > 0 {
		g.RemoveEdge(vertex, g.adjacencyList[vertex][0])
	}

	// Different Approach
	// for i := 0; i < len(g.adjacencyList[vertex]); i++ {
	// 	vertexWithConnection := g.adjacencyList[g.adjacencyList[vertex][i]]
	// 	for j := 0; j < len(vertexWithConnection); j++ {
	// 		if vertexWithConnection[j] == vertex {
	// 			g.adjacencyList[g.adjacencyList[vertex][i]] = append(g.adjacencyList[g.adjacencyList[vertex][i]][:j], g.adjacencyList[g.adjacencyList[vertex][i]][j+1:]...)
	// 		}
	// 	}
	// }
	delete(g.adjacencyList, vertex)
	return g
}

func (g *Graph) dfsHelper(start string) {
	currentVertex := g.adjacencyList[start]
	g.results = append(g.results, start)
	if g.visited == nil {
		g.visited = make(map[string]bool)
	}
	g.visited[start] = true
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
func (g *Graph) DFS(start string) []string {
	g.dfsHelper(start)

	//Store the results in temp variable
	results := g.results

	//To reset the visited and results properties of the graph
	g.results = []string{}
	g.visited = make(map[string]bool)

	return results
}

//DFSIterative traverse in a iteravtive rather than the recursive way
func (g *Graph) DFSIterative(start string) []string {
	stack := []string{start}
	g.visited = make(map[string]bool)

	for len(stack) > 0 {
		currentVertex := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		if !g.visited[currentVertex] {
			g.visited[currentVertex] = true
			g.results = append(g.results, currentVertex)

			//Looping in reverse fashion so that the immediate(array order) next neighbor would stay on top after pushing in the array
			for i := len(g.adjacencyList[currentVertex]) - 1; i >= 0; i-- {
				stack = append(stack, g.adjacencyList[currentVertex][i])
			}
		}
	}

	//Store the results in temp variable
	results := g.results

	//To reset the visited and results properties of the graph
	g.results = []string{}
	g.visited = make(map[string]bool)

	return results
}

//BFS traverse the graph in a breadth first search fashion
func (g *Graph) BFS(start string) []string {
	g.visited = make(map[string]bool)
	queue := []string{start}
	var results []string
	for len(queue) > 0 {
		if !g.visited[queue[0]] {
			g.results = append(g.results, queue[0])
			g.visited[queue[0]] = true
			currentVertex := g.adjacencyList[queue[0]]
			queue = append(queue, currentVertex...)
		}
		queue = queue[1:]
		results = g.results
	}
	g.results = []string{}
	g.visited = make(map[string]bool)
	return results
}
