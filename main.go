package main

import (
	"AdvancedAlgorithmsDesignAndAnalysis/lab1"
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
	n := 100
	inNodes := lab1.InitNode(n)
	wg := sync.WaitGroup{}
	wg.Add(3)
	res := make([][]Node, 6)

	go Simple(inNodes, res, &wg)
	go GrahamScan(inNodes, res, &wg)
	go DivideAndConquer(inNodes, res, &wg)
	
	wg.Wait()
	// ns := NewNodeSet()
	// log.Println(res)
	c.JSON(200, res)
}
func main() {
	r := gin.Default()
	r.Use(CORSMiddleware())
	r.GET("/", index)
	r.GET("/lab1", lab1index)
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
