/******************************************************************************
 *  Execution:    go run cmd/kmp/main.go pattern text
 *
 *  Reads in two strings, the pattern and the input text, and
 *  searches for the pattern in the input text using the
 *  KMP algorithm.
 *
 *  % go run cmd/kmp/main.go abracadabra abacadabrabracabracadabrabrabracad
 *  text:    abacadabrabracabracadabrabrabracad
 *  pattern:               abracadabra
 *
 *  % go run cmd/kmp/main.go rab abacadabrabracabracadabrabrabracad
 *  text:    abacadabrabracabracadabrabrabracad
 *  pattern:         rab
 *
 *  % go run cmd/kmp/main.go bcara abacadabrabracabracadabrabrabracad
 *  text:    abacadabrabracabracadabrabrabracad
 *  pattern:                                   bcara
 *
 *  % go run cmd/kmp/main.go rabrabracad abacadabrabracabracadabrabrabracad
 *  text:    abacadabrabracabracadabrabrabracad
 *  pattern:                        rabrabracad
 *
 *  % go run cmd/kmp/main.go abacad abacadabrabracabracadabrabrabracad
 *  text:    abacadabrabracabracadabrabrabracad
 *  pattern: abacad
 *
 ******************************************************************************/

package main

import (
	"fmt"
	"os"

	"github.com/shellfly/algo/algs4"
)

func main() {
	pat, txt := os.Args[1], os.Args[2]
	kmp := algs4.NewKMP(pat)
	offset := kmp.Search(txt)
	fmt.Println("offset: ", offset)
	fmt.Println("text:    " + txt)
	fmt.Print("pattern: ")
	for i := 0; i < offset; i++ {
		fmt.Print(" ")
	}
	fmt.Println(pat)
}
