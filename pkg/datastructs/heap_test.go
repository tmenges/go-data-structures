package datastructs

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewHeap(t *testing.T) {
	source := []int{1, 3, 5, 7, 9}
	h := NewHeapFrom(source)
	assert.Equal(t, h.Len(), len(source), "New heap length should be the same as the source")
}

func TestEmpty(t *testing.T) {
	h := NewHeap()

	assert.Equal(t, 0, h.Len(), "Empty Heap should have zero length")

	_, err := h.Pop()
	assert.NotNil(t, err, "Popping from an empty Heap should return an error")

	h.Push(100)
	assert.Equal(t, 1, h.Len(), "Pushing one value onto an empty Heap should make the length one")

	val, err := h.Pop()
	assert.Nil(t, err, "No error should be returned")
	assert.Equal(t, 100, val, "Expected value just pushed onto the Heap")
	assert.Equal(t, 0, h.Len(), "Heap should now be empty")
}

func TestSimplePushPop(t *testing.T) {
	source := []int{1, 3, 5, 7, 9}
	h := NewHeapFrom(source)

	val, err := h.Pop()

	assert.Nil(t, err, "No error should be returned")
	assert.Equal(t, source[0], val, "Expected first value from source")
}
