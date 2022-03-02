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

func main() {
	defer wr.Flush()
	sc.Split(bufio.ScanWords)
	sc.Buffer([]byte{}, math.MaxInt32)
	// this template is new version.
	// use getI(), getS(), getInts(), getF()
	N, M := getI(), getI()
	a := make([][]int, M)
	pos := make([][]int, N) // 色のある場所を格納
	for i := 0; i < M; i++ {
		k := getI()
		a[i] = getInts(k)
		for j := 0; j < k; j++ {
			a[i][j]--
			pos[a[i][j]] = append(pos[a[i][j]], i)
		}
	}

	cnt := make([]int, N)
	q := make([]int, 0)
	// 一番上にある色の数をcntに入れる
	for i := 0; i < M; i++ {
		t := a[i][0]
		cnt[t]++
		// ２つあれば　色をキューに入れる
		if cnt[t] == 2 {
			q = append(q, t)
		}
	}

	n := 0
	for len(q) != 0 {
		// キューから取り出す
		x := q[0]
		q = q[1:]
		n++
		// 取り出した色を除外し、下にある色をカウントし２つあればキューに入れる
		for i := 0; i < 2; i++ {
			p := pos[x][i]
			a[p] = a[p][1:]
			if len(a[p]) != 0 {
				t := a[p][0]
				cnt[t]++
				if cnt[t] == 2 {
					q = append(q, t)
				}
			}
		}
	}
	if n == N {
		out("Yes")
	} else {
		out("No")
	}
}
