package test

type HuffmanNode struct {
	Left, Right *HuffmanNode
	Code        rune
	Weight      int
}

func huffmanEncodeMap(root *HuffmanNode, result map[rune]uint64, code uint64) {
	if root == nil {
		return
	}
	if root.Code != -1 {
		result[root.Code] = code
	}
	huffmanEncodeMap(root.Left, result, code<<1)
	huffmanEncodeMap(root.Right, result, (code<<1)+1)
}

func setupHuffmanTree(codeStream []rune, weightStream []int) map[rune]uint64 {
	if len(codeStream) != len(weightStream) {
		panic("invalid params")
	}
	nodes := make([]*HuffmanNode, len(codeStream))
	for idx, code := range codeStream {
		nodes[idx] = &HuffmanNode{
			Code:   code,
			Left:   nil,
			Right:  nil,
			Weight: weightStream[idx],
		}
	}
	pq := NewPriorityQueue(func(i interface{}, i2 interface{}) int {
		return i2.(*HuffmanNode).Weight - i.(*HuffmanNode).Weight
	})
	for _, node := range nodes {
		pq.Push(node)
	}
	for pq.Length() != 1 {
		n1, err := pq.Pop()
		if err != nil {
			panic(err)
		}
		n2, err := pq.Pop()
		if err != nil {
			panic(err)
		}
		n := new(HuffmanNode)
		n.Weight = n1.(*HuffmanNode).Weight + n2.(*HuffmanNode).Weight
		n.Left = n1.(*HuffmanNode)
		n.Right = n2.(*HuffmanNode)
		n.Code = -1
		pq.Push(n)
	}
	root, err := pq.Pop()
	if err != nil {
		panic(err)
	}
	result := make(map[rune]uint64)
	huffmanEncodeMap(root.(*HuffmanNode), result, 0)
	return result
}

func GetHuffmanEncodeTable(codes []rune, weight []int) map[rune]uint64 {
	return setupHuffmanTree(codes, weight)
}
