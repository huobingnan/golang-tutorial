package main

import (
	"container/list"
	"fmt"
)

func main() {
	func() {
		// list可以当做普通的链表使用
		l := list.New()
		for i := 0; i < 3; i++ {
			l.PushBack(i)
		}
		// 本质上list是一个双向循环链表
		// list不能使用range来遍历

		// l.Front()是获取链表的首节点，用于正向迭代
		for e := l.Front(); e != nil; e = e.Next() {
			fmt.Printf("itr : %d\n", e.Value)
		}
		// l.Back()是获取链表的尾节点，用于反向迭代
		for e := l.Back(); e != nil; e = e.Prev() {
			fmt.Printf("reverse itr : %d\n", e.Value)
		}
	}()

	func() {
		// 对于栈和队列这类受限使用的线性表，list提供了相应的API
		// 下面这段代码用于模拟栈的行为
		s := list.New()
		// push
		for i := 0; i < 5; i++ {
			s.PushBack(i)
		}
		// pop
		for s.Len() != 0 {
			top := s.Remove(s.Back())
			fmt.Printf("pop : %d\n", top)
		}
	}()

	func() {
		// 使用list模拟队列行为
		q := list.New()
		// enqueue
		for i := 0; i < 5; i++ {
			q.PushBack(i)
		}
		// dequeue
		for q.Len() != 0 {
			first := q.Remove(q.Front())
			fmt.Printf("dequeue : %d\n", first)
		}
	}()
}
