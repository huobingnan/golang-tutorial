package algorithm

import (
	"testing"
)

func TestDirectInsertSort(t *testing.T) {
	t.Run("Case 1", func(t *testing.T) {
		source := []int{3, 1, 5, 6, 2, 4, 0}
		DirectInsertSort(source)
		t.Log(source)
	})
}

func TestBinaryInsertSort(t *testing.T) {
	t.Run("Case 1", func(t *testing.T) {
		source := []int{3, 1, 5, 6, 2, 4, 0}
		BinaryInsertSort(source)
		t.Log(source)
	})
}
