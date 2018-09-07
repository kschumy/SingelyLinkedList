package main

import (
	"github.com/kschumy/SingelyLinkedList/src"
	"testing"
)

// nil
func getListEmpty() singleylinkedlist.List {
	list := singleylinkedlist.List{}
	return list
}

// 42
func getListOne() singleylinkedlist.List {
	list := singleylinkedlist.List{}
	list.Insert(42)
	return list
}

// 21 -> 42
func getListOfN(n int) singleylinkedlist.List {
	list := singleylinkedlist.List{}
	for i := n; i > 0; i-- {
		list.Insert(i)
	}
	return list
}


func TestMiddleWithListOfOne(t *testing.T) {
	list := getListOfN(1)
	answer,_ := list.FindMiddleValue()
	if answer != 1 {
		t.Errorf("TestMiddle was incorrect, got: %d, want: %d.", answer, 1)
	}

	list = getListOfN(2)
	answer,_ = list.FindMiddleValue()
	if answer != 1 {
		t.Errorf("TestMiddle was incorrect, got: %d, want: %d.", answer, 1)
	}

	list = getListOfN(3)
	answer,_ = list.FindMiddleValue()
	if answer != 2 {
		t.Errorf("TestMiddle was incorrect, got: %d, want: %d.", answer, 2)
	}

	list = getListOfN(4)
	answer,_ = list.FindMiddleValue()
	if answer != 2 {
		t.Errorf("TestMiddle was incorrect, got: %d, want: %d.", answer, 2)
	}

	list = getListOfN(5)
	answer,_ = list.FindMiddleValue()
	if answer != 3 {
		t.Errorf("TestMiddle was incorrect, got: %d, want: %d.", answer, 3)
	}
}

//func TestCycle(t *testing.T) {
//	list := getListOfN(1)
//	answer,_ := list.FindMiddleValue()
//	if answer != 1 {
//		t.Errorf("TestMiddle was incorrect, got: %d, want: %d.", answer, 1)
//	}
//
//	list = getListOfN(2)
//	answer,_ = list.FindMiddleValue()
//	if answer != 1 {
//		t.Errorf("TestMiddle was incorrect, got: %d, want: %d.", answer, 1)
//	}
//
//	list = getListOfN(3)
//	answer,_ = list.FindMiddleValue()
//	if answer != 2 {
//		t.Errorf("TestMiddle was incorrect, got: %d, want: %d.", answer, 2)
//	}
//
//	list = getListOfN(4)
//	answer,_ = list.FindMiddleValue()
//	if answer != 2 {
//		t.Errorf("TestMiddle was incorrect, got: %d, want: %d.", answer, 2)
//	}
//
//	list = getListOfN(5)
//	answer,_ = list.FindMiddleValue()
//	if answer != 3 {
//		t.Errorf("TestMiddle was incorrect, got: %d, want: %d.", answer, 3)
//	}
//}