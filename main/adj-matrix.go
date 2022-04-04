package main

import "fmt"

// Go Graph for
// Adjacency matrix in directed graph
type Graph struct {
	// Graph node
	node [][]int
	// Number of nodes
	size int
}

func getGraph(size int) *Graph {
	// Assume size is valid
	var me *Graph = &Graph{make([][]int, size), size}
	for i := range me.node {
		me.node[i] = make([]int, size)
	}
	return me
}
func (this Graph) addEdge(start, end int) {
	if this.size > start && this.size > end {
		// Set the connection
		this.node[start][end] = 1
	}
}
func (this Graph) adjacencyNode() {
	if this.size > 0 {

		for row := 0; row < this.size; row++ {
			fmt.Print("", row, " : ")
			for col := 0; col < this.size; col++ {
				if this.node[row][col] == 1 {
					fmt.Print("[1] ") //fmt.Print("[", col, "] ")
				} else {
					fmt.Print("[0] ")
				}
			}
			fmt.Print("\n")
		}
	} else {
		fmt.Println("Empty Graph")
	}
}
func main() {
	var g *Graph = getGraph(6)
	//  adicinanda arestas
	g.addEdge(0, 1)
	g.addEdge(0, 3)
	g.addEdge(2, 1)
	g.addEdge(2, 3)
	g.addEdge(3, 4)
	g.addEdge(4, 2)
	g.addEdge(4, 5)
	g.addEdge(5, 2)

	// mostrando a matriz de adjancências
	g.adjacencyNode()
}
