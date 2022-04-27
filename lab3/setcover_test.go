package lab3

import (
	"testing"
)

func TestGenerateRandSlice(t *testing.T) {
	res := GenerateRandSlice(100, 20)
	t.Log(res)
}

func TestGenerateSubset(t *testing.T) {
	sc := NewSetCover()
	sc.GenerateSubset()
	t.Log(sc.SubSets)
}

func TestGreedySetCover(t *testing.T) {
	sc := NewSetCover()
	sc.GenerateSubset()
	t.Log(sc.SubSets)
	t.Log(sc.GreedySetCover())
}
func TestLPSetCover(t *testing.T) {
	sc := NewSetCover()
	sc.GenerateSubset()
	t.Log(sc.SubSets)
	t.Log(sc.LPSetCover())
}

func TestSetCover(t *testing.T) {
	sc := NewSetCover()
	sc.GenerateSubset()
	t.Log(sc.SubSets)
	t.Log(sc.GreedySetCover())
	t.Log(sc.LPSetCover())
}

// func TestSpecialSetCover(t *testing.T) {
// 	sc := NewSetCover()
// 	// sc.GenerateSubset()
// 	sc.SubSets = append(sc.SubSets, []int{6, 8, 1, 0, 7}, []int{0, 7, 1, 8}, []int{2, 3, 4, 9}, []int{5})
// 	sc.SetNum = len(sc.SubSets)
// 	t.Log(sc.SubSets)
// 	t.Log(sc.GreedySetCover())
// 	t.Log(sc.LPSetCover())
// }
