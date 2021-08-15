package cscsll

import (
	"fmt"
)

// Single element of list
type node struct {
	Next *node

	// index of stored in Storage value
	Index int
}

// CScsll represents a Circular singly linked list
type CScsll struct {

	// array containing all uniq elements
	Storage interface{}
	// Pointer first element of list
	Head *node
}

func Create(_storage interface{}) *CScsll {
	return &CScsll{
		Storage: _storage,
		Head:    nil,
	}
}

func (cs *CScsll) Add(val interface{}) {

	// calculate where to insert new value
	_index := cs.findIndex[int](cs.Storage, val)

	newNode := &node{Next: nil, Index: _index}

	switch cs.Storage.(type) {
	case []string:
		cs.Storage.([]string)[_index] = val.(string)
	case []int:
		cs.Storage.([]int)[_index] = val.(int)
	}

	curNode := cs.Head

	// list is empty
	if curNode == nil {
		cs.Head = newNode
		return
	}

	for {
		if curNode.Next == nil {
			break
		}
		curNode = curNode.Next
	}

	curNode.Next = newNode

}

func (cs *CScsll) Dump() {

	curNode := cs.Head
	for {

		fmt.Printf("%#v \n", curNode)

		if curNode == nil || curNode.Next == nil {
			break
		}
		curNode = curNode.Next
	}
}

// Returns index i at which x == a[i],
// or len(a) if there is no such index.
func (cs *CScsll) findIndex[T any](a []T, x T) int {

	_storage := cs.Storage.([]T)
	for i, n := range _storage {
		if x == n {
			return i
		}
	}
	return len(a)
}
