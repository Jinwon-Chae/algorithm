package singlelinkedlist

type Node[T any] struct {
	next  *Node[T] // 다음 노드를 가리키는 포인터
	Value T        // value
}

type LinkedList[T any] struct {
	root *Node[T]
	tail *Node[T]

	count int
}

func (l *LinkedList[T]) PushBack(value T) {
	node := &Node[T]{
		Value: value,
	}

	l.count++
	if l.root == nil {
		l.root = node
		l.tail = node
		return
	}

	l.tail.next = node
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
	l.root = node
}

func (l *LinkedList[T]) Front() *Node[T] {
	return l.root
}

func (l *LinkedList[T]) Back() *Node[T] {
	return l.tail
}

func (l *LinkedList[T]) Count() int {
	node := l.root
	cnt := 0

	for ; node != nil; node = node.next {
		cnt++
	}

	return cnt
}

func (l *LinkedList[T]) Count2() int {
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

	node.next, newNode.next = newNode, node.next
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
	if node == l.root {
		l.PushFront(value)
		return
	}

	prevNode := l.findPrevNode(node)
	if prevNode == nil {
		return
	}

	newNode := &Node[T]{
		Value: value,
	}

	prevNode.next, newNode.next = newNode, node

	l.count++
}

func (l *LinkedList[T]) findPrevNode(node *Node[T]) *Node[T] {
	inner := l.root
	for ; inner != nil; inner = inner.next {
		if inner.next == node {
			return inner
		}
	}

	return nil
}

func (l *LinkedList[T]) PopFront() *Node[T] {
	if l.root == nil {
		return nil
	}

	n := l.root
	l.root.next, l.root = nil, l.root.next

	if l.root == nil {
		l.tail = nil
	}

	l.count--

	return n
}

func (l *LinkedList[T]) Remove(node *Node[T]) {
	if node == l.root {
		l.PopFront()
		return
	}

	prev := l.findPrevNode(node)
	if prev == nil {
		return
	}

	prev.next = node.next
	node.next = nil
	l.count--
}

func (l *LinkedList[T]) Reverse() {
	newL := &LinkedList[T]{}

	for l.root != nil {
		n := l.PopFront()
		newL.PushFront(n.Value)
	}
	l.count = newL.count
	l.root = newL.root
	l.tail = newL.tail
}

func (l *LinkedList[T]) Reverse2() {
	if l.root == nil {
		return
	}

	node := l.root
	next := node.next
	l.root.next = nil // 바뀌면 root가 tail이됨

	for next != nil {
		nextnext := next.next
		next.next = node
		node = next
		next = nextnext
	}

	l.root, l.tail = l.tail, l.root
}
