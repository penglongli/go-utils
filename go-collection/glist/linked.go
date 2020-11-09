package glist

import "sync"

type Linked struct {
	sync.RWMutex

	length int

	head *node
	tail *node
}

type node struct {
}

func NewLinkedConcurrent() *Linked {

}
