package algs4

import (
	"fmt"
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
	return fmt.Sprintf("%c-%d-(%v-%v)", hn.ch, hn.freq, hn.left, hn.right)
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
	s := BinaryStdin.ReadString()
	freq := make([]int, h.R)
	for i := 0; i < len(s); i++ {
		freq[s[i]]++
	}
	// build huffman trie
	root := h.buildTrie(freq)

	// build code table
	st := make([]string, h.R)
	h.buildCode(st, root, "")
	// print trie for decoder
	h.writeTrie(root)

	// print number of bytes in original uncompressed message
	BinaryStdout.WriteInt(len(s))

	// use Huffman code to encode input
	for i := 0; i < len(s); i++ {
		code := st[s[i]]
		for j := 0; j < len(code); j++ {
			if string(code[j]) == "0" {
				BinaryStdout.WriteBit(false)
			} else if string(code[j]) == "1" {
				BinaryStdout.WriteBit(true)
			}
		}
	}
	BinaryStdout.Close()
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

func (h Huffman) writeTrie(x *huffmanNode) {
	if x.isLeaf() {
		BinaryStdout.WriteBit(true)
		BinaryStdout.WriteByte(x.ch)
		return
	}
	BinaryStdout.WriteBit(false)
	h.writeTrie(x.left)
	h.writeTrie(x.right)
}

func (h Huffman) buildCode(st []string, x *huffmanNode, s string) {
	if !x.isLeaf() {
		h.buildCode(st, x.left, s+"0")
		h.buildCode(st, x.right, s+"1")
	} else {
		st[x.ch] = s
	}
}

// Expand ...
func (h Huffman) Expand() {
	root := readTrie()
	length := BinaryStdin.ReadInt()
	for i := 0; i < length; i++ {
		x := root
		for !x.isLeaf() {
			bit := BinaryStdin.ReadBool()
			if bit {
				x = x.right
			} else {
				x = x.left
			}
		}
		BinaryStdout.WriteByte(x.ch)
	}
}

func readTrie() *huffmanNode {
	isLeaf := BinaryStdin.ReadBool()
	if isLeaf {
		return &huffmanNode{BinaryStdin.ReadByte(), -1, nil, nil}
	}
	return &huffmanNode{'0', -1, readTrie(), readTrie()}
}
