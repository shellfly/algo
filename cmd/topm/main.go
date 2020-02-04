/******************************************************************************
*  Execution:    go run cmd/topm/main.go
 * Data files:   https://algs4.cs.princeton.edu/24pq/tinyBatch.txt
*
*  Given an integer m from the command line and an input stream where
*  each line contains a String and a long value, this MinPQ client
*  prints the m lines whose numbers are the highest.
*
*  % go run cmd/topm/main.go 5 < tinyBatch.txt
*  Thompson    2/27/2000  4747.08
*  vonNeumann  2/12/1994  4732.35
*  vonNeumann  1/11/1999  4409.74
*  Hoare       8/18/1992  4381.21
*  vonNeumann  3/26/2002  4121.85
*
******************************************************************************/

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"

	"github.com/shellfly/algo"
)

func main() {
	m, _ := strconv.Atoi(os.Args[1])
	pq := algo.NewMinPQ()
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()
		pq.Insert(*algo.NewTransaction(line))
		if pq.Size() > m {
			pq.DelMin()
		}
	}

	s := algo.NewStack()
	for !pq.IsEmpty() {
		s.Push(pq.DelMin())
	}
	for !s.IsEmpty() {
		fmt.Println(s.Pop())
	}
}
