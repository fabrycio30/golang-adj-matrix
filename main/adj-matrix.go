package main

import (
	"fmt"
	"strings"
)

type Graph struct {
	node [][]int
	size int
}

func getGraph(size int) *Graph {
	var me *Graph = &Graph{make([][]int, size), size}
	for i := range me.node {
		me.node[i] = make([]int, size)
	}
	return me
}

//Adicionando arestas para grafos direcionados
func (this Graph) addEdgeDirected(u, v int) {
	if this.size > u && this.size > v {
		// Set the connection
		this.node[u][v] = 1
	}
}

//Adicionando arestas para grafos não direcionados
func (this Graph) addEdge(u, v int) {
	if this.size > u && this.size > v {
		// Set the connection
		this.node[u][v] = 1
		this.node[v][u] = 1
	}
}
func (this Graph) adjacencyNode() {
	if this.size > 0 {
		for row := 0; row < this.size; row++ {
			fmt.Print("", row, " : ")
			for col := 0; col < this.size; col++ {
				if this.node[row][col] == 1 {
					fmt.Print("[x] ")
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
func adjacencyNodeText(texto string) {
	if texto != "" {
		var t string = strings.ReplaceAll(texto, " ", "") //remove espaços em branco
		var t1 []string = strings.Split(t, "{")           //remove '{'
		//var tfinal string = strings.ReplaceAll(string((t[1]), "}", "")//remove final
		var g string = strings.Replace(t1[1], "}", "", 1)
		fmt.Println(g)
		var rows []string = strings.Split(g, "\n")
		fmt.Println(rows[1])
	} else {
		fmt.Println("Empty Graph")
	}
}
func main() {

	var g *Graph = getGraph(5) //Defina o tamanho de Vértices
	//  adicinando arestas
	g.addEdge(0, 1)
	g.addEdge(2, 1)
	//mostrando a matriz de adjancências
	g.adjacencyNode()

}
