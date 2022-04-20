package lab1

import (
	"sort"
)

type Node struct {
	x, y int
}
type NodeSet struct {
	cap        int
	Nodes      []Node
	ConvexHull []bool
}

func NewNodeSet(nodes []Node) *NodeSet {
	ns := &NodeSet{}
	ns.cap = len(nodes)
	ns.ConvexHull = make([]bool, ns.cap)
	ns.convexHullInit(true)
	ns.Nodes = make([]Node, ns.cap)
	copy(ns.Nodes, nodes)
	return ns
}
func (ns *NodeSet) convexHullInit(flag bool) {
	for i := range ns.ConvexHull {
		ns.ConvexHull[i] = flag
	}
}
func (ns *NodeSet) ConvexHullPrint() []int {
	res := make([]int, 0, ns.cap)
	for i := range ns.ConvexHull {
		if ns.ConvexHull[i] {
			res = append(res, i)
		}
	}
	return res
}
func (ns *NodeSet) SimpleFindConvexHull() {
	ns.convexHullInit(true)
	for i := range ns.Nodes {
		for j := range ns.Nodes {
			if i == j {
				continue
			}
			for k := range ns.Nodes {
				if i == k || j == k {
					continue
				}
				for z := range ns.Nodes {
					if i == z || j == z || k == z {
						continue
					}
					if isInTriangle(ns.Nodes[i], ns.Nodes[j], ns.Nodes[k], ns.Nodes[z]) {
						ns.ConvexHull[i] = false
					}
				}
			}
		}
	}
}

type PolarAngleNode struct {
	Node
	prevLocation int
}

func (ns *NodeSet) GrahamScanFindConvexHull() {
	ns.convexHullInit(false)
	pole, min := 0, ns.Nodes[0].y
	for i := range ns.Nodes {
		if ns.Nodes[i].y < min {
			min = ns.Nodes[i].y
			pole = i
		}
	}

	polarAngleNodes := make([]PolarAngleNode, ns.cap)
	for i := range polarAngleNodes {
		polarAngleNodes[i] = PolarAngleNode{
			prevLocation: i,
		}
		polarAngleNodes[i].x = ns.Nodes[i].x
		polarAngleNodes[i].y = ns.Nodes[i].y
	}
	polarAngleNodes[0], polarAngleNodes[pole] = polarAngleNodes[pole], polarAngleNodes[0]
	b := polarAngleNodes[1:]
	sort.Slice(b, func(i, j int) bool {
		return getPolarAngle(b[i], b[j], polarAngleNodes[0]) > 0
	})

	stack := make([]PolarAngleNode, 0, len(polarAngleNodes))

	stack = append(stack, polarAngleNodes[0], polarAngleNodes[1])
	for i := 2; i < len(polarAngleNodes); i++ {
		for getPolarAngle(polarAngleNodes[i], stack[len(stack)-1], stack[len(stack)-2]) >= 0 {
			stack = stack[:len(stack)-1]
			if len(stack) < 3 {
				break
			}
		}
		stack = append(stack, polarAngleNodes[i])
	}

	for i := range stack {
		ns.ConvexHull[stack[i].prevLocation] = true
	}
}
