package linkedlist

type SingleLinkedListNode struct {
	Next *SingleLinkedListNode
	Data interface{}
}

func (node *SingleLinkedListNode) ToList() []interface{} {
	data := make([]interface{}, 0)
	cur := node
	for cur != nil {
		data = append(data, cur.Data)
		cur = cur.Next
	}
	return data
}

type SingleLinkedList struct {
	Head *SingleLinkedListNode
}

func NewSingleLinkedList() *SingleLinkedList {
	return &SingleLinkedList{
		Head: &SingleLinkedListNode{
			Next: nil,
			Data: nil,
		},
	}
}

func (list *SingleLinkedList) Insert(target int, data interface{}) *SingleLinkedList {
	cur := list.Head
	count := 0
	for cur.Next != nil && count < target {
		cur = cur.Next
		count++
	}
	newNode := &SingleLinkedListNode{
		Next: cur.Next,
		Data: data,
	}
	cur.Next = newNode
	return list
}

func (list *SingleLinkedList) Append(data interface{}) *SingleLinkedList {
	cur := list.Head
	for cur.Next != nil {
		cur = cur.Next
	}
	newNode := &SingleLinkedListNode{
		Next: cur.Next,
		Data: data,
	}
	cur.Next = newNode
	return list
}

func (list *SingleLinkedList) IsEmpty() bool {
	return list.Head.Next == nil
}

func (list *SingleLinkedList) DeleteVal(val interface{}) bool {
	cur := list.Head
	for cur.Next != nil {
		if cur.Next.Data == val {
			targetNode := cur.Next
			cur.Next = targetNode.Next
			return true
		}
		cur = cur.Next
	}
	return false
}

func (list *SingleLinkedList) ToList() []interface{} {
	data := make([]interface{}, 0)
	cur := list.Head.Next
	for cur != nil {
		data = append(data, cur.Data)
		cur = cur.Next
	}
	return data
}

func (list *SingleLinkedList) AppendList(data []interface{}) {
	for _, v := range data {
		list.Append(v)
	}
}
