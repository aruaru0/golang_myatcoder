package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

func out(x ...interface{}) {
	fmt.Println(x...)
}

var sc = bufio.NewScanner(os.Stdin)

func getInt() int {
	sc.Scan()
	i, e := strconv.Atoi(sc.Text())
	if e != nil {
		panic(e)
	}
	return i
}

func getString() string {
	sc.Scan()
	return sc.Text()
}

// min, max, asub, absなど基本関数
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func asub(a, b int) int {
	if a > b {
		return a - b
	}
	return b - a
}

func abs(a int) int {
	if a >= 0 {
		return a
	}
	return -a
}

func lowerBound(a []int, x int) int {
	idx := sort.Search(len(a), func(i int) bool {
		return a[i] >= x
	})
	return idx
}

func upperBound(a []int, x int) int {
	idx := sort.Search(len(a), func(i int) bool {
		return a[i] > x
	})
	return idx
}

// X Y :
var X, Y int
var memo [][2][2]bool

func rec(x, y, dir int, a []int) bool {
	if memo[len(a)][dir][0] == true {
		return memo[len(a)][dir][1]
	}
	ret := false
	if len(a) == 0 {
		if x == X && y == Y {
			ret = true
		}
		return ret
	}
	// out(x, y, dir, a[0])

	if a[0] != 0 {
		if dir == 0 {
			ret = rec(x+a[0], y, dir, a[1:])
			ret = ret || rec(x-a[0], y, dir, a[1:])
		} else {
			ret = rec(x, y+a[0], dir, a[1:])
			ret = ret || rec(x, y-a[0], dir, a[1:])
		}
	} else {
		if dir == 0 {
			ret = rec(x, y, 1, a[1:])
		} else {
			ret = rec(x, y, 0, a[1:])
		}
	}
	memo[len(a)][dir][0] = true
	memo[len(a)][dir][0] = ret
	return ret
}

func main() {
	sc.Split(bufio.ScanWords)
	sc.Buffer([]byte{}, 1000000)
	s := getString()
	X, Y = getInt(), getInt()

	a := make([]int, 0)
	cnt := 0
	for _, v := range s {
		if v == 'F' {
			cnt++
		} else {
			if cnt != 0 {
				a = append(a, cnt)
			}
			a = append(a, 0)
			cnt = 0
		}
	}
	if s[len(s)-1] == 'F' {
		a = append(a, cnt)
	}

	const maxn = 8100

	var dpX [maxn][2 * maxn]bool
	var dpY [maxn][2 * maxn]bool

	dir := false
	flg := false
	cntX, cntY := 0, 0

	dpX[0][maxn] = true
	dpY[0][maxn] = true

	for i := 0; i < len(a); i++ {
		if a[i] == 0 {
			dir = !dir
		} else {
			if !dir { // x
				cntX++
				for j := 0; j < maxn*2; j++ {
					if dpX[cntX-1][j] {

						if j+a[i] < maxn*2 {
							dpX[cntX][j+a[i]] = true
						}
						if j-a[i] >= 0 && flg {
							dpX[cntX][j-a[i]] = true
						}
					}
				}
			} else {
				cntY++
				for j := 0; j < maxn*2; j++ {
					if dpY[cntY-1][j] {
						if j+a[i] < maxn*2 {
							dpY[cntY][j+a[i]] = true
						}
						if j-a[i] >= 0 && flg {
							dpY[cntY][j-a[i]] = true
						}
					}
				}
			}
		}
		flg = true
	}

	// for i := 0; i <= cntX; i++ {
	// 	out(dpX[i][maxn-4 : maxn+5])
	// }
	// out(cntX, cntY)

	if dpX[cntX][X+maxn] && dpY[cntY][Y+maxn] {
		out("Yes")
	} else {
		out("No")
	}
}
