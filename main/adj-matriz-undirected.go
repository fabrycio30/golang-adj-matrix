package main

/**import "fmt"
// Go Graph for
// implemented by adjacency matrix in Undirected graph

type Graph struct {
    // Graph node
    node [][] int
    // Number of nodes
    size int
}
func getGraph(size int) * Graph {
    // assume size is valid
    var me *Graph = &Graph {make([][] int, size),size}
    for i := range me.node {
        me.node[i] = make([]int, size)
    }
    return me
}
func(this Graph) addEdge(start, end int) {
    if this.size > start && this.size > end {
        // Set the connection
        this.node[start][end] = 1
        this.node[end][start] = 1
    }
}
func(this Graph) adjacencyNode() {
    if this.size > 0 {
        for row := 0 ; row < this.size ; row++ {
            fmt.Print("Adjacency Matrix of vertex ", row, " :")
            for col := 0 ; col < this.size ; col++ {
                if this.node[row][col] == 1 {
                    // When node is connect
                    fmt.Print(" ", col)
                }
            }
            fmt.Print("\n")
        }
    } else {
        fmt.Println("Empty Graph")
    }
}
func main() {
    // 5 implies the number of nodes in graph
    var g * Graph = getGraph(5)
    //  Add a edge with given nodes
    g.addEdge(0, 1)
    g.addEdge(0, 2)
    g.addEdge(0, 3)
    g.addEdge(1, 4)
    g.addEdge(2, 3)
    g.addEdge(3, 4)
    // Display node and edge status
    g.adjacencyNode()
}**/
