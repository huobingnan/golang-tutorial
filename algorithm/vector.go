package test

import "errors"

// go version minimum 1.18
//lint:file-ignore ST1006  MYSTYLE
//lint:file-ignore U1000  MYSTYLE

type Vector[T any] []T

func NewVector[T any]() Vector[T] {
	var array [10]T
	ref := array[0:0]
	return ref
}

func (self *Vector[T]) PushBack(value T) {
	*self = append([]T(*self), value)
}

func (self *Vector[T]) Delete(index int) (T, error) {
	if index < 0 || index >= len([]T(*self)) {
		var _n_ T
		return _n_, errors.New("index out of bound")
	}
	slice := []T(*self)
	boundary := len([]T(*self)) - 1
	elem := slice[index]
	for idx := index; idx < boundary; idx++ {
		slice[idx] = slice[idx+1]
	}
	*self = slice[:boundary]
	return elem, nil
}

func (self *Vector[T]) Range() []T {
	return []T(*self)
}
