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
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/shellfly/algo"
)

func main() {
	stack := algo.NewStack()
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()
		for _, word := range strings.Split(line, " ") {
			stack.Push(word)
		}
	}

	fmt.Println("size of stack = ", stack.Size())
	for item := stack.Pop(); !stack.IsEmpty(); item = stack.Pop() {
		fmt.Println(item)
	}
}
