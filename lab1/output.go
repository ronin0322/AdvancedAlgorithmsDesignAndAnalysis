package lab1

import (
	"log"
	"math/rand"
	"sort"
	"sync"
)

func generate(n int) []int {
	res := make([]int, n)
	for i := range res {
		res[i] = rand.Intn(n)
	}
	return res
}
func bsort(a []int, wg *sync.WaitGroup) {
	sort.Slice(a, func(i, j int) bool {
		return a[i] < a[j]
	})
	// for i := 0; i < len(a)-1; i++ {
	// 	for j := i + 1; j < len(a); j++ {
	// 		if a[i] > a[j] {
	// 			a[i], a[j] = a[j], a[i]
	// 		}
	// 	}
	// }
	wg.Done()
}
func qsort(a []int, wg *sync.WaitGroup) {
	sort.Ints(a)
	wg.Done()
}
func Output() {
	// f, _ := os.OpenFile("cpu.pprof", os.O_CREATE|os.O_RDWR, 0644)
	// defer f.Close()
	// pprof.StartCPUProfile(f)
	// defer pprof.StopCPUProfile()
	// res := []int{5, 4, 3, 2, 1}
	// bsort(res)
	// log.Println(res)
	n := 1000000
	wg := sync.WaitGroup{}
	wg.Add(2)
	nums := generate(n)
	// log.Println(nums)
	a, b := make([]int, n), make([]int, n)
	copy(a, nums)
	copy(b, nums)
	log.Println(b[0])
	log.Println(a[0])
	go bsort(a, &wg)
	go qsort(b, &wg)
	wg.Wait()
	log.Println(b[0])
	log.Println(a[0])
}
