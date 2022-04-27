package lab3

import (
	"math/rand"
	"time"

	"github.com/draffensperger/golp"
)

type SetCover struct {
	NodeNum int
	SetNum  int
	SetSize int
	X       []int
	SubSets [][]int
}

func NewSetCover() *SetCover {
	rand.Seed(time.Now().UnixNano())
	sc := &SetCover{}
	sc.NodeNum = 50
	sc.SetSize = 10
	// sc.SetNum = n
	return sc
}

func GenerateRandSlice(n, m int) []int {
	if n < m {
		return []int{}
	}
	randSlice := make([]int, 0, m)
	mark := make(map[int]int, m)
	for len(mark) < m {
		x := rand.Intn(n)
		_, ok := mark[x]
		if ok {
			continue
		}
		mark[x] = 1
		randSlice = append(randSlice, x)
	}
	return randSlice
}

func (sc *SetCover) GenerateSetFromAllSet(a [][]int, m int) []int {
	mark := make(map[int]int, m)
	allSet := map[int]int{}
	for _, i := range a {
		for _, j := range i {
			allSet[j] = 1
		}
	}
	set := make([]int, 0, m)
	for len(mark) < m {
		o := rand.Intn(len(allSet))
		x := 0
		for k := range allSet {
			if o == 0 {
				x = k
				break
			}
			o--
		}
		// x := randmap.Key(allSet).(int)
		_, ok := mark[x]
		if ok {
			continue
		}
		delete(allSet, x)
		mark[x] = 1
		set = append(set, x)
	}
	return set
}

func Union(a, b []int) []int {
	union := make([]int, 0)
	union = append(union, a...)
	mark := make(map[int]int)
	for _, i := range a {
		mark[i] = 1
	}
	for _, i := range b {
		_, ok := mark[i]
		if ok {
			continue
		}
		union = append(union, i)
	}
	return union
}

func (sc *SetCover) GenerateSubset() {
	// rand.Seed(time.Now().UnixNano())
	sc.SubSets = make([][]int, 0)
	sc.SubSets = append(sc.SubSets, GenerateRandSlice(sc.NodeNum, sc.SetSize))

	for {
		sItemLen := rand.Intn(sc.SetSize-1) + 1
		sLastLen := rand.Intn(sItemLen)
		setA := sc.GenerateSetFromAllSet(sc.SubSets, sLastLen)
		setB := GenerateRandSlice(sc.NodeNum, sItemLen-sLastLen)

		sc.SubSets = append(sc.SubSets, Union(setA, setB))

		allSet := map[int]int{}
		for _, i := range sc.SubSets {
			for _, j := range i {
				allSet[j] = 1
			}
		}
		if sc.NodeNum-len(allSet) < sc.SetSize {
			lastSet := make([]int, 0, sc.SetSize)
			for i := 0; i < sc.NodeNum; i++ {
				_, ok := allSet[i]
				if ok {
					continue
				}
				lastSet = append(lastSet, i)
			}
			sc.SubSets = append(sc.SubSets, lastSet)
			break
		}
	}
	sc.SetNum = len(sc.SubSets)
}

func (sc *SetCover) GreedySetCover() [][]int {
	subSets := make([][]int, len(sc.SubSets))
	copy(subSets, sc.SubSets)
	coverSets := make([][]int, 0, sc.SetNum)
	covered := make(map[int]int)
	for len(covered) < sc.NodeNum {
		maxSet := 0
		maxNum := 0
		for index, set := range subSets {
			num := 0
			for _, i := range set {
				_, ok := covered[i]
				if ok {
					continue
				}
				num++
			}
			if num > maxNum {
				maxSet = index
				maxNum = num
			}
		}
		for _, i := range subSets[maxSet] {
			covered[i] = 1
		}
		coverSets = append(coverSets, make([]int, len(subSets[maxSet])))
		copy(coverSets[len(coverSets)-1], subSets[maxSet])
		subSets = append(subSets[:maxSet], subSets[maxSet+1:]...)

	}
	return coverSets
}

func (sc *SetCover) LPSetCover() [][]int {
	subSets := make([][]int, len(sc.SubSets))
	copy(subSets, sc.SubSets)
	coverSets := make([][]int, 0, sc.SetNum)
	allF := make(map[int]int)
	max := 0
	for _, i := range subSets {
		for _, j := range i {
			allF[j]++
			if allF[j] > max {
				// f = j
				max = allF[j]
			}
		}
	}
	A := make([][]float64, sc.NodeNum)
	for i := range A {
		A[i] = make([]float64, sc.SetNum)
		for j, set := range subSets {
			for _, k := range set {
				if k == i {
					A[i][j] += 1.0
				}
			}
		}
	}
	B := make([]float64, sc.NodeNum)
	for i := range B {
		B[i] = 1.0
	}
	C := make([]float64, sc.SetNum)
	for i := range C {
		C[i] = 1.0
	}
	lp := golp.NewLP(0, sc.SetNum)

	for i := range A {
		lp.AddConstraint(A[i], golp.GE, B[i])
	}
	lp.SetObjFn(C)

	lp.Solve()
	vars := lp.Variables()
	for i := 0; i < lp.NumCols(); i++ {
		// log.Printf("x%v = %v\n", i+1, vars[i])
		if vars[i] >= 1/float64(max) {
			coverSets = append(coverSets, subSets[i])
		}
	}
	return coverSets
}
