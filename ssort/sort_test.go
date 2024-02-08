package ssort

import (
	"math/rand"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestQuickSort(t *testing.T) {
	arr := make([]int, 100)
	for i := 0; i < 100; i++ {
		arr[i] = rand.Intn(100)
	}

	assert.False(t, IsSorted(arr))
	QuickSort(arr)
	assert.True(t, IsSorted(arr))
}

func TestMergeSort(t *testing.T) {
	arr := make([]int, 0, 100)
	for i := 0; i < 100; i++ {
		arr[i] = rand.Intn(100)
	}

	sorted := MergeSort(arr)
	assert.True(t, IsSorted(sorted))
}

func TestInsertSort(t *testing.T) {
	arr := make([]int, 0, 100)
	for i := 0; i < 100; i++ {
		BinaryInsertSort[int](arr, rand.Intn(100))
	}

	t.Log(arr)
	assert.True(t, IsSorted(arr))
}

func BenchmarkQuickSort(b *testing.B) {
	arr := make([]int, b.N)
	for i := 0; i < b.N; i++ {
		arr[i] = rand.Intn(b.N)
	}

	QuickSort(arr)
}

func BenchmarkMergeSort(b *testing.B) {
	arr := make([]int, b.N)
	for i := 0; i < b.N; i++ {
		arr[i] = rand.Intn(b.N)
	}

	MergeSort(arr)
}

func BenchmarkInsertSort(b *testing.B) {
	arr := make([]int, 0, b.N)
	for i := 0; i < b.N; i++ {
		arr = BinaryInsertSort[int](arr, rand.Intn(b.N))
	}
}
