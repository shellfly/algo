/******************************************************************************
 *  Execution:    go run cmd/stack/main.go < input.txt
 *
 *  A generic stack or multiset, implemented using a singly linked list.
 *
 *  % more tobe.txt
 *  to be or not to - be - - that - - - is
 *
 *  % go run cmd/stack/main.go < tobe.txt
 *  size of stack = 14
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
	stack := algs4.NewStack()
	for _, word := range words {
		stack.Push(word)
	}

	fmt.Println("size of stack = ", stack.Size())
	for item := stack.Pop(); !stack.IsEmpty(); item = stack.Pop() {
		fmt.Println(item)
	}
}
