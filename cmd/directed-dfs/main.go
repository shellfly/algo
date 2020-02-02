/*
   Execution:    go run cmd/directed-dfs/main.go digraph.txt s
   Data files:   https: // algs4.cs.princeton.edu / 42digraph / tinyDG.txt
                 https://algs4.cs.princeton.edu/42digraph/mediumDG.txt
                 https://algs4.cs.princeton.edu/42digraph/largeDG.txt

   Determine single-source or multiple-source reachability in a digraph
   using depth first search.
   Runs in O(E + V) time.

   % go run cmd/directed-dfs/main.go tinyDG.txt 1
   1

   % go run cmd/directed-dfs/main.go tinyDG.txt 2
   0 1 2 3 4 5

   % go run cmd/directed-dfs/main.go tinyDG.txt 1 2 6
   0 1 2 3 4 5 6 8 9 10 11 12
*/

package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/shellfly/algo"
	"github.com/shellfly/algo/stdin"
)

func main() {
	graph := algo.NewDigraph(stdin.NewIn(os.Args[1]))

	sources := []int{}
	for i := 2; i < len(os.Args); i++ {
		s, _ := strconv.Atoi(os.Args[i])
		sources = append(sources, s)
	}
	reachable := algo.NewDirectedDFSSources(graph, sources)
	for v := 0; v < graph.V(); v++ {
		if reachable.Marked(v) {
			fmt.Println(v, " ")
		}
	}

}
