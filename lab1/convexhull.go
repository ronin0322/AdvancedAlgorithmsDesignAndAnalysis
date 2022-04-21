package lab1

type Node struct {
	x, y int
}
type NodeSet struct {
	cap        int
	Nodes      []Node
	ConvexHull []bool
}

func (ns *NodeSet) GetNode() [][]int {
	res := make([][]int, ns.cap)
	for i, j := range ns.Nodes {
		res[i] = []int{j.x, j.y}
	}
	return res
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
	nodes := make([]Node, 0, ns.cap)
	location := make([]int, 0, ns.cap)
	for i := range ns.ConvexHull {
		if ns.ConvexHull[i] {
			nodes = append(nodes, ns.Nodes[i])
			location = append(location, i)
		}
	}
	PolarAngleNodes := SortPolarAngleNodes(nodes)
	for _, j := range PolarAngleNodes {
		res = append(res, location[j.prevLocation])
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

	polarAngleNodes := SortPolarAngleNodes(ns.Nodes)

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

func (ns *NodeSet) Divide(l, r int, isMax int) {
	maxArea, maxIndex := 0, 0
	for i := range ns.Nodes {
		if i == l || i == r {
			continue
		}
		if isMax*calcuArea(ns.Nodes[l], ns.Nodes[r], ns.Nodes[i]) > maxArea {
			maxArea = isMax * calcuArea(ns.Nodes[l], ns.Nodes[r], ns.Nodes[i])
			maxIndex = i
		}
	}
	if maxArea == 0 {
		return
	}
	ns.ConvexHull[maxIndex] = true
	ns.Divide(l, maxIndex, isMax)
	ns.Divide(maxIndex, r, isMax)
}

func (ns *NodeSet) DivideAndConquerFindConvexHull() {
	ns.convexHullInit(false)
	xMin, xMax := 0, 0
	for i := range ns.Nodes {
		if ns.Nodes[i].x < ns.Nodes[xMin].x {
			xMin = i
		}
		if ns.Nodes[i].x > ns.Nodes[xMax].x {
			xMax = i
		}
	}
	ns.ConvexHull[xMin], ns.ConvexHull[xMax] = true, true
	ns.Divide(xMin, xMax, 1)
	ns.Divide(xMin, xMax, -1)

}
