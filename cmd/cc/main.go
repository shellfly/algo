/*
  Execution:    go run cmd/cc/main.go filename.txt
  Data files:   https://algs4.cs.princeton.edu/41graph/tinyG.txt
                https://algs4.cs.princeton.edu/41graph/mediumG.txt
                https://algs4.cs.princeton.edu/41graph/largeG.txt

  Compute connected components using depth first search.
  Runs in O(E + V) time.

  % go run cmd/cc/main.go tinyG.txt
  3 components
  0 1 2 3 4 5 6
  7 8
  9 10 11 12

  % pytyon cc.py mediumG.txt
  1 components
  0 1 2 3 4 5 6 7 8 9 10 ...

  % go run cmd/cc/main.go largeG.txt
  1 components
  0 1 2 3 4 5 6 7 8 9 10 ...

  Note: This implementation uses a recursive DFS. To avoid needing
        a potentially very large stack size, replace with a non-recurisve
        DFS ala NonrecursiveDFS.

*/

package main

import (
	"fmt"
	"os"

	"github.com/shellfly/algo"
	"github.com/shellfly/algo/stdin"
)

func main() {
	graph := algo.NewGraph(stdin.NewIn(os.Args[1]))
	cc := algo.NewCC(graph)

	fmt.Println(cc.Count(), " components")
	var components = []*algo.Bag{}
	for i := 0; i < cc.Count(); i++ {
		components = append(components, algo.NewBag())
	}

	for v := 0; v < graph.V(); v++ {
		components[cc.ID(v)].Add(v)
	}

	for i := 0; i < cc.Count(); i++ {
		for _, v := range components[i].Slice() {
			fmt.Print(v, " ")
		}
		fmt.Println()
	}

}
