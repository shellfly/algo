/******************************************************************************
 *  Execution:    go run cmd/uf/main.go < input.txt
 *
 *  Weighted quick-union by rank with path compression by halving.
 *
 *  % go run cmd/uf/main.go < tinyUF.txt
 *  4 3
 *  3 8
 *  6 5
 *  9 4
 *  2 1
 *  5 0
 *  7 2
 *  6 1
 *  2 components
 *
 ******************************************************************************/

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/shellfly/algo/algs4"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())
	uf := algs4.NewUF(n)
	for scanner.Scan() {
		line := scanner.Text()
		values := strings.Split(line, " ")
		p, _ := strconv.Atoi(values[0])
		q, _ := strconv.Atoi(values[1])
		uf.Union(p, q)
		// fmt.Println(p, q)
	}
	fmt.Println(uf.Count, "components")
}
