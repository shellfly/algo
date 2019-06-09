/******************************************************************************
 *  Execution:    go run bag.go < input.txt
 *
 *  A generic bag or multiset, implemented using a singly linked list.
 *
 *  % more tobe.txt
 *  to be or not to - be - - that - - - is
 *
 *  % go run bag.go < tobe.txt
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

	f, err := os.Open(os.Args[1])
	if err != nil {
		fmt.Println("error open file", os.Args[1], err)
	}
	defer f.Close()

	bag := algo.NewBag()
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		for _, word := range strings.Split(line, " ") {
			bag.Add(word)
		}
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading file", f, err)
	}

	fmt.Println("size of bag = ", bag.Size())
	head := bag.First
	for head != nil {
		fmt.Println(head.Item)
		head = head.Next
	}
}
