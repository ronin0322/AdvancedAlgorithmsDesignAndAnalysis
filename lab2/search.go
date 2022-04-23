package lab2

import (
	"math"
)

type Node struct {
	X int `json:"x"`
	Y int `json:"y"`
}
type FCost struct {
	G, H     float64
	Previous Node
}
type Search struct {
	Length, Width   int
	Desert          [][]int
	River           [][]int
	Barrier         [][]int
	Start, Terminal []int
	Open            map[Node]FCost
	Close           map[Node]FCost
}

func NewSearch(l, w int, s, t []int, desert, river, barrier [][]int) *Search {
	srh := &Search{}
	srh.Length = l
	srh.Width = w
	srh.Desert = make([][]int, l)
	for i := range srh.Desert {
		srh.Desert[i] = make([]int, w)
	}
	srh.River = make([][]int, l)
	for i := range srh.River {
		srh.River[i] = make([]int, w)
	}
	srh.Barrier = make([][]int, l)
	for i := range srh.Barrier {
		srh.Barrier[i] = make([]int, w)
	}
	copy(srh.Desert, desert)
	copy(srh.River, river)
	copy(srh.Barrier, barrier)
	srh.Start = []int{s[0], s[1]}
	srh.Terminal = []int{t[0], t[1]}
	srh.Open = make(map[Node]FCost, srh.Length*srh.Width)
	srh.Close = make(map[Node]FCost, srh.Length*srh.Width)
	srh.Open[Node{s[0], s[1]}] = FCost{0, srh.GetHCost(Node{s[0], s[1]}), Node{-1, -1}}
	return srh
}

func (srh *Search) GetHCost(cur Node) float64 {
	return math.Sqrt(math.Pow(float64(cur.X-srh.Terminal[0]), 2) + math.Pow(float64(cur.Y-srh.Terminal[1]), 2))
}

func (srh *Search) FindMinFCostNode() Node {
	min, minNode := 0.0, Node{-1, -1}
	for k, v := range srh.Open {
		if minNode.X == -1 {
			min = v.G + v.H
			minNode = k
		} else if v.G+v.H < min {
			min = v.G + v.H
			minNode = k
		}
	}
	return minNode
}

func (srh *Search) InSide(cur Node) bool {
	return cur.X >= 0 && cur.X < srh.Length && cur.Y >= 0 && cur.Y < srh.Width
}

func (srh *Search) GetDRCost(cur Node) float64 {
	var drcost float64
	drcost += 4.0 * float64(srh.Desert[cur.X][cur.Y])
	drcost += 2.0 * float64(srh.River[cur.X][cur.Y])
	return drcost
}

func (srh *Search) GetGCost(i, j int) float64 {
	return math.Sqrt(math.Abs(float64(i)) + math.Abs(float64(j)))
}
func (srh *Search) IsBarrier(cur Node) bool {
	return srh.Barrier[cur.X][cur.Y] == 1
}
func (srh *Search) NextStep() {
	minCost := srh.FindMinFCostNode()
	for i := -1; i <= 1; i++ {
		for j := -1; j <= 1; j++ {
			if i == 0 && j == 0 {
				continue
			}
			nxt := Node{minCost.X + i, minCost.Y + j}
			_, inClose := srh.Close[nxt]
			if !inClose && srh.InSide(nxt) && !srh.IsBarrier(nxt) {
				nxtFCost := FCost{}
				nxtFCost.G = srh.Open[minCost].G + srh.GetDRCost(nxt) + srh.GetGCost(i, j)
				nxtFCost.H = srh.GetHCost(nxt)
				nxtFCost.Previous = minCost
				v, ok := srh.Open[nxt]
				if ok {
					if v.G <= nxtFCost.G {
						continue
					}
				}
				srh.Open[nxt] = nxtFCost
			}
		}
	}
	srh.Close[minCost] = srh.Open[minCost]
	delete(srh.Open, minCost)
}

func (srh *Search) FindTerminal() {
	for {
		_, ok := srh.Open[Node{srh.Terminal[0], srh.Terminal[1]}]
		if ok {
			break
		}
		if len(srh.Open) == 0 {
			break
		}
		srh.NextStep()
	}
}

func (srh *Search) PrintPath(cur Node) []Node {
	path := make([]Node, 0, srh.Length*srh.Width)
	// cur := Node{srh.Terminal[0], srh.Terminal[1]}
	v, ok := srh.Open[cur]
	if !ok {
		return path
	}
	for cur.X != -1 {
		path = append(path, cur)
		cur = v.Previous
		v, ok = srh.Close[cur]
		if !ok {
			return path
		}
	}
	return path
}
