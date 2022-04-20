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
	n := 99
	wg := sync.WaitGroup{}
	wg.Add(2)
	nodes := generateNode(n, n)
	ns := NewNodeSet(nodes)
	// log.Println(" ", nodes)
	go func() {
		ns.SimpleFindConvexHull()
		log.Println("SimpleFindConvexHull    ", ns.ConvexHullPrint())
		wg.Done()
	}()
	go func() {
		ns.GrahamScanFindConvexHull()
		log.Println("GrahamScanFindConvexHull", ns.ConvexHullPrint())
		wg.Done()
	}()
	wg.Wait()

}
