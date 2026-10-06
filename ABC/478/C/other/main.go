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

func main() {
	defer wr.Flush()
	sc.Split(bufio.ScanWords)
	sc.Buffer([]byte{}, math.MaxInt32)

	N, K := getI(), getI()
	a := getInts(N)

	// p[x]: 位置 x で終わる非減少連続長（1-indexed, p[0]=0）
	p := make([]int, N+1)
	p[1] = 1
	for i := 1; i < N; i++ {
		if a[i] >= a[i-1] {
			p[i+1] = p[i] + 1
		} else {
			p[i+1] = 1
		}
	}

	// 各窓 l の最小値・最大値を deque で求める
	wmin := make([]int, N-K+1)
	wmax := make([]int, N-K+1)
	minQ := []int{}
	maxQ := []int{}
	for j := 0; j < N; j++ {
		for len(minQ) > 0 && a[minQ[len(minQ)-1]] >= a[j] {
			minQ = minQ[:len(minQ)-1]
		}
		minQ = append(minQ, j)
		if minQ[0] <= j-K {
			minQ = minQ[1:]
		}

		for len(maxQ) > 0 && a[maxQ[len(maxQ)-1]] <= a[j] {
			maxQ = maxQ[:len(maxQ)-1]
		}
		maxQ = append(maxQ, j)
		if maxQ[0] <= j-K {
			maxQ = maxQ[1:]
		}

		if j >= K-1 {
			l := j - K + 1
			wmin[l] = a[minQ[0]]
			wmax[l] = a[maxQ[0]]
		}
	}

	for l := 0; l < N-K+1; l++ {
		r := l + K

		// 左部分 [0,l) が昇順か
		if p[l]-p[0] != l {
			continue
		}
		// 右部分 [r,N) が昇順か
		if p[N]-p[r] != N-r {
			continue
		}
		// 境界: a[l-1] <= 窓の最小値, 窓の最大値 <= a[r]
		if l > 0 && a[l-1] > wmin[l] {
			continue
		}
		if r < N && wmax[l] > a[r] {
			continue
		}

		out("Yes")
		return
	}

	out("No")
}
