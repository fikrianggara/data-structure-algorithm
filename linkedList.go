package main

import "fmt"

type sllNode struct {
	item int
	next *sllNode
}

type singlyLinkedList struct {
	head *sllNode
	len  int
}

func (l *singlyLinkedList) getAt(pos int) *sllNode {
	ptr := l.head
	if pos < 0 {
		return ptr
	}
	if pos > (l.len - 1) {
		return nil
	}
	for i := 0; i < pos; i++ {
		ptr = ptr.next
	}
	return ptr
}

func (l *singlyLinkedList) traverse() {
	ptr := l.head
	for i := 0; i < l.len; i++ {
		fmt.Println("Node :", ptr)
		ptr = ptr.next
	}
}

func (l *singlyLinkedList) insert(item int) {
	node := sllNode{item: item}

	if l.len == 0 {
		l.head = &node
		l.len++
		return
	}
	//mulai dari paling depan, yaitu pointer head
	pointer := l.head

	//insert ke node paling belakang
	for i := 0; i < l.len; i++ {

		//node paling belakang punya pointer next == nil,
		//kita masukin node baru di situ
		if pointer.next == nil {
			pointer.next = &node
			l.len++
			return
		}
		//traverse linkedListnya
		pointer = pointer.next
	}
}

func (l *singlyLinkedList) insert_at(pos int, item int) {

	node := sllNode{item: item}

	if pos < 0 {
		return
	}
	if pos == 0 && l.len != 0 {
		node.next = l.head
		l.head = &node
		l.len++
		return
	}
	if pos > l.len {
		return
	}

	n := l.getAt(pos)
	node.next = n
	prevNode := l.getAt(pos - 1)
	prevNode.next = &node
	l.len++
}

func (l *singlyLinkedList) delete_at(pos int) {
	if pos < 0 {
		return
	}
	if l.len == 0 {
		return
	}
	if pos > l.len {
		return
	}
	prevNode := l.getAt(pos - 1)
	if prevNode == nil {
		return
	}
	prevNode.next = l.getAt(pos).next
	l.len--
}

func (l *singlyLinkedList) reverse() {
	if l.len <= 0 {
		return
	}
	if l.len == 1 {
		return
	}

	tempPtr := l.head
	var prevNode *sllNode

	for tempPtr.next != nil {

		tempNode := tempPtr
		tempPtr = tempPtr.next
		tempNode.next = prevNode
		l.head = tempPtr
		prevNode = tempNode
		fmt.Println("tempNode :", tempNode, ", tempPtr :", tempPtr, ", prev :", prevNode)
	}
	tempPtr.next = prevNode

}
