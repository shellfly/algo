/******************************************************************************
 *  Execution:    go run cmd/stack/main.go < input.txt
 *
 *  A generic stack or multiset, implemented using a singly linked list.
 *
 *  % more tobe.txt
 *  to be or not to - be - - that - - - is
 *
 *  % go run cmd/stack/main.go < tobe.txt
 *  to be not that or be (2 left on stack)
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
		if !(word == "-") {
			stack.Push(word)
		} else if !stack.IsEmpty() {
			fmt.Print(stack.Pop(), " ")
		}
	}
	fmt.Println("(", stack.Size(), " left on stack)")
}
