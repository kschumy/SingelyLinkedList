package singleylinkedlist

// Accessing field 'next' of 'toDelete' (type '**node') requires explicit dereference.
// Inspection info: Reports method calls and field accesses on expressions with a double pointer type

import (
	"fmt"
)

type List struct {
	head *node

}

type node struct {
	data int
	next *node
}

// method to add a new node with the specific data value in the linked list
// insert the new node at the beginning of the linked list
func (list *List) Insert(value int) {
	list.head = &node{data: value, next: list.head}
}


// method to find if the linked list contains a node with specified value
// returns true if found, false otherwise
func (list *List) Search(value int) bool {
	for curr := list.head; curr != nil; curr = curr.next {
		if curr.data == value {
			return true
		}
	}
	return false
}

// method to return the max value in the linked list
// returns the data value and not the node
func (list *List) FindMax() (int, error) {
	if list.isEmpty() {
		return 0, emptyListError()
	}

	return list.findMax(), nil
}

func (list *List) findMax() int {
	max := list.head.data
	for curr := list.head.next; curr != nil; curr = curr.next {
		if curr.data > max {
			max = curr.data
		}
	}
	return max
}


// method to return the min value in the linked list
// returns the data value and not the node
func (list *List) FindMin() (int, error) {
	if list.isEmpty() {
		return 0, emptyListError()
	}
	min := list.head.data
	for curr := list.head.next; curr != nil; curr = curr.next {
		if curr.data < min {
			min = curr.data
		}
	}
	return min, nil
}

//func getData(curr *node) int {
//	return curr.data
//}


// method that returns the length of the singly linked list
func (list *List) Length() int {
	length := 0
	for curr := list.head; curr != nil; curr = curr.next {
		length++
	}
	return length
}

// method to return the value of the nth element from the beginning
// assume indexing starts at 0 while counting to n
func (list *List) FindNthFromBeginning(n int) (int, error) {
	nthNode, err := list.getNthFromBeginning(n)

	if nthNode == nil {
		return 0, err
	}

	return nthNode.data, nil
}

// method to return the value of the nth element from the beginning
// assume indexing starts at 0 while counting to n
func (list *List) getNthFromBeginning(n int) (*node, error) {
	if n < 0 {
		return nil, invalidIndexList(n) // QUESTION: is x right?
	}

	if list.isEmpty() {
		return nil, emptyListError()
	}

	curr := list.head
	for i := 0; i < n; i++ {
		if curr == nil {
			return nil, fmt.Errorf("singleylinkedlist: not index %x", n) // QUESTION: is x right?
		}
		curr = curr.next
	}

	return curr, nil
}


// method to insert a new node with specific data value, assuming the linked
// list is sorted in ascending order
func (list *List) InsertAscending(value int) {
	if list.isEmpty() {
		list.Insert(value)
	} else {
		//for curr := list.head; ; curr = curr.next {
		//	if  curr.next == nil || curr.next.data > value {
		//		curr.next = &node{data: value, next: curr.next}
		//		return
		//	}
		//}
		curr := list.head
		for ; curr.next != nil && curr.next.data < value; curr = curr.next { }
		curr.next = &node{data: value, next: curr.next}
	}
}

// method to print all the values in the linked list
// NOTE: does nothing if list is empty
func (list *List) Visit() {
	if !list.isEmpty() {
		list.printList()
	}
	fmt.Println()
 }

// Prints list. List cannot be empty
func (list *List) printList() {
	fmt.Print(list.head.data)
	for curr := list.head.next; curr != nil; curr = curr.next {
		fmt.Print(", ", curr.data)
	}
}

func (list *List) isEmpty() bool {
	return list.head == nil
}

// method to delete the first node found with specified value
// NOTE: does nothing if list is empty or if value is not found
func (list *List) Delete(value int) error {
	if list.isEmpty() { return emptyListError() }


	if !list.removeNode(value) {
		return fmt.Errorf("singleylinkedlist: value %x not found to delete", value)
	}
	return nil
}


func (list *List) removeNode(value int) bool {
	var toRemove *node

	if list.head.data == value {
		toRemove = list.head
		list.head = list.head.next
	} else {
		nodeBefore :=  list.getNodeBefore(value)
		if nodeBefore == nil { // nodeBefore is nil if value is not in list
			return false // indicates delete was not successful
		}
		toRemove = nodeBefore.next
		nodeBefore.next = toRemove.next
	}
	toRemove.next = nil
	toRemove = nil
	return true // indicates delete was successful
}

// list must not be empty
func (list *List) getNodeBefore(value int) *node {
	for curr := list.head; curr.next != nil; curr = curr.next {
		if curr.next.data == value {
			return curr
		}
	}
	return nil
}


// method to reverse the singly linked list
// note: the nodes should be moved and not just the values in the nodes
func (list *List) Reverse() {
	if list.isEmpty() { return }

	curr := list.head
	nextNode := curr.next
	curr.next = nil
	for nextNode != nil {
		temp := nextNode.next
		nextNode.next = curr
		if temp == nil {
			nextNode.next = curr
			list.head = nextNode
			return
		}
		curr = nextNode
		nextNode = temp
	}
	fmt.Println("hello")
}

// Advanced Exercises
// returns the value at the middle element in the singly linked list
func (list *List) FindMiddleValue() (int, error) {
	if list.isEmpty() { return 0, emptyListError() }

	toReturn := list.head

	for curr := list.head; curr.next != nil && curr.next.next != nil; curr = curr.next.next {
		toReturn = toReturn.next
	}

	return toReturn.data, nil

}

// find the nth node from the end and return its value
// assume indexing starts at 0 while counting to n
func (list *List) FindNthFromEnd(n int) (int, error){
	frontNode, err := list.getNthFromBeginning(n)
	if err != nil {
		return 0, err
	}
	return list.iterateInParallel(frontNode).data, nil
}


// frontNode cannot be nil and list cannot be empty,
// otherwise does whatever nullpointerexception is in Go
func (list *List) iterateInParallel(frontNode *node) *node {
	toReturn := list.head
	for ; frontNode.next != nil; frontNode = frontNode.next {
		toReturn = toReturn.next
	}
	return toReturn
}

func emptyListError() error {
	return fmt.Errorf("singleylinkedlist: empty list")
}

func invalidIndexList(n int) error {
	return fmt.Errorf("singleylinkedlist: negative index %x", n)
}

// checks if the linked list has a cycle. A cycle exists if any node in the
// linked list links to a node already visited.
// returns true if a cycle is found, false otherwise.
func (list *List) HasCycle() bool {
	if list.isEmpty() { return false }
	fast, slow := list.head.next, list.head
	for ; fast.next != nil && fast.next.next != nil; fast = fast.next.next {
		fmt.Println(fast.next.data, ", ", slow.data)
		if fast == slow || fast.next == slow {
			return true
		}
		slow = slow.next
	}
	return false
}


// Creates a cycle in the linked list for testing purposes
// Assumes the linked list has at least one node
func (list *List) CreateCycle() {
	if list.head == nil { // don't do anything if the linked list is empty
		return
	}

	// navigate to last node
	current := list.head
	for current.next != nil {
		current = current.next
	}
	current.next = list.head // make the last node link to first node
}



