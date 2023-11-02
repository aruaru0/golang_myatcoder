package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)
var wr = bufio.NewWriter(os.Stdout)

func out(x ...interface{}) {
	fmt.Fprintln(wr, x...)
}

func getI() int {
	sc.Scan()
	i, e := strconv.Atoi(sc.Text())
	if e != nil {
		panic(e)
	}
	return i
}

func getF() float64 {
	sc.Scan()
	i, e := strconv.ParseFloat(sc.Text(), 64)
	if e != nil {
		panic(e)
	}
	return i
}

func getInts(N int) []int {
	ret := make([]int, N)
	for i := 0; i < N; i++ {
		ret[i] = getI()
	}
	return ret
}

func getS() string {
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

// min for n entry
func nmin(a ...int) int {
	ret := a[0]
	for _, e := range a {
		ret = min(ret, e)
	}
	return ret
}

// max for n entry
func nmax(a ...int) int {
	ret := a[0]
	for _, e := range a {
		ret = max(ret, e)
	}
	return ret
}

func chmin(a *int, b int) bool {
	if *a < b {
		return false
	}
	*a = b
	return true
}

func chmax(a *int, b int) bool {
	if *a > b {
		return false
	}
	*a = b
	return true
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

func solve() {
	n := getI()
	x := getI()
	y := getI()
	aa := getInts(n)
	if n == 1 {
		if aa[0] == y && x == 0 {
			out("Yes")
			out("L")
			return
		}
		if aa[0] == -y && x == 0 {
			out("Yes")
			out("R")
			return
		}
		out("No")
		return
	}
	yy := make([]int, 0)
	xx := make([]int, 0)
	for i := 0; i < n; i++ {
		if i%2 == 0 {
			yy = append(yy, aa[i])
			continue
		}
		xx = append(xx, aa[i])
	}
	yok, hy := half(yy, y)
	xok, hx := half(xx, x)
	if !yok || !xok {
		out("No")
		return
	}
	out("Yes")
	d := 0
	for i := 0; i < n; i++ {
		switch d {
		case 0:
			if hy[0] == 0 {
				fmt.Fprint(wr, "R")
				d = 3
			} else {
				fmt.Fprint(wr, "L")
				d = 1
			}
			hy = hy[1:]
		case 1:
			if hx[0] == 0 {
				fmt.Fprint(wr, "L")
				d = 2
			} else {
				fmt.Fprint(wr, "R")
				d = 0
			}
			hx = hx[1:]
		case 2:
			if hy[0] == 0 {
				fmt.Fprint(wr, "L")
				d = 3
				hy = hy[1:]
			} else {
				d = 1
				fmt.Fprint(wr, "R")
				hy = hy[1:]
			}
		case 3:
			if hx[0] == 0 {
				fmt.Fprint(wr, "R")
				d = 2
			} else {
				fmt.Fprint(wr, "L")
				d = 0
			}
			hx = hx[1:]
		}
	}
	out("")
}
func half(yy []int, y int) (bool, []int) {
	n := len(yy)
	if n == 1 {
		if yy[0] == -y {
			return true, []int{0}
		}
		if yy[0] == y {
			return true, []int{1}
		}
		return false, []int{}
	}
	m := n / 2
	ml := map[int]int{}
	for i := 0; i < 1<<uint(m); i++ {
		sum := 0
		for j := 0; j < m; j++ {
			sum += yy[j] * ((i>>uint(j)&1)*2 - 1)
		}
		ml[sum] = i
	}
	for i := 0; i < 1<<uint(n-m); i++ {
		sum := 0
		for j := 0; j < n-m; j++ {
			sum += yy[m+j] * ((i>>uint(j)&1)*2 - 1)
		}
		if x, ok := ml[y-sum]; ok {
			ans := make([]int, 0)
			for j := 0; j < m; j++ {
				ans = append(ans, (x >> uint(j) & 1))
			}
			for j := 0; j < n-m; j++ {
				ans = append(ans, (i >> uint(j) & 1))
			}
			return true, ans
		}
	}
	return false, []int{}
}

func main() {
	defer wr.Flush()
	sc.Split(bufio.ScanWords)
	sc.Buffer([]byte{}, math.MaxInt32)
	// this template is new version.
	// use getI(), getS(), getInts(), getF()
	solve()
}
