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
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/shellfly/algo"
)

func main() {
	queue := algo.NewQueue()
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()
		for _, word := range strings.Split(line, " ") {
			queue.Enqueue(word)
		}
	}

	fmt.Println("size of queue = ", queue.Size())
	for item := queue.Dequeue(); !queue.IsEmpty(); item = queue.Dequeue() {
		fmt.Println(item)
	}
}
