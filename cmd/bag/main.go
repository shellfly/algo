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
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/shellfly/algo"
)

func main() {
	bag := algo.NewBag()
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()
		for _, word := range strings.Split(line, " ") {
			bag.Add(word)
		}
	}

	fmt.Println("size of bag = ", bag.Size())
	for _, item := range bag.Slice() {
		fmt.Println(item)
	}
}
