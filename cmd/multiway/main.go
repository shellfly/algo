/******************************************************************************
*  Execution:  go run cmd/multiway/main.go
*  Data files:   https://algs4.cs.princeton.edu/24pq/m1.txt
*                https://algs4.cs.princeton.edu/24pq/m2.txt
*                https://algs4.cs.princeton.edu/24pq/m3.txt
*
*  Merges together the sorted input stream given as command-line arguments
*  into a single sorted output stream on standard output.
*
*  % more m1.txt
*  A B C F G I I Z
*
*  % more m2.txt
*  B D H P Q Q
*
*  % more m3.txt
*  A B E F J N
*
*  % go run cmd/multiway/main.go m1.txt m2.txt m3.txt
*  A A B B B C D E F F G H I I J N P Q Q Z
*
******************************************************************************/

package main

import (
	"fmt"
	"os"

	"github.com/shellfly/algo/algs4"
	"github.com/shellfly/algo/stdin"
)

type StringPQItem string

func (s StringPQItem) CompareTo(other interface{}) int {
	ss := other.(StringPQItem)
	if s < ss {
		return -1
	} else if s > ss {
		return 1
	} else {
		return 0
	}
}
func (s StringPQItem) String() string {
	return string(s)
}

func merge(streams []*stdin.In) {
	n := len(streams)
	pq := algs4.NewIndexMinPQ(n)
	for i := 0; i < n; i++ {
		if !streams[i].IsEmpty() {
			str := streams[i].ReadString()
			pq.Insert(i, StringPQItem(str))
		}
	}
	for !pq.IsEmpty() {
		fmt.Print(pq.Min(), " ")
		i := pq.DelMin()
		if !streams[i].IsEmpty() {
			pq.Insert(i, StringPQItem(streams[i].ReadString()))
		}
	}
}
func main() {
	args := os.Args[1:]
	m := len(args)
	streams := make([]*stdin.In, m)
	for i := 0; i < m; i++ {
		streams[i] = stdin.NewIn(args[i])
	}
	merge(streams)
}
