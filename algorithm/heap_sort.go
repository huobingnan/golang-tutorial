package test

func heapSortParent(rank int) int {
	if rank%2 == 0 {
		// 奇数，那么它是左孩子
		return (rank - 1) >> 1
	} else {
		// 偶数，那么它是右孩子
		return (rank >> 1) - 1
	}
}

func heapSortLeft(rank int) int {
	return (rank << 1) + 1
}

func heapSortRight(rank int) int {
	return (rank + 1) << 1
}

func HeapSort(source []int) {
	// 建堆

}
