package algs4

import (
	"fmt"
	"io"
	"os"
)

// BinaryStdout ...
var BinaryStdout = binaryStdout{os.Stdout, 0, 0}

type binaryStdout struct {
	out    io.Writer
	buffer int // 8-bit buffer of bits to write
	n      int // number of bits remaining in buffer
}
// WriteInt writes the 32-bit int to standard output
func (bs *binaryStdout)WriteInt(x int)error{
	bs.WriteByte(byte((x >> 24) & 0xff))
	bs.WriteByte(byte((x >> 16) & 0xff))
	bs.WriteByte(byte((x >>  8) & 0xff))
	bs.WriteByte(byte((x >>  0) & 0xff))
	return nil
}

// Close ...
func (bs *binaryStdout)Close(){
	bs.clearBuffer()
}

// Write...
func (bs *binaryStdout) Write(p []byte) error {
	return nil
}


// WriteByte writes the 8-bit byte to standard output.
func (bs *binaryStdout) WriteByte(b byte) error {
	if !(int(b) >= 0 && int(b) < 256) {
		panic("invalid byte")
	}
	// // optimized if byte-aligned
	// if bs.n == 0 {
	// 	_, err := bs.out.Write([]byte{b})
	// 	if err != nil {
	// 		fmt.Println("write byte to stdout error: ", err)
	// 		return err
	// 	}
	// }
	for i := 0; i < 8; i++ {
		bit := ((int(b) >> (8 - i - 1)) & 1) == 1
		bs.WriteBit(bit)
	}
	return nil
}

// WriteBit writes the specified bit to standard output.
func (bs *binaryStdout) WriteBit(bit bool) error {
	bs.buffer <<= 1
	if bit {
		bs.buffer |= 1
	}
	bs.n++
	if bs.n == 8 {
		return bs.clearBuffer()
	}
	return nil
}

func (bs *binaryStdout) clearBuffer() error {
	if bs.n == 0 {
		return nil
	}
	if bs.n > 0 {
		bs.buffer <<= (8 - bs.n)
	}
	_, err := bs.out.Write([]byte{byte(bs.buffer)})
	if err != nil {
		fmt.Println("write byte to stdout error: ", err)
		panic("err")
	}
	bs.n = 0
	bs.buffer = 0
	return nil
}