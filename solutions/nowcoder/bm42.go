package nowcoder

// link: https://www.nowcoder.com/practice/54275ddae22f475981afa2244dd448c6
// ACCEPT

var stack1 []int
var stack2 []int
var top2 int

func Push(node int) {
	if stack1 == nil {
		stack1 = make([]int, 0, 10)
	}
	stack1 = append(stack1, node)
}

func Pop() int {
	if stack2 == nil {
		stack2 = make([]int, len(stack1))
	}
	if top2 < 0 {
		ptr := len(stack1) - 1
		idx := 0
		for ; idx < len(stack2); idx++ {
			if ptr < 0 {
				break
			}
			stack2[idx] = stack1[ptr]
			ptr -= 1
		}
		top2 = idx
	}
	ret := stack2[top2]
	top2 -= 1
	return ret
}
