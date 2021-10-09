package linkedlist

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"leetcode/lib/linkedlist"
)

func TestRemoveElements(t *testing.T) {
	list := linkedlist.NewSingleLinkedList()
	list.AppendList([]interface{}{1, 2, 6, 3, 4, 5, 6})
	node := removeElements(list.Head.Next, 6)
	assert.Equal(t, []interface{}{1, 2, 3, 4, 5}, node.ToList())

	list = linkedlist.NewSingleLinkedList()
	node = removeElements(list.Head.Next, 1)
	assert.Equal(t, 0, len(node.ToList()))

	list = linkedlist.NewSingleLinkedList()
	list.AppendList([]interface{}{7, 7, 7, 7})
	node = removeElements(list.Head.Next, 7)
	assert.Equal(t, 0, len(node.ToList()))

	list = linkedlist.NewSingleLinkedList()
	list.AppendList([]interface{}{1, 2, 2, 1})
	node = removeElements(list.Head.Next, 2)
	assert.Equal(t, []interface{}{1, 1}, node.ToList())
}

func TestRemoveElementsWithVirtualHead(t *testing.T) {
	list := linkedlist.NewSingleLinkedList()
	list.AppendList([]interface{}{1, 2, 6, 3, 4, 5, 6})
	node := removeElementsWithVirtualHead(list.Head.Next, 6)
	assert.Equal(t, []interface{}{1, 2, 3, 4, 5}, node.ToList())

	list = linkedlist.NewSingleLinkedList()
	node = removeElementsWithVirtualHead(list.Head.Next, 1)
	assert.Equal(t, 0, len(node.ToList()))

	list = linkedlist.NewSingleLinkedList()
	list.AppendList([]interface{}{7, 7, 7, 7})
	node = removeElementsWithVirtualHead(list.Head.Next, 7)
	assert.Equal(t, 0, len(node.ToList()))

	list = linkedlist.NewSingleLinkedList()
	list.AppendList([]interface{}{1, 2, 2, 1})
	node = removeElementsWithVirtualHead(list.Head.Next, 2)
	assert.Equal(t, []interface{}{1, 1}, node.ToList())
}
