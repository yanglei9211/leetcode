package main

import (
	"container/list"
	"errors"
	"fmt"
)

type Queue struct {
	list *list.List
}

func NewQueue() *Queue {
	ls := list.New()
	return &Queue{ls}
}

func (q *Queue) Push(value interface{}) {
	q.list.PushBack(value)
}

func (q *Queue) Pop() (interface{}, error) {
	if q.Empty() {
		return nil, errors.New("queue empty")
	} else {
		e := q.list.Front()
		if e != nil {
			q.list.Remove(e)
			return e.Value, nil
		}
		return nil, errors.New("inter error")
	}
}

func (q *Queue) Front() (interface{}, error) {
	if q.Empty() {
		return nil, errors.New("queue empty")
	} else {
		e := q.list.Front()
		if e != nil {
			return e.Value, nil
		} else {
			return nil, errors.New("inter error")
		}
	}
}

func (q *Queue) Len() int {
	return q.list.Len()
}

func (q *Queue) Empty() bool {
	return q.list.Len() == 0
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type QueueItem struct {
	node  *TreeNode
	level int
}

func levelOrder(root *TreeNode) [][]int {
	q := NewQueue()
	ans := [][]int{}
	if root == nil {
		return ans
	}
	q.Push(QueueItem{node: root, level: 0})
	for !q.Empty() {
		res, _ := q.Pop()
		cur := res.(QueueItem)
		if cur.level >= len(ans) {
			ans = append(ans, []int{})
		}
		ans[cur.level] = append(ans[cur.level], cur.node.Val)
		if cur.node.Left != nil {
			q.Push(QueueItem{node: cur.node.Left, level: cur.level + 1})
		}
		if cur.node.Right != nil {
			q.Push(QueueItem{node: cur.node.Right, level: cur.level + 1})
		}
	}
	return ans
}

func main() {
	t1 := TreeNode{Val: 1}
	t21 := TreeNode{Val: 2}
	t22 := TreeNode{Val: 2}
	t31 := TreeNode{Val: 3}
	t32 := TreeNode{Val: 3}
	//t41 := TreeNode{Val:4}
	//t42 := TreeNode{Val:4}
	//t1.Left = &t21
	//t21.Left = &t31
	//t21.Right = &t41
	//t1.Right = &t22
	//t22.Left = &t42
	//t22.Right = &t32
	t1.Left = &t21
	t21.Left = &t31
	t1.Right = &t32
	t32.Left = &t22
	fmt.Println(levelOrder(&t1))
}
