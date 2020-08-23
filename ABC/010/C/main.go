package main

import (
	"bufio"
	"fmt"
	"math"
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

const eps = 1e-10

func main() {
	sc.Split(bufio.ScanWords)
	Txa, Tya, Txb, Tyb, T, V := getInt(), getInt(), getInt(), getInt(), getInt(), getInt()
	n := getInt()
	x := make([]int, n)
	y := make([]int, n)
	for i := 0; i < n; i++ {
		x[i], y[i] = getInt(), getInt()
	}

	L := T * V

	flg := false
	for i := 0; i < n; i++ {
		dx := abs(x[i] - Txa)
		dy := abs(y[i] - Tya)
		l0 := math.Hypot(float64(dx), float64(dy))
		dx = abs(x[i] - Txb)
		dy = abs(y[i] - Tyb)
		l1 := math.Hypot(float64(dx), float64(dy))

		if l0+l1 <= float64(L)+eps {
			flg = true
		}
		if l0+l1 <= float64(L)-eps {
			flg = true
		}
	}
	if flg {
		out("YES")
		return
	}
	out("NO")
}
