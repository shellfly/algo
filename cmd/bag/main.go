/******************************************************************************
 *  Execution:    go run cmd/bag/main.go < input.txt
 *
 *  A generic bag or multiset, implemented using a singly linked list.
 *
 *  % more tobe.txt
 *  to be or not to - be - - that - - - is
 *
 *  % go run cmd/bag/main.go < tobe.txt
 *  size of bag = 14
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
	bag := algs4.NewBag()
	for _, word := range words {
		bag.Add(word)

	}

	fmt.Println("size of bag = ", bag.Size())
	for _, item := range bag.StringSlice() {
		fmt.Println(item)
	}
}
