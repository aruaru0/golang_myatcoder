package main

import (
	"bufio"
	"fmt"
	"math"
	"math/bits"
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

const inf int = 9

func solve() {
	h, w := getI(), getI()
	g := make([]int, h)
	for i := 0; i < h; i++ {
		s := getS()
		for j := 0; j < w; j++ {
			if s[j] == '#' {
				g[i] |= 1 << j
			}
		}
	}

	// テーブルを作成（dp[S] -> g[0]をSにするためのコスト
	dp := make([]int, 1<<w)
	for s := 0; s < 1<<w; s++ {
		if (s | g[0]) == g[0] {
			dp[s] = bits.OnesCount(uint(s ^ g[0]))
		} else {
			dp[s] = inf
		}
	}

	// 次の行から全パターンを捜査
	for i := 1; i < h; i++ {
		tmp := make([]int, 1<<w)
		for s := 0; s < 1<<w; s++ {
			tmp[s] = inf
			if (s | g[i]) != g[i] {
				continue
			}

			inc := bits.OnesCount(uint(s ^ g[i]))

			for t := 0; t < 1<<w; t++ {
				if (t | g[i-1]) != g[i-1] {
					continue
				}

				and := t & s
				if (and & (and << 1)) != 0 {
					continue
				}

				tmp[s] = min(tmp[s], dp[t]+inc)
			}
		}

		dp = tmp
	}

	ans := inf
	for s := 0; s < 1<<w; s++ {
		ans = min(ans, dp[s])
	}

	out(ans)
}

func main() {
	defer wr.Flush()
	sc.Split(bufio.ScanWords)
	sc.Buffer([]byte{}, math.MaxInt32)
	// this template is new version.
	// use getI(), getS(), getInts(), getF()
	T := getI()

	for ; T > 0; T-- {
		solve()
	}
}
