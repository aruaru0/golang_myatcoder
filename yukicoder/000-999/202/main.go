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

func getInts(N int) []int {
	ret := make([]int, N)
	for i := 0; i < N; i++ {
		ret[i] = getInt()
	}
	return ret
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

func main() {
	sc.Split(bufio.ScanWords)
	N := getInt()
	x := make([]int, N)
	y := make([]int, N)
	var m [2010][2010][]int
	for i := 0; i < N; i++ {
		x[i], y[i] = getInt(), getInt()
	}
	cnt := 0
	for i := 0; i < N; i++ {
		xpos := (x[i]+19)/20 + 1
		ypos := (y[i]+19)/20 + 1
		flg := true
		// out(i, x[i], y[1], xpos, ypos)
	L0:
		for dx := -1; dx <= 1; dx++ {
			for dy := -1; dy <= 1; dy++ {
				xx := xpos + dx
				yy := ypos + dy
				// out("xx, yy", xx, yy)
				for _, e := range m[xx][yy] {
					rx := abs(x[i] - x[e])
					ry := abs(y[i] - y[e])
					// out(i, ":", x[i], y[i], "vs", e, ":", x[e], y[e], rx, ry)
					if rx*rx+ry*ry < 400 {
						flg = false
						break L0
					}
				}
			}
		}
		if flg {
			m[xpos][ypos] = append(m[xpos][ypos], i)
			cnt++
		}
	}
	out(cnt)
}
