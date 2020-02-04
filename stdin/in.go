package stdin

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

// In wraps a scanner using ScanWords as split function
type In struct {
	scanner *bufio.Scanner
}

// NewInByLine return a pointer of In
func NewInByLine(path string) *In {
	inFile, err := os.Open(path)
	if err != nil {
		fmt.Println(err)
		panic("Can not open file: " + path)
	}
	scanner := bufio.NewScanner(inFile)
	return &In{scanner}
}

// NewIn return a pointer of In
func NewIn(path string) *In {
	inFile, err := os.Open(path)
	if err != nil {
		fmt.Println(err)
		panic("Can not open file: " + path)
	}
	scanner := bufio.NewScanner(inFile)
	scanner.Split(bufio.ScanWords)
	return &In{scanner}
}

// IsEmpty reports if the In is empty
func (in In) IsEmpty() bool {
	return !in.scanner.Scan()
}

// ReadString return next string by delimiter of ' '
func (in In) ReadString() string {

	return in.scanner.Text()
}

// ReadInt return next integer by delimiter of ' '
func (in In) ReadInt() int {
	in.scanner.Scan()
	v, _ := strconv.Atoi(in.scanner.Text())
	return v
}

// ReadFloat32 return next integer by delimiter of ' '
func (in In) ReadFloat32() float32 {
	in.scanner.Scan()
	v, _ := strconv.ParseFloat(in.scanner.Text(), 10)
	return float32(v)
}
