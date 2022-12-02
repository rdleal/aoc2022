package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type heap[T any] struct {
	items   []T
	cmpFunc func(a, b T) bool
}

func newHeap[T any](cmp func(a, b T) bool) *heap[T] {
	return &heap[T]{
		items:   make([]T, 0),
		cmpFunc: cmp,
	}
}

func (h *heap[T]) DelTop() T {
	if len(h.items) == 0 {
		panic("empty heap")
	}
	v := h.items[0]
	l := len(h.items)
	h.items[0] = h.items[l-1]
	h.items = h.items[:l-1]
	h.heapifyDown(0)

	return v
}

func (h *heap[T]) leftChild(i int) int {
	return i*2 + 1
}

func (h *heap[T]) rightChild(i int) int {
	return i*2 + 2
}

func (h *heap[T]) heapifyDown(i int) {
	for j := h.leftChild(i); j < len(h.items); j = h.leftChild(i) {
		if r := h.rightChild(i); r < len(h.items) && h.cmpFunc(h.items[r], h.items[j]) {
			j = r
		}

		if !h.cmpFunc(h.items[j], h.items[i]) {
			break
		}

		h.items[i], h.items[j] = h.items[j], h.items[i]
		i = j
	}
}

func (h *heap[T]) Insert(v T) {
	h.items = append(h.items, v)
	h.heapifyUp(len(h.items) - 1)
}

func (h *heap[T]) parentNode(i int) int {
	return (i - 1) / 2
}

func (h *heap[T]) heapifyUp(i int) {
	for j := h.parentNode(i); h.cmpFunc(h.items[i], h.items[j]); j = h.parentNode(i) {
		h.items[i], h.items[j] = h.items[j], h.items[i]
		i = j
	}
}

func main() {
	h := newHeap(func(i, j int) bool { return i > j })

	var max, cals int // max = answer for part one

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		calStr := scanner.Text()
		if calStr == "" {
			h.Insert(cals)
			cals = 0
		}

		cal, _ := strconv.Atoi(calStr)

		cals += cal

		if cals > max {
			max = cals
		}
	}

	var res int // res = answer for part two
	if len(h.items) > 0 {
		res = h.DelTop() + h.DelTop() + h.DelTop()
	}

	fmt.Println(max, res)
}
