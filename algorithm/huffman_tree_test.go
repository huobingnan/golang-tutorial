package algorithm

import (
	"fmt"
	"testing"
)

func TestHuffmanTreeSetup(t *testing.T) {
	t.Run("Case 1", func(t *testing.T) {
		codes := []rune{rune('A'), rune('B'), rune('C'), rune('D'), rune('E')}
		weights := []int{20, 5, 40, 10, 25}
		table := GetHuffmanEncodeTable(codes, weights)
		for char, code := range table {
			fmt.Printf("%c %b\n", char, code)
		}
	})

	t.Run("Case 2", func(t *testing.T) {
		codes := []rune{'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h'}
		weights := []int{5, 29, 7, 8, 14, 23, 3, 11}
		table := GetHuffmanEncodeTable(codes, weights)
		for char, code := range table {
			fmt.Printf("%c %b\n", char, code)
		}
	})
}
