package algs4

import (
	"fmt"
	"github.com/shellfly/algo/stdin"
)

type huffmanNode struct {
	ch          byte
	freq        int
	left, right *huffmanNode
}

func (hn *huffmanNode) isLeaf() bool {
	return hn.left == nil && hn.right == nil
}
func (hn *huffmanNode) CompareTo(other interface{}) int {
	that := other.(*huffmanNode)
	if hn.freq < that.freq {
		return -1
	} else if hn.freq > that.freq {
		return 1
	}
	return 0
}
func (hn *huffmanNode) String() string {
	return fmt.Sprintf("%s-%s", hn.ch, hn.freq)
}

// Huffman ...
type Huffman struct {
	R int // ASCII alphabet
}

// NewHuffman ...
func NewHuffman() *Huffman {
	return &Huffman{256}
}

// Compress ...
func (h Huffman) Compress() {
	s := stdin.NewStdIn().ReadString()
	freq := make([]int, h.R)
	for i := 0; i < len(s); i++ {
		freq[s[i]]++
	}
	// build huffman trie
	root := buildTrie(freq)

	// build code table
	st := make([]string, h.R)
	h.buildCode(st, root, "")

	// print trie for decoder
	h.writeTrie(root)

	// print number of bytes in original uncompressed message
	fmt.Println(len(s))

	// use Huffman code to encode input
	for i := 0; i < len(s); i++ {
		code := st[s[i]]
		for j := 0; j < len(code); j++ {
			if string(code[j]) == "0" {
				fmt.Print(0)
			} else if string(code[j]) == "1" {
				fmt.Print(1)
			}
		}
	}
	fmt.Println(s)
}

func (h Huffman) buildTrie(freq []int) *huffmanNode {
	pq := NewMinPQ()
	for c := 0; c < h.R; c++ {
		if freq[c] > 0 {
			pq.Insert(&huffmanNode{byte(c), freq[c], nil, nil})
		}
	}
	for pq.Size() > 1 {
		left := pq.DelMin().(*huffmanNode)
		right := pq.DelMin().(*huffmanNode)
		parent := &huffmanNode{byte(0), left.freq + right.freq, left, right}
		pq.Insert(parent)
	}
	return pq.DelMin().(*huffmanNode)
}
