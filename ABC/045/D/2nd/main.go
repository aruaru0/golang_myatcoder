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

type pos struct {
	x, y int
}

func main() {
	sc.Split(bufio.ScanWords)
	H, W, N := getInt(), getInt(), getInt()
	total := (H - 2) * (W - 2)

	p := make([]pos, N)
	for i := 0; i < N; i++ {
		x, y := getInt()-1, getInt()-1
		p[i] = pos{x, y}
	}

	q := make(map[pos]int)
	for _, i := range p {
		// out("black", i)
		for x := -1; x <= 1; x++ {
			for y := -1; y <= 1; y++ {
				px := i.x + x
				py := i.y + y
				// out("---", px, py)
				if px <= 0 || px >= H-1 {
					continue
				}
				if py <= 0 || py >= W-1 {
					continue
				}
				// out(px, py)
				q[pos{px, py}]++
			}
		}
	}

	J := make([]int, 10)
	for _, v := range q {
		J[v]++
		total--
	}
	out(total)
	for i := 1; i < 10; i++ {
		out(J[i])
	}
}
