package lab1

import (
	"math/rand"
	"sort"
	"time"
)

func getV(x, y Node) Node {
	return Node{y.x - x.x, y.y - x.y}
}

func crossProduct(x, y Node) int {
	return x.x*y.y - x.y*y.x
}
func generateNode(n, max int) []Node {
	rand.Seed(time.Now().UnixNano())
	resNode := make([]Node, n)
	mark := make(map[Node]int, n)
	for i := range resNode {
		resNode[i].x, resNode[i].y = rand.Intn(max), rand.Intn(max)
		for mark[resNode[i]] == 1 {
			resNode[i].x, resNode[i].y = rand.Intn(max), rand.Intn(max)
		}
		mark[resNode[i]] = 1
	}
	return resNode
}
func isInTriangle(z, i, j, k Node) bool {
	// sameside := false
	vector1, vector2 := getV(i, j), getV(i, z)
	vector3, vector4 := getV(j, k), getV(j, z)
	vector5, vector6 := getV(k, i), getV(k, z)
	if (crossProduct(vector1, vector2) > 0 && crossProduct(vector3, vector4) > 0 && crossProduct(vector5, vector6) > 0) || (crossProduct(vector1, vector2) < 0 && crossProduct(vector3, vector4) < 0 && crossProduct(vector5, vector6) < 0) {
		return true
	}
	if crossProduct(vector1, vector2)*crossProduct(vector3, vector4)*crossProduct(vector5, vector6) == 0 {
		return crossProduct(vector1, vector2)*crossProduct(vector3, vector4) > 0 || crossProduct(vector1, vector2)*crossProduct(vector5, vector6) > 0 || crossProduct(vector3, vector4)*crossProduct(vector5, vector6) > 0
	}
	return false
}

func getPolarAngle(i, j, k PolarAngleNode) int {
	return ((i.x-k.x)*(j.y-k.y) - (j.x-k.x)*(i.y-k.y))
}

func calcuArea(i, j, k Node) int {
	return i.x*j.y + k.x*i.y + j.x*k.y - k.x*j.y - j.x*i.y - i.x*k.y
}

func SortPolarAngleNodes(nodes []Node) []PolarAngleNode {
	pole, min := 0, nodes[0].y
	for i := range nodes {
		if nodes[i].y < min {
			min = nodes[i].y
			pole = i
		}
	}

	polarAngleNodes := make([]PolarAngleNode, len(nodes))
	for i := range polarAngleNodes {
		polarAngleNodes[i] = PolarAngleNode{
			prevLocation: i,
		}
		polarAngleNodes[i].x = nodes[i].x
		polarAngleNodes[i].y = nodes[i].y
	}
	polarAngleNodes[0], polarAngleNodes[pole] = polarAngleNodes[pole], polarAngleNodes[0]
	b := polarAngleNodes[1:]
	sort.Slice(b, func(i, j int) bool {
		return getPolarAngle(b[i], b[j], polarAngleNodes[0]) > 0
	})

	return polarAngleNodes
}
