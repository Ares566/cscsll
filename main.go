package cscsll

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
	// first element of list
	Head *node
}

func Create(_storage interface{}) *CScsll {
	return &CScsll{
		Storage: _storage,
		Head:    &node{Next: nil, Index: 0},
	}
}

func (cs *CScsll) Add(val interface{}) {

	curNode := cs.Head
	for {

		if curNode.Next == nil {
			break
		}
		curNode = curNode.Next
	}

	newNode := &node{Next: nil, Index: 1}

	curNode.Next = newNode

}
