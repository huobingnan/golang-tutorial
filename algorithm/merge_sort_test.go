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
}

func BenchmarkMergeSort(b *testing.B) {
	rand.Seed(time.Now().Unix())
	src := make([]int, b.N)
	for idx := 0; idx < b.N; idx++ {
		src[idx] = rand.Intn(b.N)
	}
	MergeSort(src)
}
