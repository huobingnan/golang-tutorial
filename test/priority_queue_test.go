package test

import (
	"fmt"
	"math/rand"
	"testing"
)
import "../algorithm"

type Comparable int

func (c Comparable) Compare(other interface{}) int {
	return int(c) - other.(int)
}

func IntegerComparator(i1, i2 interface{}) int {
	return i1.(int) - i2.(int)
}

func GenerateRandomIntArray() []int {
	const length int = 100000
	ret := make([]int, length)
	for i := 0; i < length; i++ {
		ret[i] = rand.Intn(100)
	}
	return ret
}

func TestComparable(t *testing.T) {
	c := Comparable(123)
	t.Log(c.Compare(123))
	t.Log(c.Compare(12))
	t.Log(c.Compare(1234))
}

func TestNewPriorityQueue(t *testing.T) {
	queue := algorithm.NewPriorityQueue(IntegerComparator)
	t.Log("length = ", queue.Length())
	t.Log("capacity = ", queue.Capacity())
	t.Log("level = ", queue.Level())
}

func TestPush(t *testing.T) {
	q := algorithm.NewPriorityQueue(IntegerComparator)
	array := []int{2, 5, 1, 9, 7, 0}
	for _, each := range array {
		q.Push(each)
	}
	t.Log("top = ", q.Top())
}

func TestPop(t *testing.T) {
	q := algorithm.NewPriorityQueue(IntegerComparator)
	array := GenerateRandomIntArray()
	for _, each := range array {
		q.Push(each)
	}
	fmt.Printf("generated array:\n\t%v\n", array)
	fmt.Printf("sorted array:\n\t")
	fmt.Printf("[ ")
	for q.IsNotEmpty() {
		m, err := q.Pop()
		if err != nil {
			break
		}
		fmt.Printf("%d ", m)
	}
	fmt.Printf("]\n")
}
