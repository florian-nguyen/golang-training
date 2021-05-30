package linkedlist

import (
	"errors"
)

type Element struct {
	data int
	next *Element
}

type List struct {
	head *Element
	size int
}

func New(list []int) *List {
	l := List{nil, 0}
	for _, value := range list {
		l.Push(value)
	}
	return &l
}

func (l *List) Size() int {
	return l.size
}

func (l *List) Push(n int) {
	l.head = &Element{n, l.head}
	l.size++
}

func (l *List) Pop() (int, error) {
	if l.size == 0 {
		return 0, errors.New("Error: Impossible to pop empty list!")
	}

	l.size--
	value := l.head.data
	l.head = l.head.next
	return value, nil
}

func (l *List) Array() []int {
	output := make([]int, l.size)
	for i, item := l.size-1, l.head; item != nil; i, item = i-1, item.next {
		output[i] = item.data
	}
	return output
}

func (l *List) Reverse() *List {
	output := New([]int{})
	for item, err := l.Pop(); err == nil; item, err = l.Pop() {
		output.Push(item)
	}
	return output
}
