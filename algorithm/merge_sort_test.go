package algorithm

import (
	"math/rand"
	"testing"
	"time"
)

func TestMergeSort(t *testing.T) {
	rand.Seed(time.Now().Unix())
	size := rand.Intn(15)
	src := make([]int, size)
	t.Run("V1", func(t *testing.T) {
		for idx := 0; idx < size; idx++ {
			src[idx] = rand.Intn(100)
		}
		t.Log("---Before sort")
		t.Log(src)
		MergeSort(src)
		t.Log("---After sort")
		t.Log(src)
		// check
		for idx := 1; idx < len(src); idx++ {
			if src[idx] < src[idx-1] {
				t.Fail()
			}
		}
	})
	t.Run("V2", func(t *testing.T) {
		for idx := 0; idx < size; idx++ {
			src[idx] = rand.Intn(100)
		}
		t.Log("---Before sort")
		t.Log(src)
		MergeSortV2(src)
		t.Log("---After sort")
		t.Log(src)
		// check
		for idx := 1; idx < len(src); idx++ {
			if src[idx] < src[idx-1] {
				t.Fail()
			}
		}
	})

}

func BenchmarkMergeSort(b *testing.B) {
	rand.Seed(time.Now().Unix())
	src := make([]int, b.N)
	for idx := 0; idx < b.N; idx++ {
		src[idx] = rand.Intn(b.N)
	}
	t1 := make([]int, len(src))
	copy(t1, src)
	b.Run("V1", func(b *testing.B) {
		MergeSort(t1)
	})
	t2 := make([]int, len(src))
	copy(t2, src)
	b.Run("V2", func(b *testing.B) {
		MergeSortV2(t2)
	})
}
