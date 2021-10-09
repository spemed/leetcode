package linkedlist

import "testing"

var singleLinkedList *SingleLinkedList

func init() {
	singleLinkedList = NewSingleLinkedList()
	singleLinkedList.Append(1)
	singleLinkedList.Append(2)
	singleLinkedList.Append(3)
	singleLinkedList.Append(4)
	singleLinkedList.Append(5)
	//singleLinkedList.Insert(-1, 11)
	singleLinkedList.Insert(3, 11)
}

func TestSingleLinkedList_ToList(t *testing.T) {
	for _, v := range singleLinkedList.ToList() {
		t.Log(v)
	}
}

func TestSingleLinkedList_DeleteVal(t *testing.T) {
	singleLinkedList.DeleteVal(1)
	singleLinkedList.DeleteVal(2)
	singleLinkedList.DeleteVal(3)
	for _, v := range singleLinkedList.ToList() {
		t.Log(v)
	}
}
