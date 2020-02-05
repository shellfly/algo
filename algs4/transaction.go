package algs4

import (
	"fmt"
	"strconv"
	"strings"
)

// Transaction get top M elements
type Transaction struct {
	Who, When string
	Amount    float64
}

// NewTransaction ...
func NewTransaction(line string) *Transaction {
	items := strings.Fields(line)
	who, when, amountStr := items[0], items[1], items[2]
	amount, _ := strconv.ParseFloat(amountStr, 64)
	return &Transaction{who, when, amount}
}

// String ...
func (t Transaction) String() string {
	return fmt.Sprintf("%15s %10s %-10.2f", t.Who, t.When, t.Amount)
}

// CompareTo implements PQItem interface
func (t Transaction) CompareTo(other interface{})int{
	tt := other.(Transaction)
	if t.Amount > tt.Amount{
		return 1
	} else if t.Amount < tt.Amount{
		return -1
	} else {
		return 0
	}
}