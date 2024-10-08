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

const inf = int(1e18)

func f(s, K int) int {
	// out("f()")
	used := make([]int, N)
	for i := 0; i < N; i++ {
		used[i] = -1
	}
	cur := s
	tot := 0
	loop := 0
	cost := []int{0}
	ret := -inf // ループするまでの最大の値
	// ループができるまで探索する
	for i := 0; i < N; i++ {
		next := p[cur]
		if used[next] != -1 {
			loop = used[next]
			break
		}
		tot += c[next]
		used[next] = i
		cur = next
		if len(cost) <= K {
			ret = max(ret, tot)
		}
		cost = append(cost, tot)
	}
	// out(cost, loop)
	if len(cost)-1 >= K {
		return ret
	}

	pre := cost[loop] - cost[0]           // ループ手前
	lps := cost[len(cost)-1] - cost[loop] // ループ部分の点数
	lpsize := len(cost) - 1 - loop        // ループのサイズ
	npre := loop
	nlps := (K - npre) / lpsize
	rest := K - loop - lpsize*nlps
	if rest == 0 && nlps != 0 {
		rest = lpsize
		nlps--
	}

	// out(pre, lps, "size", npre, lpsize, "cnt", nlps, "ret", ret, "rest", rest)

	cnt := npre
	ret2 := pre
	if lps > 0 {
		ret2 += lps * nlps
		cnt += lps * lpsize
	}
	// out("ret2", ret2)
	maxCost := -inf
	maxCnt := 1
	for i := 0; i <= rest; i++ {
		// out(cost[i+loop], cost[loop])
		if maxCost <= cost[i+loop]-cost[loop] {
			maxCost = cost[i+loop] - cost[loop]
			maxCnt = i
		}
	}
	ret2 += maxCost
	cnt += maxCnt

	if cnt == 0 {
		ret2 = c[p[s]]
	}
	// out(cost, ret, ret2, "cnt", cnt)
	return max(ret, ret2)
}

var N, K int
var p []int
var c []int

func main() {
	defer wr.Flush()
	sc.Split(bufio.ScanWords)
	sc.Buffer([]byte{}, math.MaxInt32)
	// this template is new version.
	// use getI(), getS(), getInts(), getF()
	N, K = getI(), getI()

	p = make([]int, N)
	for i := 0; i < N; i++ {
		p[i] = getI() - 1
	}
	c = getInts(N)
	// out(p)
	// out(c, K)

	ans := -inf

	for i := 0; i < N; i++ {
		ret := f(i, K)
		ans = max(ans, ret)
	}
	out(ans)
}
