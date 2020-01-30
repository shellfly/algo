package stdin

import (
	"bufio"
	"os"
	"strings"
)

// StdIn ...
type StdIn struct {
	scanner *bufio.Scanner
}

// NewStdIn ...
func NewStdIn() *StdIn {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)
	return &StdIn{scanner}
}

// IsEmpty reports if the In is empty
func (s *StdIn) IsEmpty() bool {
	return !s.scanner.Scan()
}

// ReadString return next string by delimiter of ' '
func (s *StdIn) ReadString() string {
	return s.scanner.Text()
}

// ReadAllStrings returns all strings as a slice from the input
func ReadAllStrings() (words []string) {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()
		words = append(words, strings.Split(line, " ")...)
	}
	return
}
