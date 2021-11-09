package engine

import (
	"github.com/wenchangshou2/vd-node-manage/module/agent/dto"
	"sync"
)

// DoubleObject 链表元素对象烦死
type DoubleObject dto.TaskItem

// DoubleNode 单个节点
type DoubleNode struct {
	Data DoubleObject
	Prev *DoubleNode
	Next *DoubleNode
}

// DoubleList 双向链表
type DoubleList struct {
	mutex *sync.RWMutex
	Size  uint
	Head  *DoubleNode
	Tail  *DoubleNode
}

func NewDoubleList() *DoubleList {
	return &DoubleList{
		mutex: new(sync.RWMutex),
		Size:  0,
		Head:  nil,
		Tail:  nil,
	}
}

// Append 追加一个新的元素
func (list *DoubleList) Append(node *DoubleNode) bool {
	if node == nil {
		return false
	}
	list.mutex.Lock()
	defer list.mutex.Unlock()
	//如果是空表
	if list.Size == 0 {
		list.Head = node
		list.Tail = node
		node.Next = nil
		node.Prev = nil
	} else {
		node.Prev = list.Tail
		node.Next = nil
		list.Tail.Next = node
		list.Tail = node
	}
	list.Size++
	return true
}
func (list *DoubleList) Insert(index uint, node *DoubleNode) bool {
	if index > list.Size || node == nil {
		return false
	}
	if index == list.Size {
		return list.Append(node)
	}
	list.mutex.Lock()
	defer list.mutex.Lock()
	if index == 0 {
		node.Next = list.Head
		list.Head = node
		list.Head.Prev = nil
		list.Size++
		return true
	}
	nextNode := list.Get(index)
	node.Prev = nextNode.Prev
	node.Next = nextNode
	nextNode.Prev.Next = node
	nextNode.Prev = node
	list.Size++
	return true
}

// Delete 删除指定的元素
func (list *DoubleList) Delete(index uint) bool {
	if index > list.Size-1 {
		return false
	}
	list.mutex.Lock()
	defer list.mutex.Unlock()
	if index == 0 {
		if list.Size == 1 {
			list.Head = nil
			list.Tail = nil
		} else {
			list.Head.Next.Prev = nil
			list.Head = list.Head.Next
		}
		list.Size--
		return true
	}
	if index == list.Size-1 {
		list.Tail.Prev.Next = nil
		list.Tail = list.Tail.Prev
		list.Size--
		return true
	}
	node := list.Get(index)
	node.Prev.Next = node.Next
	node.Next.Prev = node.Prev
	list.Size--
	return true
}

func (list *DoubleList) Get(index uint) *DoubleNode {
	if list.Size == 0 || index > list.Size-1 {
		return nil
	}
	if index == 0 {
		return list.Head
	}
	node := list.Head
	var i uint
	for i = 1; i < index; i++ {
		node = node.Next
	}
	return node
}

type TaskLinkedList struct {
	DoubleList
}

func NewTaskLinedList() *TaskLinkedList {
	return &TaskLinkedList{struct {
		mutex *sync.RWMutex
		Size  uint
		Head  *DoubleNode
		Tail  *DoubleNode
	}{mutex: new(sync.RWMutex), Size: 0, Head: nil, Tail: nil}}
}
func (list *TaskLinkedList) GetByTaskId(id string) (uint, bool) {
	if list.Size == 0 {
		return 0, false
	}
	node := list.Head
	if id == node.Data.ID {
		return 0, true
	}
	var i uint = 0
	for node.Next != nil {
		i++
		node = node.Next
		if node.Data.ID == id {
			return i, true
		}
	}
	return 0, false

}

func (list *TaskLinkedList) Foreach(callback func(node *DoubleNode) bool) {
	if list == nil || list.Size == 0 {
		return
	}
	list.mutex.RLock()
	defer list.mutex.RUnlock()
	ptr := list.Head
	for ptr != nil {
		if !callback(ptr) {
			return
		}
		ptr = ptr.Next
	}
}
