package algo

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBag(t *testing.T) {
	assert := assert.New(t)

	bag := NewBag()
	assert.True(bag.IsEmpty(), "Bag should be empty")
	bag.Add(1)
	bag.Add(2)
	assert.False(bag.IsEmpty(), "Bag should not be empty")
	assert.Equal(bag.Size(), 2, "Bag size should be 2")

	for item := range bag.Slice() {
		fmt.Println(item)
	}
}
