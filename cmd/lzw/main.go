/******************************************************************************
 *  Execution:    go run cmd/lzw/main.go - < input.txt   (compress)
 *  Execution:    go run cmd/lzw/main.go + < input.txt   (expand)
 *  Data files:   https://algs4.cs.princeton.edu/55compression/abraLZW.txt
 *                https://algs4.cs.princeton.edu/55compression/ababLZW.txt
 *
 *  Compress or expand binary input from standard input using LZW.
 *
 *  WARNING: STARTING WITH ORACLE JAVA 6, UPDATE 7 the SUBSTRING
 *  METHOD TAKES TIME AND SPACE LINEAR IN THE SIZE OF THE EXTRACTED
 *  SUBSTRING (INSTEAD OF CONSTANT SPACE AND TIME AS IN EARLIER
 *  IMPLEMENTATIONS).
 *
 *  See <a href = "http://java-performance.info/changes-to-string-java-1-7-0_06/">this article</a>
 *  for more details.
 *
 ******************************************************************************/

package main

import (
	"os"

	"github.com/shellfly/algo/algs4"
)

func main() {
	huffman := algs4.NewLZW()
	if os.Args[1] == "-" {
		huffman.Compress()
	} else {
		huffman.Expand()
	}
}
