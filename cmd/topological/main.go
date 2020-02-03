/******************************************************************************
 *  Execution:    go run cmd/topological/main.go filename.txt delimiter
 *  Data files:   https://algs4.cs.princeton.edu/42digraph/jobs.txt
 *
 *  Compute topological ordering of a DAG or edge-weighted DAG.
 *  Runs in O(E + V) time.
 *
 *  % go run cmd/topological/main.go jobs.txt "/"
 *  Calculus
 *  Linear Algebra
 *  Introduction to CS
 *  Advanced Programming
 *  Algorithms
 *  Theoretical CS
 *  Artificial Intelligence
 *  Robotics
 *  Machine Learning
 *  Neural Networks
 *  Databases
 *  Scientific Computing
 *  Computational Biology
 *
 ******************************************************************************/

package main

import (
	"fmt"
	"os"

	"github.com/shellfly/algo"
)

func main() {
	filename, sp := os.Args[1], os.Args[2]
	sg := algo.NewSymbolDigraph(filename, sp)
	top := algo.NewTopological(sg.G())
	for _, v := range top.Order() {
		fmt.Println(sg.Name(v))
	}
}
