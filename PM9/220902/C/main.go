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

// 区間の[L, R)の左区間がPになるように分割する
func f(x []int, L, R, P int) int {
	l, r := L, R
	// 二分探索
	for l+1 != r {
		m := (l + r) / 2
		if x[m]-x[L] > P {
			r = m
		} else {
			l = m
		}
	}
	// 見つかった区間が丁度Pでなければ-1
	if x[l]-x[L] != P {
		return -1
	}
	return l
}

func main() {
	defer wr.Flush()
	sc.Split(bufio.ScanWords)
	sc.Buffer([]byte{}, math.MaxInt32)
	// this template is new version.
	// use getI(), getS(), getInts(), getF()
	N, P, Q, R := getI(), getI(), getI(), getI()
	a := getInts(N)
	b := make([]int, N+1)
	for i := 0; i < N; i++ {
		b[i+1] += b[i] + a[i]
	}

	tot := P + Q + R
	l, r := 0, 0
	for l <= N {
		// 区間の和がtotより小さい間rを加算
		for r <= N && b[r]-b[l] < tot {
			r++
		}
		if r > N {
			break
		}
		// もし区間和がちょっきりtotの場合
		if b[r]-b[l] == tot {
			// [l, r+1)の区間で値が左区間の値がPになる部分を探索
			x := f(b, l, r+1, P)
			// 見つかったら
			if x != -1 {
				// [x: r+1)の区間で左側がQになる区間を探索
				y := f(b, x, r+1, Q)
				// out(x, y)
				if y != -1 {
					// 見つかれば最後の区間がRになるか調べて（なるはずだけど一応）yesとして終了
					if b[r]-b[y] == R {
						out("Yes")
						return
					}
				}
			}
		}
		// 区間先頭を１つ左にずらす
		l++
	}
	// 見つからない場合はNo
	out("No")
}
