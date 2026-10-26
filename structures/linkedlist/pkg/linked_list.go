package linkedlist

type Node struct {
	data any
	next *Node
}

type LinkedList struct {
	head *Node
	tail *Node
}

func NewLinkedList() *LinkedList {
	return &LinkedList{}
}

func (ll *LinkedList) InsertEnd(data any) {
	newNode := &Node{
		data: data,
	}
	if ll.head == nil {
		ll.head = newNode
		ll.tail = ll.head
		return
	}
	ll.tail.next = newNode
	ll.tail = newNode
}

func (ll *LinkedList) InsertFront(data any) {
	newNode := &Node{
		data: data,
	}
	if ll.head == nil {
		ll.head = newNode
		ll.tail = ll.head
		return
	}
	newNode.next = ll.head
	ll.head = newNode
}

func (ll *LinkedList) Travesal() []any {
	items := []any{}
	current := ll.head
	for current != nil {
		items = append(items, current.data)
		current = current.next
	}

	return items
}

func (ll *LinkedList) Delete(value any) bool {
	var prev *Node
	current := ll.head

	for current != nil && current.data != value {
		prev = current
		current = current.next
	}

	if current == nil {
		return false
	}

	if prev == nil {
		ll.head = current.next
	} else {
		prev.next = current.next
	}

	if current == ll.tail {
		ll.tail = prev
	}

	current.next = nil
	return true
}
