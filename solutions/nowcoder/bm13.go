package nowcoder

// link: https://www.nowcoder.com/practice/3fed228444e740c8be66232ce8b87c2f
// ACCEPT

func Bm13IsPail(head *ListNode) bool {
	// write code here
	container := make([]int, 0, 10)
	for head != nil {
		container = append(container, head.Val)
		head = head.Next
	}
	low, high := 0, len(container)-1
	for low < high {
		if container[low] != container[high] {
			return false
		}
		low += 1
		high -= 1
	}
	return true
}
