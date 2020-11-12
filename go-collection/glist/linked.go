package glist

import "sync"

type Linked struct {
	sync.RWMutex

	length int

	head *node
	tail *node
}

type node struct {
	E    Element
	prev *node
	next *node
}

func NewLinkedConcurrent() *Linked {
	head := &node{}
	tail := &node{}

	head.next = tail
	tail.prev = head

	return &Linked{
		length: 0,
		head:   head,
		tail:   tail,
	}
}

func (linked *Linked) Add(e Element) {
	linked.Lock()
	defer func() {
		linked.length++
		linked.Unlock()
	}()

	node := &node{
		E:    e,
		next: linked.tail,
	}

	prevNode := linked.tail.prev
	prevNode.next = node
	linked.tail.prev = node
	node.prev = prevNode
}

func (linked *Linked) AddSlice(es []Element) {
	linked.Lock()
	defer func() {
		linked.length += len(es)
		linked.Unlock()
	}()

	esLen := len(es)
	slice := make([]*node, esLen, esLen)
	for i := range es {
		slice[i] = &node{
			E: es[i],
		}
	}

	prevNode := linked.tail.prev
	for i := range slice {
		if i == 0 {
			prevNode.next = slice[i]
		}
		if i != 0 {
			slice[i].prev = slice[i-1]
		}
		if i != esLen-1 {
			slice[i].next = slice[i+1]
		}
		if i == esLen-1 {
			linked.tail.prev = slice[i]
		}
	}
}

func (linked *Linked) Contains(e Element) bool {
	linked.RLock()
	defer func() {
		linked.RUnlock()
	}()

	currentNode := linked.head.next
	for ; currentNode.next != linked.tail; currentNode = currentNode.next {
		if currentNode.E.Equals(e) {
			return true
		}
	}
	return false
}

func (linked *Linked) Get(index int) Element {
	linked.RLock()
	defer func() {
		linked.RUnlock()
	}()

	if index < 0 || index > linked.length-1 {
		return nil
	}
	if index < linked.length/2 {
		currentNode := linked.head
		for ans := 0; ans <= index; ans++ {
			currentNode = currentNode.next
		}
		return currentNode.E
	}
	currentNode := linked.tail
	for ans := linked.length - 1; ans >= index; ans-- {
		currentNode = currentNode.prev
	}
	return currentNode.E
}

func (linked *Linked) IndexOf(e Element) int {
	linked.RLock()
	defer func() {
		linked.RUnlock()
	}()

	ans := 0
	currentNode := linked.head.next
	for ; currentNode.next != linked.tail; currentNode = currentNode.next {
		if currentNode.E.Equals(e) {
			return ans
		}
		ans++
	}
	return -1
}

func (linked *Linked) RemoveIndex(index int) {

}

func (linked *Linked) Remove(e Element) {

}

func (linked *Linked) Sort() {

}

func (linked *Linked) IsEmpty() bool {
	linked.RLock()
	defer func() {
		linked.RUnlock()
	}()

	if linked.length == 0 {
		return true
	}
	return false
}

func (linked *Linked) Len() int {
	linked.RLock()
	defer func() {
		linked.RUnlock()
	}()

	return linked.length
}

func (linked *Linked) Clear() {
	linked.Lock()
	defer func() {
		linked.Unlock()
	}()

}
