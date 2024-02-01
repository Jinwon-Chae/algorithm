package main

import "fmt"

type Node[T any] struct {
	next *Node[T] // 다음 노드를 가리키는 포인터
	prev *Node[T] // 있으면 double linked list, 없으면  single linked list
	val  T        // value
}

func Append[T any](node *Node[T], next *Node[T]) *Node[T] {
	node.next = next
	return next
}

func main() {
	// 총 3개의 노드
	root := &Node[int]{nil, nil, 10}                //int인 경우 16바이트
	root.next = &Node[int]{nil, root, 20}           // next에 다음 노드의 포인터 할당
	root.next.next = &Node[int]{nil, root.next, 30} // next의 next에 다음 노드의 포인터 할당

	for n := root; n != nil; n = n.next {
		fmt.Printf("node val: %d\n", n.val)
	}

	tail := root.next.next
	for n := tail; n != nil; n = n.prev {
		fmt.Printf("prev node val: %d\n", n.val)
	}
}
