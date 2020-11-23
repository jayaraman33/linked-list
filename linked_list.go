package linkedlist

import "errors"

type Node struct {
	Val        interface{}
	next, prev *Node
}

func (n *Node) Next() *Node {
	return n.next
}

func (n *Node) Prev() *Node {
	return n.prev
}

type List struct {
	head *Node
	tail *Node
}

var ErrEmptyList error = errors.New("list is empty")

func NewList(args ...interface{}) *List {
	l := &List{}
	for _, arg := range args {
		l.PushBack(arg)
	}
	return l
}

func (l *List) PushFront(v interface{}) {
	node := &Node{
		Val:  v,
		next: l.head,
	}
	if l.head == nil {
		l.tail = node
	} else {
		l.head.prev = node
	}
	l.head = node
}

func (l *List) PushBack(v interface{}) {
	node := &Node{
		Val:  v,
		prev: l.tail,
	}
	if l.tail == nil {
		l.head = node
	} else {
		l.tail.next = node
	}
	l.tail = node
}

func (l *List) PopFront() (interface{}, error) {
	if l.isEmpty() {
		return nil, ErrEmptyList
	}
	val := l.head.Val
	if l.head.next == nil {
		l.head = nil
		l.tail = nil
	} else {
		l.head = l.head.next
		l.head.prev.next = nil
		l.head.prev = nil
	}
	return val, nil
}

func (l *List) PopBack() (interface{}, error) {
	if l.isEmpty() {
		return nil, ErrEmptyList
	}
	val := l.tail.Val
	if l.tail.prev == nil {
		l.tail = nil
		l.head = nil
	} else {
		l.tail = l.tail.prev
		l.tail.next.prev = nil
		l.tail.next = nil
	}
	return val, nil
}

func (l *List) Reverse() *List {
	if l.isEmpty() {
		return l
	}
	l.head, l.tail = l.tail, l.head
	node := l.head
	for node != nil {
		node.prev, node.next = node.next, node.prev
		node = node.next
	}
	return l
}

func (l *List) First() *Node {
	return l.head
}

func (l *List) Last() *Node {
	return l.tail
}

func (l *List) isEmpty() bool {
	return l.head == nil && l.tail == nil
}
