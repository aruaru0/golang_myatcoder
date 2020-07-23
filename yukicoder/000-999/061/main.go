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

// GCD : greatest common divisor (GCD) via Euclidean algorithm
func GCD(a, b int) int {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}

func solve() {
	w, h, d := getInt(), getInt(), getInt()
	mx, my := getInt(), getInt()
	hx, hy := getInt(), getInt()
	vx, vy := getInt(), getInt()
	var xy [20][20]int
	var g int
	if vx == 0 {
		g = abs(vy)
	} else if vy == 0 {
		g = abs(vx)
	} else {
		g = GCD(abs(vx), abs(vy))
	}
	d *= g
	vx /= g
	vy /= g
	for i := 0; i < min(d+1, 1100); i++ {
		dx := abs(hx + i*vx)
		dy := abs(hy + i*vy)
		dx %= 2 * w
		if dx > w {
			dx = 2*w - dx
		}
		dy %= 2 * h
		if dy > h {
			dy = 2*h - dy
		}
		xy[dx][dy] = 1
	}
	if xy[mx][my] == 1 {
		out("Hit")
		return
	}
	out("Miss")
}

func main() {
	sc.Split(bufio.ScanWords)
	Q := getInt()
	for i := 0; i < Q; i++ {
		solve()
	}
}
