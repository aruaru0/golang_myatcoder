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

func getStrings(N int) []string {
	ret := make([]string, N)
	for i := 0; i < N; i++ {
		ret[i] = getS()
	}
	return ret
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

// 値を圧縮した配列を返す
func compressArray(a []int) []int {
	m := make(map[int]int)
	for _, e := range a {
		m[e] = 1
	}
	b := make([]int, 0)
	for e := range m {
		b = append(b, e)
	}
	sort.Ints(b)
	for i, e := range b {
		m[e] = i
	}

	ret := make([]int, len(a))
	for i, e := range a {
		ret[i] = m[e]
	}
	return ret
}

func solve() {
	s := getS()
	k := getI()

	n := len(s)
	const inf = int(1e18)

	// 2次元スライスのディープコピー用ヘルパー
	makeEmp := func() [][3]int {
		res := make([][3]int, k+2)
		for i := 0; i < k+2; i++ {
			for j := 0; j < 3; j++ {
				res[i][j] = inf
			}
		}
		return res
	}

	emp := make([][3]int, k+2)
	for i := 0; i < k+2; i++ {
		for j := 0; j < 3; j++ {
			emp[i][j] = inf
		}
	}
	dp := emp
	dp[0][0] = 0
	ch := "ABC?"
	for i := 0; i < n; i++ {
		old := dp
		dp = makeEmp()
		for j := 0; j < k+2; j++ {
			for a := 0; a < 3; a++ {
				for ci := 0; ci < len(ch); ci++ {
					c := ch[ci]
					now := old[j][a]
					if c != '?' {
						now++
					}
					if c == '?' {
						c = s[i]
					}
					nj, na := j, a
					if ch[a] == c {
						na++
					} else if c == 'A' {
						na = 1
					} else {
						na = 0
					}
					if na == 3 {
						na = 0
						nj++
					}
					if nj > k+1 {
						continue
					}
					chmin(&dp[nj][na], now)
				}
			}
		}
		if i >= 2 && s[i-2:i+1] == "ABC" {
			dp = dp[1:]
			dp = append(dp, [3]int{inf, inf, inf})
		}
	}

	ans := inf
	for i := 0; i < 3; i++ {
		ans = min(ans, dp[k][i])
	}
	if ans == inf {
		out(-1)
	} else {
		out(ans)
	}
}

func main() {
	defer wr.Flush()
	sc.Split(bufio.ScanWords)
	sc.Buffer([]byte{}, math.MaxInt32)
	// this template is new version.
	// use getI(), getS(), getInts(), getF()
	t := getI()
	for i := 0; i < t; i++ {
		solve()
	}

}
