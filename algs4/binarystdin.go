package algs4

import (
	"io"
	"os"
	"strings"
)

// BinaryStdin ...
var BinaryStdin = &binaryStdin{os.Stdin, 0, 0, false, nil}

type binaryStdin struct {
	in            io.Reader
	buffer        int // one character buffer
	n             int // number of bits left in buffer
	isInitialized bool
	err           error
}

func (bs *binaryStdin) fillBuffer() {
	byte := make([]byte, 1)
	_, err := bs.in.Read(byte)
	if err != nil {
		if err == io.EOF {
			bs.err = io.EOF
			bs.n = -1
		} else {
			panic(err)
		}
	}
	bs.n = 8

	bs.buffer = int(byte[0])
}

// ReadBool reads the next bit of data from standard input and return as a boolean.
func (bs *binaryStdin) ReadBool() bool {
	if bs.IsEmpty() {
		panic("reading from empty input stream")
	}
	bs.n--
	bit := (bs.buffer >> bs.n & 1) == 1
	if bs.n == 0 {
		bs.fillBuffer()
	}
	return bit
}

// ReadInt reads the next 32 bits from standard input and return as a 32-bit int.
func (bs *binaryStdin) ReadInt() int {
	if bs.IsEmpty() {
		panic("reading from empty input stream")
	}
	x := 0
	for i := 0; i < 4; i++ {
		b := bs.ReadByte()
		x <<= 8
		x |= int(b)
	}
	return x
}

// ReadIntR reads the next r bits from standard input and return as a 32-bit int.
func (bs *binaryStdin) ReadIntR(r int) int {
	if r < 1 || r > 32 {
		panic("invalid bits r")
	}
	if r == 32 {
		return bs.ReadInt()
	}

	x := 0
	for i := 0; i < r; i++ {
		x <<= 1
		bit := bs.ReadBool()
		if bit {
			x |= 1
		}
	}
	return x
}

// ReadByte the next 8 bits from standard input and return as a byte
func (bs *binaryStdin) ReadByte() byte {
	if bs.IsEmpty() {
		panic("reading from empty input stream")
	}
	if bs.n == 8 {
		b := byte(bs.buffer)
		bs.fillBuffer()
		return b
	}
	x := bs.buffer
	x <<= (8 - bs.n)
	oldN := bs.n
	bs.fillBuffer()
	bs.n = oldN
	x |= (bs.buffer >> bs.n)
	return byte(x)
}

// ReadString reads the remaining bytes of data from standard input and return as a string.
func (bs *binaryStdin) ReadString() string {
	if bs.IsEmpty() {
		panic("reading from empty input stream")
	}

	var sb strings.Builder
	for !bs.IsEmpty() {
		b := bs.ReadByte()
		sb.WriteString(string(b))
	}
	return sb.String()
}

func (bs *binaryStdin) initialize() {
	bs.fillBuffer()
	bs.isInitialized = true
}

// IsEmpty ...
func (bs *binaryStdin) IsEmpty() bool {
	if !bs.isInitialized {
		bs.initialize()
	}
	return bs.err == io.EOF
}
