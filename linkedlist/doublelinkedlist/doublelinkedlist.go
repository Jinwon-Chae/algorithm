package doublelinkedlist

type Node[T any] struct {
	next *Node[T] // 다음 노드를 가리키는 포인터
	prev *Node[T] // 이전 노드를 가리키는 포인터

	Value T // value
}

type LinkedList[T any] struct {
	root *Node[T] // 첫 노드
	tail *Node[T] // 마지막 노드

	count int // 리스트의 노드 갯수
}

func (l *LinkedList[T]) PushBack(value T) {
	node := &Node[T]{
		Value: value,
	}

	l.count++
	if l.tail == nil {
		l.root = node
		l.tail = node
		return
	}

	l.tail.next = node
	node.prev = l.tail
	l.tail = node
}

func (l *LinkedList[T]) PushFront(value T) {
	node := &Node[T]{
		Value: value,
	}

	l.count++
	if l.root == nil {
		l.root = node
		l.tail = node
		return
	}

	node.next = l.root
	l.root.prev = node
	l.root = node
}

func (l *LinkedList[T]) Front() *Node[T] {
	return l.root
}

func (l *LinkedList[T]) Back() *Node[T] {
	return l.tail
}

func (l *LinkedList[T]) Count() int {
	return l.count
}

func (l *LinkedList[T]) GetAt(idx int) *Node[T] {
	if idx > l.count {
		return nil
	}

	i := 0
	for node := l.root; node != nil; node = node.next {
		if i == idx {
			return node
		}

		i++
	}

	return nil
}

func (l *LinkedList[T]) InsertAfter(node *Node[T], value T) {
	if !l.isIncluded(node) {
		return
	}

	newNode := &Node[T]{
		Value: value,
	}

	nextNode := node.next
	node.next = newNode
	newNode.next = nextNode
	newNode.prev = node

	if nextNode != nil {
		nextNode.prev = newNode
	}

	l.count++

	if node == l.tail {
		l.tail = newNode
	}
}

func (l *LinkedList[T]) isIncluded(node *Node[T]) bool {
	inner := l.root

	for ; inner != nil; inner = inner.next {
		if inner == node {
			return true
		}
	}

	return false
}

func (l *LinkedList[T]) InsertBefore(node *Node[T], value T) {
	if !l.isIncluded(node) {
		return
	}

	if node == l.root {
		l.PushFront(value)
		return
	}

	prevNode := node.prev
	newNode := &Node[T]{
		Value: value,
	}

	prevNode.next = newNode
	newNode.next = node
	newNode.prev = prevNode
	node.prev = newNode
	l.count++
}

func (l *LinkedList[T]) PopFront() *Node[T] {
	if l.root == nil {
		return nil
	}

	n := l.root
	l.root = n.next

	if l.root != nil {
		l.root.prev = nil
	} else {
		l.tail = nil
	}
	n.next = nil

	l.count--

	return n
}

func (l *LinkedList[T]) Remove(node *Node[T]) {
	if node == l.root {
		l.PopFront()
		return
	} else if node == l.tail {
		l.PopBack()
		return
	}

	if !l.isIncluded(node) {
		return
	}

	prev := node.prev
	next := node.next

	if prev != nil {
		prev.next = next
	}

	if next != nil {
		next.prev = prev
	}

	node.next, node.prev = nil, nil
	l.count--
}

func (l *LinkedList[T]) PopBack() *Node[T] {
	if l.tail == nil {
		return nil
	}

	n := l.tail
	l.tail = n.prev

	if l.tail != nil {
		l.tail.next = nil
	} else {
		l.root = nil
	}
	n.prev = nil

	l.count--

	return n
}
