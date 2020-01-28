package stdin

import (
	"bufio"
	"fmt"
	"os"
)

// In wraps a scanner using ScanWords as split function
type In struct {
	scanner *bufio.Scanner
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
