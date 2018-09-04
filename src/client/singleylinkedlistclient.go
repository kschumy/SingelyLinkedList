package main

import (
	"fmt"
	"github.com/kschumy/SingelyLinkedList/src"
)

func main() {

	// Create an object of linked list class
	myLinkedList := singleylinkedlist.List{}

	myLinkedList.Visit() // should do nothing

	// Add some values to the linked list
	fmt.Println("Adding 5, 3 and 1 to the linked list.")
	myLinkedList.Insert(5)
	myLinkedList.Visit() // should print: 5
	myLinkedList.Insert(3)
	myLinkedList.Visit() // should print: 3, 5
	myLinkedList.Insert(1)

	// print all elements
	fmt.Println("Printing elements in the linked list:")
	myLinkedList.Visit() // should print: 1, 3, 5

	// Find the value at the nth node
	fmt.Println("Confirming values in the linked list using find_nth_from_beginning method.")

	value, _ := myLinkedList.FindNthFromBeginning(2)
	if value != 5 {
		fmt.Println("BUG: Value at index 2 should be 5 and is", value)
	}

	value, _ = myLinkedList.FindNthFromBeginning(1)
	if value != 3 {
		fmt.Println("BUG: Value at index 1 should be 3 and is", value)
	}

	value, _ = myLinkedList.FindNthFromBeginning(0)
	if value != 1 {
		fmt.Println("BUG: Value at index 0 should be 1 and is", value)
	}

	// print all elements
	fmt.Println("Printing elements in the linked list:")
	myLinkedList.Visit()

	// Insert ascending
	fmt.Println("Adding 4 in ascending order.")
	myLinkedList.InsertAscending(4)
	// check newly inserted value
	fmt.Println("Checking values by calling find_nth_from_beginning method.")
	value, _ = myLinkedList.FindNthFromBeginning(2)
	if value != 4 {
		fmt.Println("BUG: Value at index 2 should be 4 and is", value)
	}

	value, _ = myLinkedList.FindNthFromBeginning(3)
	if value != 5 {
		fmt.Println("BUG: Value at index 3 should be 5 and is", value)
	}

	value, _ = myLinkedList.FindNthFromBeginning(1)
	if value != 3 {
		fmt.Println("BUG: Value at index 1 should be 3 and is", value)
	}

	// Insert ascending
	fmt.Println("Adding 6 in ascending order.")
	myLinkedList.InsertAscending(6)

	// print all elements
	fmt.Println("Printing elements in the linked list:")
	myLinkedList.Visit()

	// validate length
	fmt.Println("Confirming length of the linked list.")
	myLinkedListLength := myLinkedList.Length()
	if myLinkedListLength != 5 {
		fmt.Println("BUG: Length should be 5 and not", myLinkedListLength)
	}


	// find min and max
	fmt.Println("Confirming min and max values in the linked list.")
	min,_ := myLinkedList.FindMin()
	if min != 1 {
		fmt.Println("BUG: Min value should be 1 and not #{min}")
	}

	max,_ := myLinkedList.FindMax()
	if max != 6 {
		fmt.Println("BUG: Max value should be 6 and not #{max}")
	}


	// delete value
	fmt.Println("Deleting node with value 5 from the linked list.")
	myLinkedList.Delete(5)
	// print all elements
	fmt.Println("Printing elements in the linked list:")
	myLinkedList.Visit()
	// validate length
	fmt.Println("Confirming length of the linked list.")
	myLinkedListLength = myLinkedList.Length()
	if myLinkedListLength != 4 {
		fmt.Println("BUG: Length should be 4 and not", myLinkedListLength)
	}

	// delete value
	fmt.Println("Deleting node with value 1 from the linked list.")
	myLinkedList.Delete(1)
	// print all elements
	fmt.Println("Printing elements in the linked list:")
	myLinkedList.Visit()
	// validate length
	fmt.Println("Confirming length of the linked list.")
	myLinkedListLength = myLinkedList.Length()
	if myLinkedListLength != 3 {
		fmt.Println("BUG: Length should be 3 and not", myLinkedListLength)
	}

	// find middle element
	fmt.Println("Confirming middle value in the linked list.")
	middle := myLinkedList.FindMiddleValue()
	if middle != 4 {
		fmt.Println("BUG: Middle value should be 4 and not #{middle}")
	}




}










//// reverse the linked list
//fmt.Println("Reversing the linked list.")
//myLinkedList.reverse
//// print all elements
//fmt.Println("Printing elements in the linked list:"
//myLinkedList.visit
//// verify the reversed list
//fmt.Println("Verifying the reveresed linked list by calling find_nth_from_beginning method.")
//value = myLinkedList.find_nth_from_beginning(2)
//fmt.Println("BUG: Value at index 2 should be 3 and is #{value}" if value != 3
//value = myLinkedList.find_nth_from_beginning(1)
//fmt.Println("BUG: Value at index 1 should be 4 and is #{value}" if value != 4
//value = myLinkedList.find_nth_from_beginning(0)
//fmt.Println("BUG: Value at index 0 should be 6 and is #{value}" if value != 6
//
//// nth from the end
//fmt.Println("Verifying find_nth_from_end method.")
//value = myLinkedList.find_nth_from_end(0)
//fmt.Println("BUG: Value at index 0 from the end, should be 3 and is #{value}" if value != 3
//value = myLinkedList.find_nth_from_end(1)
//fmt.Println("BUG: Value at index 1 from the end, should be 4 and is #{value}" if value != 4
//value = myLinkedList.find_nth_from_end(2)
//fmt.Println("BUG: Value at index 2 from the end, should be 6 and is #{value}" if value != 6
//
//// check for cycle
//fmt.Println("Checking the linked list for cycle.")
//fmt.Println("BUG: Should not have a cycle.") if myLinkedList.has_cycle
//
//// create cycle and then test for it
//fmt.Println("Creating a cycle in the linked list.")
//myLinkedList.create_cycle
//fmt.Println("Checking the linked list for cycle.")
//fmt.Println("BUG: Should not have a cycle.") if !(myLinkedList.has_cycle)
