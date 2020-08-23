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

type pair struct {
	x, y int
}

func main() {
	sc.Split(bufio.ScanWords)
	sc.Buffer([]byte{}, 1000000)
	_, _, M := getInt(), getInt(), getInt()
	h := make(map[int]int)
	w := make(map[int]int)
	xy := make(map[pair]int)
	for i := 0; i < M; i++ {
		hh, ww := getInt(), getInt()
		h[hh]++
		w[ww]++
		xy[pair{hh, ww}]++
	}
	maxH := 0
	maxW := 0
	hpos := make([]int, 0, 300000)
	wpos := make([]int, 0, 300000)
	for i, n := range h {
		if maxH < n {
			hpos = make([]int, 0, 300000)
			hpos = append(hpos, i)
			maxH = n
		} else if maxH == n {
			hpos = append(hpos, i)
		}
	}
	for i, n := range w {
		if maxW < n {
			wpos = make([]int, 0, 300000)
			wpos = append(wpos, i)
			maxW = n
		} else if maxW == n {
			wpos = append(wpos, i)
		}
	}

	ans := maxH + maxW - 1
	for _, hh := range hpos {
		for _, ww := range wpos {
			if xy[pair{hh, ww}] == 0 {
				out(ans + 1)
				return
			}
		}
	}
	out(ans)
}
