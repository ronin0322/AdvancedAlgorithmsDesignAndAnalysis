package lab1

import (
	"math/rand"
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
