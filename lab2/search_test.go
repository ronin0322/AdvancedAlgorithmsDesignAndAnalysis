package lab2

import (
	"testing"
)

func TestFindTerminal(t *testing.T) {
	n, m := 3, 3
	desert := make([][]int, n)
	river := make([][]int, n)
	barrier := make([][]int, n)
	for i := range desert {
		desert[i] = make([]int, m)
	}
	for i := range river {
		river[i] = make([]int, m)
	}
	for i := range barrier {
		barrier[i] = make([]int, m)
	}
	desert[2][0] = 1
	desert[2][1] = 1
	river[2][2] = 1
	barrier[1][1] = 1
	srh := NewSearch(n, m, []int{0, 0}, []int{n - 1, m - 1}, desert, river, barrier)
	srh.FindTerminal()
	t.Log(srh.PrintPath(Node{srh.Terminal[0], srh.Terminal[1]}))
}

func TestFindTerminal1(t *testing.T) {
	n, m := 7, 13
	desert := make([][]int, n)
	river := make([][]int, n)
	barrier := make([][]int, n)
	for i := range desert {
		desert[i] = make([]int, m)
	}
	for i := range river {
		river[i] = make([]int, m)
	}
	for i := range barrier {
		barrier[i] = make([]int, m)
	}
	// desert[2][0] = 1
	// desert[2][1] = 1
	// river[2][2] = 1
	barrier[0][3] = 1
	barrier[1][3] = 1
	barrier[2][4] = 1
	barrier[3][4] = 1
	barrier[4][4] = 1
	barrier[4][5] = 1
	barrier[5][5] = 1
	barrier[6][5] = 1
	srh := NewSearch(n, m, []int{3, 0}, []int{4, 11}, desert, river, barrier)
	srh.FindTerminal()
	t.Log(srh.PrintPath(Node{srh.Terminal[0], srh.Terminal[1]}))
}

func TestFindTerminal2(t *testing.T) {
	n, m := 21, 41
	desert := make([][]int, n)
	river := make([][]int, n)
	barrier := make([][]int, n)
	for i := range desert {
		desert[i] = make([]int, m)
	}
	for i := range river {
		river[i] = make([]int, m)
	}
	for i := range barrier {
		barrier[i] = make([]int, m)
	}
	// desert[2][0] = 1
	// desert[2][1] = 1
	a := []int{25, 25, 26, 27, 27, 27, 28, 28, 30}
	b := []int{40, 40, 40, 40, 37, 36, 33, 33, 33}
	for i := 0; i <= 8; i++ {
		for j := a[i]; j <= b[i]; j++ {
			desert[i][j] = 1
		}
	}
	desert[2][35] = 0
	desert[3][34] = 0
	desert[4][33] = 0
	// river[2][2] = 1
	river[2][35] = 1
	river[3][34] = 1
	river[4][33] = 1
	river[5][34] = 1
	c := []int{34, 34, 34, 33, 33, 33, 33}
	d := []int{35, 35, 36, 36, 35, 37, 36}
	for i := 6; i <= 12; i++ {
		for j := c[i-6]; j <= d[i-6]; j++ {
			river[i][j] = 1
		}
	}
	for i := 13; i <= 20; i++ {
		for j := 33; j <= 36; j++ {
			river[i][j] = 1
		}
	}
	barrier[1][4] = 1
	barrier[1][8] = 1
	barrier[2][8] = 1
	barrier[4][9] = 1

	for i := 1; i <= 11; i++ {
		if i == 7 {
			continue
		}
		barrier[3][i] = 1
	}
	barrier[6][8] = 1
	barrier[6][9] = 1
	for i := 3; i <= 8; i++ {
		barrier[7][i] = 1
	}

	barrier[8][3] = 1
	for i := 8; i <= 13; i++ {
		barrier[i][6] = 1
	}

	barrier[8][8] = 1
	barrier[10][8] = 1
	barrier[11][8] = 1
	barrier[11][3] = 1
	barrier[12][3] = 1
	barrier[12][4] = 1
	barrier[12][5] = 1
	for i := 13; i <= 20; i++ {
		if i == 18 {
			continue
		}
		barrier[i][4] = 1
	}
	for i := 11; i <= 16; i++ {
		barrier[i][9] = 1
	}
	barrier[14][10] = 1
	for i := 5; i <= 8; i++ {
		barrier[16][i] = 1
	}
	barrier[18][8] = 1
	barrier[19][8] = 1
	barrier[20][8] = 1
	for i := 0; i <= 20; i++ {
		if i >= 9 && i <= 12 {
			continue
		}
		barrier[i][13] = 1
	}
	barrier[14][12] = 1
	for i := 11; i <= 13; i++ {
		for j := 20; j <= 22; j++ {
			barrier[i][j] = 1
		}
	}
	for i := 16; i <= 17; i++ {
		for j := 25; j <= 26; j++ {
			barrier[i][j] = 1
		}
	}
	barrier[11][29] = 1
	barrier[12][32] = 1
	barrier[14][32] = 1
	barrier[8][37] = 1
	barrier[10][37] = 1
	srh := NewSearch(n, m, []int{11, 5}, []int{1, 36}, desert, river, barrier)
	srh.FindTerminal()
	t.Log(srh.PrintPath(Node{srh.Terminal[0], srh.Terminal[1]}))
}
