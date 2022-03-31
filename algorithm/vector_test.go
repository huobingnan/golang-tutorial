package test

import (
	"testing"
)

func TestVectorPushBack(t *testing.T) {
	t.Run("case 1", func(t *testing.T) {
		vec := NewVector[int]()
		for i := 1; i <= 10; i++ {
			vec.PushBack(i)
		}
	})
}

func TestVectorAccess(t *testing.T) {
	t.Run("range access", func(t *testing.T) {
		vec := NewVector[int]()
		for i := 0; i < 10; i++ {
			vec.PushBack(i)
		}
		for _, each := range vec {
			t.Log(each)
		}
	})
	t.Run("index acess", func(t *testing.T) {
		vec := NewVector[int]()
		for i := 0; i < 10; i++ {
			vec.PushBack(i)
		}
		if vec[0] != 0 {
			t.Fatal("vec[0] expect 0, actual ", vec[0])
		}
		if vec[len([]int(vec))-1] != 9 {
			t.Fatal("vec[len(vec)-1] expect 9, actual ", vec[len([]int(vec))-1])
		}
		if vec[5] != 5 {
			t.Fatal("vec[5] expect 5, actual ", vec[5])
		}
	})
}
