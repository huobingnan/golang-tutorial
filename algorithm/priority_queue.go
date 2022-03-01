package algorithm

import "errors"

const (
	PriorityQueueLevel = 5
)

type PriorityQueue struct {
	level        int // 层数
	length       int // 长度
	elementArray []interface{}
	comparator   func(one, other interface{}) int
}

// 获取优先队列中元素的左孩子的秩
func leftChildRank(rank int) int {
	return (rank << 1) + 1
}

// 获取优先队列中元素的右孩子的秩
func rightChildRank(rank int) int {
	return (rank + 1) << 1
}

// 获取优点队列中元素的父节点的秩
func parentRank(rank int) int {
	if rank == 0 {
		return -1
	}
	if rank%2 != 0 {
		// 奇数是左孩子
		return (rank - 1) >> 1
	} else {
		return (rank >> 1) - 1
	}
}

// NewPriorityQueue 新建一个优先队列
// param: comparator 队列中元素的比较器
func NewPriorityQueue(comparator func(interface{}, interface{}) int) *PriorityQueue {
	return &PriorityQueue{
		level:        0,
		length:       0,
		elementArray: nil,
		comparator:   comparator,
	}
}

// 检查优先队列的数组状态
func (p *PriorityQueue) checkElementArrayState() {
	if p.elementArray == nil {
		p.elementArray = make([]interface{}, (1<<PriorityQueueLevel)-1)
		p.level = PriorityQueueLevel
	}
}

// Level 获取优先队列的层数
func (p *PriorityQueue) Level() int {
	return p.level
}

// Length get the length information of priority queue current state
func (p *PriorityQueue) Length() int {
	return p.length
}

// Capacity get the capacity of low level array used by priority queue
func (p *PriorityQueue) Capacity() int {
	return cap(p.elementArray)
}

func (p *PriorityQueue) ensureArrayCapacity() {
	if p.length >= cap(p.elementArray) {
		p.level += 1
		newSpace := make([]interface{}, (1<<p.level)-1)
		copy(newSpace, p.elementArray)
		p.elementArray = newSpace
	}
}

func (p *PriorityQueue) percolateUp(rank int) {
	if rank == 0 || rank == -1 {
		return
	}
	parentRank := parentRank(rank)
	if parentRank == -1 {
		return
	}
	parent := p.elementArray[parentRank]
	current := p.elementArray[rank]
	if p.comparator(current, parent) > 0 {
		// swap
		p.elementArray[parentRank] = current
		p.elementArray[rank] = parent
		// recursive call
		p.percolateUp(parentRank)
	}
}

func (p *PriorityQueue) percolateDown(rank int) {
	leftRank := leftChildRank(rank)
	rightRank := rightChildRank(rank)
	// select maximum child from left & right
	var max interface{}
	var maxRank int
	if leftRank < p.length && rightRank >= p.length {
		max = p.elementArray[leftRank]
		maxRank = leftRank
	} else if rightRank < p.length && leftRank >= p.length {
		max = p.elementArray[rightRank]
		maxRank = rightRank
	} else if rightRank < p.length && leftRank < p.length {
		// select max
		left := p.elementArray[leftRank]
		right := p.elementArray[rightRank]
		if p.comparator(left, right) > 0 {
			max = left
			maxRank = leftRank
		} else {
			max = right
			maxRank = rightRank
		}
	} else {
		return
	}

	current := p.elementArray[rank]
	// compare current node value and maximum node value
	if p.comparator(current, max) < 0 {
		// swap
		p.elementArray[maxRank] = current
		p.elementArray[rank] = max
		// recursive call
		p.percolateDown(maxRank)
	}

}

// Push push an element into priority queue
func (p *PriorityQueue) Push(element interface{}) {
	p.checkElementArrayState()
	p.ensureArrayCapacity()
	p.elementArray[p.length] = element
	p.percolateUp(p.length)
	p.length += 1
}

// Pop pop an element from priority queue
func (p *PriorityQueue) Pop() (interface{}, error) {
	p.checkElementArrayState()
	if p.length == 0 {
		return nil, errors.New("underflow")
	}
	ret := p.elementArray[0]
	p.elementArray[0] = p.elementArray[p.length-1]
	p.length -= 1
	p.percolateDown(0)
	return ret, nil
}

// Top get the maximum/minimum element of priority queue
func (p *PriorityQueue) Top() interface{} {
	p.checkElementArrayState()
	return p.elementArray[0]
}

// IsEmpty check if the priority queue is empty
func (p *PriorityQueue) IsEmpty() bool {
	return p.length == 0
}

func (p *PriorityQueue) IsNotEmpty() bool {
	return !p.IsEmpty()
}
