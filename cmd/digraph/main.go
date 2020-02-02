/*
   Execution:    go run cmd/digraph/main.go input.txt
   Dependencies: Bag.java Stack.java In.java StdOut.java
   Data files:   https://algs4.cs.princeton.edu/41graph/tinyDG.txt
                 https://algs4.cs.princeton.edu/41graph/mediumDG.txt
                 https://algs4.cs.princeton.edu/41graph/largeDG.txt

   A graph, implemented using an array of sets.
   Parallel edges and self-loops are permitted.

   % go run cmd/digraph/main.go < tinyDG.txt
   13 vertices, 22 edges
   0: 5 1
   1:
   2: 0 3
   3: 5 2
   4: 3 2
   5: 4
   6: 9 4 8 0
   7: 6 9
   8: 6
   9: 11 10
   10: 12
   11: 4 12
   12: 9
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
	fmt.Println(graph)
}
