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

func outSlice[T any](s []T) {
	if len(s) == 0 {
		return
	}
	for i := 0; i < len(s)-1; i++ {
		fmt.Fprint(wr, s[i], " ")
	}
	fmt.Fprintln(wr, s[len(s)-1])
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

func max64(x, y uint64) uint64 {
	if x < y {
		return uint64(y)
	}
	return uint64(x)
}

func main() {
	defer wr.Flush()
	sc.Split(bufio.ScanWords)
	sc.Buffer([]byte{}, math.MaxInt32)
	// this template is new version.
	// use getI(), getS(), getInts(), getF()
	n, m := getI(), getI()
	a := make([][]int, n)
	for i := 0; i < n; i++ {
		a[i] = getInts(n)
	}

	enumPath := func(grid [][]int) [][]uint64 {
		res := make([][]uint64, n)
		var dfs func(i, j int, x uint64)
		dfs = func(i, j int, x uint64) {
			if i+j == n-1 {
				res[i] = append(res[i], x)
				return
			}
			x = x*10 + uint64(grid[i][j])
			if i+1 < n {
				dfs(i+1, j, x)
			}
			if j+1 < n {
				dfs(i, j+1, x)
			}
		}
		dfs(0, 0, 0)
		return res
	}

	ss := enumPath(a)

	// 反転
	for i := range a {
		for j := 0; j < n/2; j++ {
			a[i][j], a[i][n-1-j] = a[i][n-1-j], a[i][j]
		}
	}
	for i := 0; i < n/2; i++ {
		a[i], a[n-1-i] = a[n-1-i], a[i]
	}

	ts := enumPath(a)

	for i := 0; i < n/2; i++ {
		ts[i], ts[n-1-i] = ts[n-1-i], ts[i]
	}

	var ans uint64 = 0
	f := func(s, t []uint64, d int) {
		ten := uint64(1)
		for i := 0; i < n; i++ {
			ten = ten * 10 % uint64(m)
		}
		for i := range s {
			s[i] = s[i] % uint64(m) * ten % uint64(m)
		}
		for i := range t {
			y := t[i]
			x := uint64(d)
			for j := 0; j < n-1; j++ {
				x = (x*10 + y%10) % uint64(m)
				y /= 10
			}
			t[i] = x
		}
		sort.Slice(s, func(i, j int) bool { return s[i] < s[j] })
		sort.Slice(t, func(i, j int) bool { return t[i] < t[j] })
		for _, x := range s {
			i := sort.Search(len(t), func(i int) bool { return t[i] >= uint64(m)-x })
			if i > 0 {
				ans = max64(ans, x+t[i-1])
			}
			ans = max64(ans, (x+t[len(t)-1])%uint64(m))
		}
	}

	for i := 0; i < n; i++ {
		f(ss[i], ts[i], a[n-1-i][i])
	}

	fmt.Println(ans)
}
