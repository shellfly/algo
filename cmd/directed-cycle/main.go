/*
  Execution:    go run cmd/directed-cycle/main.go input.txt
  Data files:   https://algs4.cs.princeton.edu/42digraph/tinyDG.txt
                https://algs4.cs.princeton.edu/42digraph/tinyDAG.txt

  Finds a directed cycle in a digraph.
  Runs in O(E + V) time.

  % go run cmd/directed-cycle/main.go tinyDG.txt
  Directed cycle: 3 5 4 3

  %  go run cmd/directed-cycle/main.go tinyDAG.txt
  No directed cycle

*/

package main

import (
	"fmt"
	"os"

	"github.com/shellfly/algo"
	"github.com/shellfly/algo/stdin"
)

func main() {
	graph := algo.NewDigraph(stdin.NewIn(os.Args[1]))
	finder := algo.NewDirectedCycle(graph)
	if finder.HasCycle() {
		for _, v := range finder.Cycle().Slice() {
			fmt.Print(v, " ")
		}
		fmt.Println()
	} else {
		fmt.Println("Graph is acyclic")
	}

}
