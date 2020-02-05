/*
   Execution:  go run cmd/depth-first-paths/main.go G s

   Data files:   https://algs4.cs.princeton.edu/41graph/tinyCG.txt
                 https://algs4.cs.princeton.edu/41graph/tinyG.txt
                 https://algs4.cs.princeton.edu/41graph/mediumG.txt
                 https://algs4.cs.princeton.edu/41graph/largeG.txt

   Run depth first search on an undirected graph.
   Runs in O(E + V) time.

   %  go run cmd/graph/main.go tinyCG.txt
   6 8
   0: 2 1 5
   1: 0 2
   2: 0 1 3 4
   3: 5 4 2
   4: 3 2
   5: 3 0

   % go run cmd/depth-first-paths/main.go tinyCG.txt 0
   0 to 0:  0
   0 to 1:  0-2-1
   0 to 2:  0-2
   0 to 3:  0-2-3
   0 to 4:  0-2-3-4
   0 to 5:  0-2-3-5

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
	p := algs4.NewDepthFirstPaths(graph, s)

	for v := 0; v < graph.V(); v++ {
		if p.HasPathTo(v) {
			fmt.Printf("%d to %d: ", s, v)
			for _, x := range p.PathTo(v) {
				if x == s {
					fmt.Print(x)
				} else {
					fmt.Printf("-%d", x)
				}
			}
			fmt.Println()
		} else {
			fmt.Printf("%d and %d: not connected\n", s, v)
		}
	}

}
