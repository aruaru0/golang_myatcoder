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

var N int
var A []int

type TP [3]int

var mp = make(map[TP]float64)

func solve(c, state, win int) float64 {
	tmp := [3]int{c, state, win}
	v, ok := mp[tmp]
	if ok {
		return v
	}
	if state == 0 {
		s, t := 0, 0
		for i := 0; i < N; i++ {
			if win&(1<<i) != 0 {
				s += A[i]
			} else {
				t += A[i]
			}
		}
		if (c == 0 && s > t) || (c == 1 && s <= t) {
			mp[tmp] = 1
			return mp[tmp]
		} else {
			mp[tmp] = 0
			return mp[tmp]
		}
	}

	res := 0.0
	for a := 0; a < N; a++ {
		if state&(1<<a) == 0 {
			continue
		}
		if A[a] == 1 {
			nw := win
			if c == 0 {
				nw ^= 1 << a
				res = math.Max(res, (1 - solve(1-c, state^(1<<a), nw)))
				continue
			}
		}
		score := 1.0
		for b := 0; b < N; b++ {
			if state&(1<<b) == 0 {
				continue
			}
			coef := float64(A[a]*A[b]) / float64(A[a]+A[b]-1)
			val := 0.0
			if c == 0 {
				val += coef * (1 - solve(1-c, state^(1<<a), win^(1<<a))) / float64(A[a])
				val += coef * (solve(c, state^(1<<b), win)) * (float64(A[a]-1) / float64(A[a])) / float64(A[b])
			} else {
				val += coef * (1 - solve(1-c, state^(1<<a), win)) / float64(A[a])
				val += coef * (solve(c, state^(1<<b), win^(1<<b))) * (float64(A[a]-1) / float64(A[a])) / float64(A[b])
			}
			score = math.Min(score, val)
		}
		res = math.Max(res, score)
	}
	mp[tmp] = res
	return res
}

// 解けなかったので写経。確率問題は不得意
func main() {
	defer wr.Flush()
	sc.Split(bufio.ScanWords)
	sc.Buffer([]byte{}, 1000000)
	// this template is new version.
	// use getI(), getS(), getInts(), getF()
	N = getI()
	A = getInts(N)

	ans := solve(0, (1<<N)-1, 0)
	out(ans)
}
