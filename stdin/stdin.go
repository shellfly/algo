package stdin

import (
	"bufio"
	"os"
	"strings"
)

// ReadAllStrings returns all strings as a slice from the input
func ReadAllStrings() []string {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	line := scanner.Text()
	return strings.Split(line, " ")
}
