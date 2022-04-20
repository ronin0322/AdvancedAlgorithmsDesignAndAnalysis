package lab1

import (
	"testing"
)

func TestConvexHull(t *testing.T) {
	nodes := []Node{{1, 1}, {0, 0}, {10, 10}, {0, 10}}
	ConvexHullAlgorithms(nodes, t)
}
func TestConvexHullRand(t *testing.T) {
	nodes := generateNode(10, 100)
	ConvexHullAlgorithms(nodes, t)
	// res := []bool{false, true, true, true}

}

func ConvexHullAlgorithms(nodes []Node, t *testing.T) {
	ns := NewNodeSet(nodes)
	t.Log("TestSimpleConvexHull")
	{
		ns.SimpleFindConvexHull()
		resp := ns.ConvexHull
		t.Log("Nodes", nodes)
		t.Log("SimpleConvexHull", resp)
	}
	t.Log("TestGrahamScanFindConvexHull")
	{
		ns.GrahamScanFindConvexHull()
		resp := ns.ConvexHull
		t.Log("Nodes", nodes)
		t.Log("GrahamScanConvexHull", resp)
	}
}
