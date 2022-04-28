package lab4

import "math/rand"

type QuickSort struct {
	Arr []int
}

func NewQuickSort() *QuickSort {
	qs := &QuickSort{}
	return qs
}

func (qs *QuickSort) GenerateRandom(size, dup int) {
	if dup >= size {
		dup = size - 1
	}
	numSize := size - dup
	qs.Arr = make([]int, size)
	for i := 0; i < size; i++ {
		if i < numSize {
			qs.Arr[i] = i
		} else {
			qs.Arr[i] = rand.Intn(numSize)
		}
	}
	for i := 0; i < numSize; i++ {
		j := rand.Intn(numSize)
		qs.Arr[i], qs.Arr[j] = qs.Arr[j], qs.Arr[i]
	}
}

func (qs *QuickSort) Qsort(l, r int) {
	if l < r {
		rand := rand.Intn(r-l) + l
		qs.Arr[l], qs.Arr[rand] = qs.Arr[rand], qs.Arr[l]
		x := l
		for y := l + 1; y < r; y++ {
			if qs.Arr[y] < qs.Arr[l] {
				x++
				qs.Arr[x], qs.Arr[y] = qs.Arr[y], qs.Arr[x]
			}
		}
		qs.Arr[x], qs.Arr[l] = qs.Arr[l], qs.Arr[x]
		qs.Qsort(l, x)
		qs.Qsort(x+1, r)
	}
}
