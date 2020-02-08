/******************************************************************************
 *  Execution:    go run cmd/lzw/main.go - < input.txt   (compress)
 *  Execution:    go run cmd/lzw/main.go + < input.txt   (expand)
 *  Data files:   https://algs4.cs.princeton.edu/55compression/abraLZW.txt
 *                https://algs4.cs.princeton.edu/55compression/ababLZW.txt
 *
 *  Compress or expand binary input from standard input using LZW.
 *
 *  % go run cmd/lzw/main.go - < abraLZW.txt | go run cmd/lzw/main.go +
 *  ABRACADBRACRAACRA
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
