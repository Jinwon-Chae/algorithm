package queue

type QueueSlice[T any] struct {
	arr []T
}

func NewSliceQueue[T any]() *QueueSlice[T] {
	return &QueueSlice[T]{}
}

func (q *QueueSlice[T]) Push(val T) {
	q.arr = append(q.arr, val)
}

func (q *QueueSlice[T]) Pop() T {
	var front T
	if len(q.arr) == 0 {
		return front
	}
	front = q.arr[0]
	q.arr = q.arr[1:]

	return front
}
