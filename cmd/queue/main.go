/******************************************************************************
 *  Execution:    go run cmd/queue/main.go < input.txt
 *
 *  A generic queue or multiset, implemented using a singly linked list.
 *
 *  % more tobe.txt
 *  to be or not to - be - - that - - - is
 *
 *  % go run cmd/queue/main.go < tobe.txt
 *  size of queue = 14
 *  is
 *  -
 *  -
 *  -
 *  that
 *  -
 *  -
 *  be
 *  -
 *  to
 *  not
 *  or
 *  be
 *  to
 *
 ******************************************************************************/

package main

import (
	"fmt"

	"github.com/shellfly/algo/algs4"
	"github.com/shellfly/algo/stdin"
)

func main() {
	words := stdin.ReadAllStrings()
	queue := algs4.NewQueue()
	for _, word := range words {
		queue.Enqueue(word)
	}

	fmt.Println("size of queue = ", queue.Size())
	for _, item := range queue.Slice() {
		fmt.Println(item)
	}
}
