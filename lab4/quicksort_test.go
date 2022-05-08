/*
 * @Author: ronin0322
 * @Date: 2022-05-08 20:43:46
 * @LastEditors: ronin0322
 * @LastEditTime: 2022-05-08 20:44:58
 * @FilePath: /AdvancedAlgorithmsDesignAndAnalysis/lab4/quicksort_test.go
 * @Description:
 *
 * Copyright (c) 2022 by ronin0322, All Rights Reserved.
 */
package lab4

import (
	"log"
	"testing"
)

func TestQuickSort(t *testing.T) {
	qs := NewQuickSort()
	qs.Arr = []int{0, 5, 4, 3, 1, 2, 5, 2, 4, 3, 1, 1}
	qs.Qsort(0, len(qs.Arr))
	t.Log(qs.Arr)
}

func TestGenerateRandomQuickSort(t *testing.T) {
	qs := NewQuickSort()
	qs.GenerateRandom(1e3, 1e2)
	t.Log(qs.Arr)
	log.Println("ready")
	qs.Qsort(0, len(qs.Arr))
	t.Log(qs.Arr)
	log.Println("ok")
}
func TestTreQuickSort(t *testing.T) {
	qs := NewQuickSort()
	qs.Arr = []int{0, 5, 4, 3, 1, 2, 5, 2, 4, 3, 1, 1}
	qs.TreQsort(0, len(qs.Arr))
	t.Log(qs.Arr)
}
