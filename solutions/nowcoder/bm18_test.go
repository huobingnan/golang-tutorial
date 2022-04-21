package nowcoder

import "testing"

func TestBm18BinarySearch(t *testing.T) {
	t.Run("Case1", func(t *testing.T) {
		array := []int{1, 2, 3, 4, 5, 6}
		t.Log(_bm18BinarySearch(array, 7))
		t.Log(_bm18BinarySearch(array, 0))
	})
}

func TestFind(t *testing.T) {
	m := [][]int{
		{1, 2, 8, 9},
		{2, 4, 9, 12},
		{4, 7, 10, 13},
		{6, 8, 11, 15},
	}
	t.Run("Case1", func(t *testing.T) {
		r := Find(4, m)
		t.Log(r)
	})
}
