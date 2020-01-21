package stdin

import (
	"bufio"
	"os"
	"strings"
)

// ReadAllStrings returns all strings as a slice from the input
func ReadAllStrings() (words []string) {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()
		words = append(words, strings.Split(line, " ")...)
	}
	return
}
