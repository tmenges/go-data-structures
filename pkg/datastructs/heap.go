package datastructs

import "errors"

type Heap struct {
	data []int
}

func NewHeap() Heap {
	h := Heap{}
	h.data = append(h.data, 0)

	return h
}

func NewHeapFrom(source []int) Heap {
	h := Heap{}
	h.data = append(h.data, 0)
	h.data = append(h.data, source...)

	return h
}

func (h Heap) Len() int {
	return len(h.data) - 1
}

func (h *Heap) Push(val int) {
	h.data = append(h.data, val)
}

func (h *Heap) Pop() (int, error) {
	if len(h.data) <= 1 {
		return 0, errors.New("heap is empty")
	}

	val := h.data[1]
	h.data = h.data[1:]

	return val, nil
}
