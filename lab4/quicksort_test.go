package lab4

import "testing"

func TestQuickSort(t *testing.T) {
	qs := NewQuickSort()
	qs.Arr = []int{0, 5, 4, 3, 1, 2, 5, 2, 4, 3, 1, 1}
	qs.Qsort(0, len(qs.Arr))
	t.Log(qs.Arr)
}

func TestGenerateRandomQuickSort(t *testing.T) {
	qs := NewQuickSort()
	qs.GenerateRandom(1e6, 0)
	t.Log(qs.Arr)
	qs.Qsort(0, len(qs.Arr))
	t.Log(qs.Arr)
}
