/*
  Execution:    go run cmd/depth-first-search/main.go filename.txt s
  Data files:   https: // algs4.cs.princeton.edu / 41graph / tinyG.txt
                https: // algs4.cs.princeton.edu / 41graph / mediumG.txt

  Run depth first search on an undirected graph.
  Runs in O(E + V) time.

 % go run cmd/depth-first-search/main.go tinyG.txt 0
  0 1 2 3 4 5 6
  NOT connected

 % go run cmd/depth-first-search/main.go tinyG.txt 9
  9 10 11 12
  NOT connected
*/

package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/shellfly/algo/algs4"
	"github.com/shellfly/algo/stdin"
)

func main() {
	graph := algs4.NewGraph(stdin.NewIn(os.Args[1]))
	s, _ := strconv.Atoi(os.Args[2])
	search := algs4.NewDepthFirstSearch(graph, s)
	for v := 0; v < graph.V(); v++ {
		if search.Marked(v) {
			fmt.Println(v, " ")
		}
	}
	if search.Count() == graph.V() {
		fmt.Println("connected")
	} else {
		fmt.Println("not connected")
	}

}
