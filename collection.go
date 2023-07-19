package collection

import (
	"github.com/jkittell/array"
)

type Collection[T any] struct {
	index int
	items *array.Array[T]
}

type Iterator[T any] interface {
	hasNext() bool
	getNext() T
}

func New[T any](n int) *Collection[T] {
	return &Collection[T]{
		items: array.New[T](),
	}
}

func (c *Collection[T]) hasNext() bool {
	if c.index < len(c.items) {
		return true
	}
	return false

}
func (c *Collection[T]) getNext() *Collection[T] {
	if c.hasNext() {
		item := c.items.Lookup(c.index)
		c.index++
		return item
	}
	return nil
}
