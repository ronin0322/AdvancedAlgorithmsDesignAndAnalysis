package lab1

import "testing"

func TestIsInTriangle(t *testing.T) {
	z, i, j, k := Node{1, 2}, Node{0, 0}, Node{10, 10}, Node{0, 10}
	res := true
	t.Log("TestIsInTriangle")
	{
		resp := isInTriangle(z, i, j, k)
		if resp == res {
			t.Log("correct", resp)
		} else {
			t.Error("error", resp)
		}
	}
}
func TestGenerateNode(t *testing.T) {

	// res := true
	t.Log("TestGenerateNode")
	{
		resp := generateNode(10, 100)
		t.Log(resp)
		// if resp == res {
		// 	t.Log("correct", resp)
		// } else {
		// 	t.Error("error", resp)
		// }
	}
}
