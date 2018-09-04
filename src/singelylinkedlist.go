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
	if n < 0 {
		return 0, fmt.Errorf("singleylinkedlist: negative index %x", n) // QUESTION: is x right?
	}

	if list.isEmpty() {
		return 0, emptyListError()
	}

	curr := list.head
	for i := 0; i < n; i++ {
		if curr == nil {
			return 0, fmt.Errorf("singleylinkedlist: not index %x", n) // QUESTION: is x right?
		}
		curr = curr.next
	}

	return curr.data, nil
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


func emptyListError() error {
	return fmt.Errorf("singleylinkedlist: empty list")
}

//// method to reverse the singly linked list
//// note: the nodes should be moved and not just the values in the nodes
//func reverse
//fmt.Println("Not implemented"
//end
//
//// Advanced Exercises
//// returns the value at the middle element in the singly linked list
//func find_middle_value
//fmt.Println("Not implemented"
//end
//
//// find the nth node from the end and return its value
//// assume indexing starts at 0 while counting to n
//func find_nth_from_end(n)
//fmt.Println("Not implemented"
//end
//
//// checks if the linked list has a cycle. A cycle exists if any node in the
//// linked list links to a node already visited.
//// returns true if a cycle is found, false otherwise.
//func has_cycle
//fmt.Println("Not implemented"
//end
//
//// Creates a cycle in the linked list for testing purposes
//// Assumes the linked list has at least one node
//func create_cycle
//return if @head == nil // don't do anything if the linked list is empty
//
//// navigate to last node
//current = @head
//while current.next != nil
//current = current.next
//end
//
//current.next = @head // make the last node link to first node
//end
//end
