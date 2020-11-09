package glist

import (
	"sync"
)

type Array struct {
	sync.RWMutex

	slice      []Element
	concurrent bool
}

func NewArray() *Array {
	return &Array{}
}

func NewArrayConcurrent() *Array {
	return &Array{concurrent: true}
}

// Add appends the specified element to the end of this list
func (array *Array) Add(e Element) {
	if array.concurrent {
		array.Lock()
	}
	defer func() {
		if array.concurrent {
			array.Unlock()
		}
	}()

	array.slice = append(array.slice, e)
}

// Appends all of the elements in the specified slice to the end of this list
func (array *Array) AddSlice(es []Element) {
	if array.concurrent {
		array.Lock()
	}
	defer func() {
		if array.concurrent {
			array.Unlock()
		}
	}()

	array.slice = append(array.slice, es...)
}

// Returns true if this list contains the specified element.
func (array *Array) Contains(e Element) bool {
	if array.concurrent {
		array.RLock()
	}
	defer func() {
		if array.concurrent {
			array.RUnlock()
		}
	}()

	for i := range array.slice {
		if array.slice[i] != nil && array.slice[i].Equals(e) {
			return true
		}
	}
	return false
}

// Returns the element at the specified position in this list.
func (array *Array) Get(index int) Element {
	if array.concurrent {
		array.RLock()
	}
	defer func() {
		if array.concurrent {
			array.RUnlock()
		}
	}()

	l := len(array.slice)
	if index < 0 || index >= l {
		return nil
	}
	return array.slice[index]
}

// Returns the index of the first occurrence of the specified element in this list
// Or -1 if this list does not contain the element.
func (array *Array) IndexOf(e Element) int {
	if array.concurrent {
		array.RLock()
	}
	defer func() {
		if array.concurrent {
			array.RUnlock()
		}
	}()

	for i := range array.slice {
		if array.slice[i] != nil && array.slice[i].Equals(e) {
			return i
		}
	}
	return -1
}

// Removes the element at the specified position in this list
func (array *Array) RemoveIndex(index int) {
	if array.concurrent {
		array.Lock()
	}
	defer func() {
		if array.concurrent {
			array.Unlock()
		}
	}()

	if index < 0 || index > len(array.slice) {
		return
	}
	array.slice = append(array.slice[:index], array.slice[index+1:]...)
}

// Removes the first occurrence of the specified element from this list, if it is present (optional operation).
func (array *Array) Remove(e Element) bool {
	if array.concurrent {
		array.Lock()
	}
	defer func() {
		if array.concurrent {
			array.Unlock()
		}
	}()

	newSlice := make([]Element, 0, len(array.slice))
	for i := range array.slice {
		if array.slice[i] != nil && array.slice[i].Equals(e) {
			continue
		}
		newSlice = append(newSlice, array.slice[i])
	}
	array.slice = newSlice
}

// Sorts this list according to the order induced by the specified Comparator.
func (array *Array) Sort() {
	if array.concurrent {
		array.Lock()
	}
	defer func() {
		if array.concurrent {
			array.Unlock()
		}
	}()
	// TODO
}

// Returns true if this list contains no elements.
func (array *Array) IsEmpty() bool {
	if array.concurrent {
		array.RLock()
	}
	defer func() {
		if array.concurrent {
			array.RUnlock()
		}
	}()
	if len(array.slice) == 0 {
		return true
	}
	return false
}

func (array *Array) Len() int {
	if array.concurrent {
		array.RLock()
	}
	defer func() {
		if array.concurrent {
			array.RUnlock()
		}
	}()

	return len(array.slice)
}

// Removes all of the elements from this list (optional operation).
func (array *Array) Clear() {
	if array.concurrent {
		array.Lock()
	}
	defer func() {
		if array.concurrent {
			array.Unlock()
		}
	}()

	array.slice = nil
}
