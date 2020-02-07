/******************************************************************************
 *  Execution:    go run cmd/nfa/main.go regexp text
 *
 *  % go run cmd/nfa/main.go "(A*B|AC)D" AAAABD
 *  true
 *
 *  % go run cmd/nfa/main.go "(A*B|AC)D" AAAAC
 *  false
 *
 *  % go run cmd/nfa/main.go "(a|(bc)*d)*" abcbcd
 *  true
 *
 *  % go run cmd/nfa/main.go "(a|(bc)*d)*" abcbcbcdaaaabcbcdaaaddd
 *  true
 *
 *  Remarks
 *  -----------
 *  The following features are not supported:
 *    - The + operator
 *    - Multiway or
 *    - Metacharacters in the text
 *    - Character classes.
 *
 ******************************************************************************/

package main

import (
	"fmt"
	"os"

	"github.com/shellfly/algo/algs4"
)

func main() {
	pattern, txt := os.Args[1], os.Args[2]
	nfa := algs4.NewNFA(pattern)
	fmt.Println(nfa.Recognizes(txt))
}
