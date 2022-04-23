package main

import (
	"AdvancedAlgorithmsDesignAndAnalysis/lab1"
	"AdvancedAlgorithmsDesignAndAnalysis/lab2"
	"log"
	"sync"

	"github.com/gin-contrib/pprof"

	"github.com/gin-gonic/gin"
)

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}
func index(c *gin.Context) {
	c.JSON(200, gin.H{
		"res": "ok",
	})
}

type Node struct {
	X int `json:"x"`
	Y int `json:"y"`
}

func Simple(inNodes []lab1.Node, res [][]Node, wg *sync.WaitGroup) {
	ns := lab1.NewNodeSet(inNodes)
	ns.SimpleFindConvexHull()
	log.Println("SimpleFindConvexHull          ", ns.ConvexHullPrint())
	nodes := ns.GetNode()
	for _, x := range nodes {
		res[3] = append(res[3], Node{x[0], x[1]})
	}
	for _, x := range ns.ConvexHullPrint() {
		res[0] = append(res[0], Node{nodes[x][0], nodes[x][1]})
	}
	wg.Done()
}

func GrahamScan(inNodes []lab1.Node, res [][]Node, wg *sync.WaitGroup) {
	ns := lab1.NewNodeSet(inNodes)
	ns.GrahamScanFindConvexHull()
	log.Println("GrahamScanFindConvexHull      ", ns.ConvexHullPrint())
	nodes := ns.GetNode()
	for _, x := range nodes {
		res[4] = append(res[4], Node{x[0], x[1]})
	}
	for _, x := range ns.ConvexHullPrint() {
		res[1] = append(res[1], Node{nodes[x][0], nodes[x][1]})
	}
	wg.Done()
}

func DivideAndConquer(inNodes []lab1.Node, res [][]Node, wg *sync.WaitGroup) {
	ns := lab1.NewNodeSet(inNodes)
	ns.DivideAndConquerFindConvexHull()
	log.Println("DivideAndConquerFindConvexHull", ns.ConvexHullPrint())
	nodes := ns.GetNode()
	// log.Println(nodes)
	for _, x := range nodes {
		res[5] = append(res[5], Node{x[0], x[1]})
	}
	for _, x := range ns.ConvexHullPrint() {
		res[2] = append(res[2], Node{nodes[x][0], nodes[x][1]})
	}
	wg.Done()
}

func lab1index(c *gin.Context) {
	n := 3000
	inNodes := lab1.InitNode(n)
	wg := sync.WaitGroup{}
	wg.Add(2)
	res := make([][]Node, 6)

	// go Simple(inNodes, res, &wg)
	go GrahamScan(inNodes, res, &wg)
	go DivideAndConquer(inNodes, res, &wg)

	wg.Wait()
	// ns := NewNodeSet()
	// log.Println(res)
	c.JSON(200, res)
}

func lab2Input1() *lab2.Search {
	n, m := 7, 13
	desert := make([][]int, n)
	river := make([][]int, n)
	barrier := make([][]int, n)
	for i := range desert {
		desert[i] = make([]int, m)
	}
	for i := range river {
		river[i] = make([]int, m)
	}
	for i := range barrier {
		barrier[i] = make([]int, m)
	}
	// desert[2][0] = 1
	// desert[2][1] = 1
	// river[2][2] = 1
	barrier[0][3] = 1
	barrier[1][3] = 1
	barrier[2][4] = 1
	barrier[3][4] = 1
	barrier[4][4] = 1
	barrier[4][5] = 1
	barrier[5][5] = 1
	barrier[6][5] = 1
	srh := lab2.NewSearch(n, m, []int{3, 0}, []int{4, 11}, desert, river, barrier)
	return srh
}

func lab2Input2() *lab2.Search {
	n, m := 21, 41
	desert := make([][]int, n)
	river := make([][]int, n)
	barrier := make([][]int, n)
	for i := range desert {
		desert[i] = make([]int, m)
	}
	for i := range river {
		river[i] = make([]int, m)
	}
	for i := range barrier {
		barrier[i] = make([]int, m)
	}
	// desert[2][0] = 1
	// desert[2][1] = 1
	a := []int{25, 25, 26, 27, 27, 27, 28, 28, 30}
	b := []int{40, 40, 40, 40, 37, 36, 33, 33, 33}
	for i := 0; i <= 8; i++ {
		for j := a[i]; j <= b[i]; j++ {
			desert[i][j] = 1
		}
	}
	desert[2][35] = 0
	desert[3][34] = 0
	desert[4][33] = 0
	// river[2][2] = 1
	river[2][35] = 1
	river[3][34] = 1
	river[4][33] = 1
	river[5][34] = 1
	c := []int{34, 34, 34, 33, 33, 33, 33}
	d := []int{35, 35, 36, 36, 35, 37, 36}
	for i := 6; i <= 12; i++ {
		for j := c[i-6]; j <= d[i-6]; j++ {
			river[i][j] = 1
		}
	}
	for i := 13; i <= 20; i++ {
		for j := 33; j <= 36; j++ {
			river[i][j] = 1
		}
	}
	river[11][35] = 0
	river[12][34] = 0
	river[13][33] = 0
	barrier[1][4] = 1
	barrier[1][8] = 1
	barrier[2][8] = 1
	barrier[4][9] = 1

	for i := 1; i <= 11; i++ {
		if i == 7 {
			continue
		}
		barrier[3][i] = 1
	}
	barrier[6][8] = 1
	barrier[6][9] = 1
	for i := 3; i <= 8; i++ {
		barrier[7][i] = 1
	}

	barrier[8][3] = 1
	for i := 8; i <= 13; i++ {
		barrier[i][6] = 1
	}

	barrier[8][8] = 1
	barrier[10][8] = 1
	barrier[11][8] = 1
	barrier[11][3] = 1
	barrier[12][3] = 1
	barrier[12][4] = 1
	barrier[12][5] = 1
	for i := 13; i <= 20; i++ {
		if i == 18 {
			continue
		}
		barrier[i][4] = 1
	}
	for i := 11; i <= 16; i++ {
		barrier[i][9] = 1
	}
	barrier[14][10] = 1
	for i := 5; i <= 8; i++ {
		barrier[16][i] = 1
	}
	barrier[18][8] = 1
	barrier[19][8] = 1
	barrier[20][8] = 1
	for i := 0; i <= 20; i++ {
		if i >= 9 && i <= 12 {
			continue
		}
		barrier[i][13] = 1
	}
	barrier[14][12] = 1
	for i := 11; i <= 13; i++ {
		for j := 20; j <= 22; j++ {
			barrier[i][j] = 1
		}
	}
	for i := 16; i <= 17; i++ {
		for j := 25; j <= 26; j++ {
			barrier[i][j] = 1
		}
	}
	barrier[11][29] = 1
	barrier[12][32] = 1
	barrier[14][32] = 1
	barrier[8][37] = 1
	barrier[10][37] = 1

	srh := lab2.NewSearch(n, m, []int{11, 5}, []int{1, 36}, desert, river, barrier)
	return srh
	// srh.FindTerminal()
	// t.Log(srh.PrintPath())
}

func lab2index(c *gin.Context) {
	srh := lab2Input2()
	srh.FindTerminal()
	srh1 := lab2Input2()
	srh2 := lab2Input2()
	copy(srh2.Start, srh1.Terminal)
	copy(srh2.Terminal, srh1.Start)
	delete(srh2.Open, lab2.Node{X: srh1.Start[0], Y: srh1.Start[1]})
	srh2.Open[lab2.Node{X: srh2.Start[0], Y: srh2.Start[1]}] = lab2.FCost{0, srh.GetHCost(lab2.Node{srh2.Start[0], srh2.Start[1]}), lab2.Node{-1, -1}}
	check := func() bool {
		min, minNode := 0.0, lab2.Node{-1, -1}
		for k, v1 := range srh1.Open {
			v2, ok := srh2.Open[k]
			if ok && (minNode.X == -1 || v1.G+v1.H+v2.G+v2.H < min) {
				min = v1.G + v1.H + v2.G + v2.H
				minNode = k
			}
		}
		if minNode.X != -1 {
			c.JSON(200, gin.H{
				"path":     srh.PrintPath(lab2.Node{X: srh.Terminal[0], Y: srh.Terminal[1]}),
				"desert":   srh.Desert,
				"barrier":  srh.Barrier,
				"river":    srh.River,
				"start":    srh.Start,
				"terminal": srh.Terminal,
				"path1":    srh1.PrintPath(minNode),
				"path2":    srh2.PrintPath(minNode),
			})
			return true
		}
		return false
	}
	for {
		if check() {
			return
		}
		srh1.NextStep()
		if check() {
			return
		}
		srh2.NextStep()
	}

}

func main() {
	r := gin.Default()
	r.Use(CORSMiddleware())
	r.GET("/", index)
	r.GET("/lab1", lab1index)
	r.GET("/lab2", lab2index)
	pprof.Register(r)
	r.Run()
}

// var cpuprofile = flag.String("cpuprofile", "", "write cpu profile to `file`")
// var memprofile = flag.String("memprofile", "", "write memory profile to `file`")

// func main() {
// 	flag.Parse()
// 	if *cpuprofile != "" {
// 		f, err := os.Create(*cpuprofile)
// 		if err != nil {
// 			log.Fatal("could not create CPU profile: ", err)
// 		}
// 		defer f.Close() // error handling omitted for example
// 		if err := pprof.StartCPUProfile(f); err != nil {
// 			log.Fatal("could not start CPU profile: ", err)
// 		}
// 		defer pprof.StopCPUProfile()
// 	}
// 	// ... rest of the program ...
// 	lab1.Output()
// 	// ... program end ...

// 	if *memprofile != "" {
// 		f, err := os.Create(*memprofile)
// 		if err != nil {
// 			log.Fatal("could not create memory profile: ", err)
// 		}
// 		defer f.Close() // error handling omitted for example
// 		runtime.GC()    // get up-to-date statistics
// 		if err := pprof.WriteHeapProfile(f); err != nil {
// 			log.Fatal("could not write memory profile: ", err)
// 		}
// 	}
// }
