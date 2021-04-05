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

const mod = int(1e9 + 7)

func main() {
	defer wr.Flush()
	sc.Split(bufio.ScanWords)
	sc.Buffer([]byte{}, math.MaxInt32)
	// this template is new version.
	// use getI(), getS(), getInts(), getF()
	N, M := getI(), getI()
	a := getInts(N)
	b := getInts(M)
	sort.Ints(a)
	sort.Ints(b)
	ma := make(map[int]bool)
	mb := make(map[int]bool)
	for i := range a {
		if ma[a[i]] {
			out(0)
			return
		}
		ma[a[i]] = true
	}
	for i := range b {
		if mb[b[i]] {
			out(0)
			return
		}
		mb[b[i]] = true
	}
	sort.Ints(a)
	sort.Ints(b)

	// Editional通りの実装
	ans := 1
	for i := N * M; i > 0; i-- {
		okA := ma[i]
		okB := mb[i]
		if okA && okB { //　両方出現する場合は一意に定まる
			continue
		}
		if okA { // Aだけにある場合は、iより大きなBに隠れている
			j := sort.SearchInts(b, i)
			ans *= M - j
			ans %= mod
			continue
		}
		if okB { // Bだけにある場合は、iより大きなAに隠れている
			j := sort.SearchInts(a, i)
			ans *= N - j
			ans %= mod
			continue
		}
		// どちらにも無い場合は、iより大きなAとBに隠れている
		j := sort.SearchInts(a, i)
		k := sort.SearchInts(b, i)
		ans *= (N-j)*(M-k) - (N*M - i)
		ans %= mod
	}
	out(ans)
}
