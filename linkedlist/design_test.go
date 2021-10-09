package linkedlist

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var linkedList MyLinkedList

func init() {
	linkedList = Constructor()
}

func TestMyLinkedList_Get(t *testing.T) {
	linkedList.AddAtHead(1)
	linkedList.AddAtTail(3)
	linkedList.AddAtIndex(1, 2)           //链表变为1-> 2-> 3
	assert.Equal(t, 2, linkedList.Get(1)) //返回2
	linkedList.DeleteAtIndex(1)           //现在链表是1-> 3
	assert.Equal(t, 3, linkedList.Get(1)) //返回2
}
