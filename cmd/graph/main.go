/*
   Execution:    go run cmd/graph/main.go input.txt
   Dependencies: Bag.java Stack.java In.java StdOut.java
   Data files:   https://algs4.cs.princeton.edu/41graph/tinyG.txt
                 https://algs4.cs.princeton.edu/41graph/mediumG.txt
                 https://algs4.cs.princeton.edu/41graph/largeG.txt

   A graph, implemented using an array of sets.
   Parallel edges and self-loops allowed.

   % go run cmd/graph/main.go < tinyG.txt
   13 vertices, 13 edges
   0: 6 2 1 5
   1: 0
   2: 0
   3: 5 4
   4: 5 6 3
   5: 3 4 0
   6: 0 4
   7: 8
   8: 7
   9: 11 10 12
   10: 9
   11: 9 12
   12: 11 9

   % go run cmd/graph/main.go < mediumG.txt
   250 vertices, 1273 edges
   0: 225 222 211 209 204 202 191 176 163 160 149 114 97 80 68 59 58 49 44 24 15
   1: 220 203 200 194 189 164 150 130 107 72
   2: 141 110 108 86 79 51 42 18 14
   ...
*/

package main

import (
	"fmt"
	"os"

	"github.com/shellfly/algo/algs4"
	"github.com/shellfly/algo/stdin"
)

func main() {
	graph := algs4.NewGraph(stdin.NewIn(os.Args[1]))
	fmt.Println(graph)
}
