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

func main() {
	defer wr.Flush()
	sc.Split(bufio.ScanWords)
	sc.Buffer([]byte{}, math.MaxInt32)
	// this template is new version.
	// use getI(), getS(), getInts(), getF()
	N, K := getI(), getI()
	a := getInts(N)

	const D int = 18
	judge := func(x int) int {
		to := make([][]int, D)
		co := make([][]int, D)
		for i := 0; i < D; i++ {
			to[i] = make([]int, N)
			co[i] = make([]int, N)
		}
		j, w, sum := 0, 0, 0
		for i := 0; i < N; i++ {
			for sum < x {
				sum += a[j]
				w++
				j = (j + 1) % N
			}
			to[0][i] = j
			co[0][i] = w
			sum -= a[i]
			w--
		}
		for i := 0; i < D-1; i++ {
			for j := 0; j < N; j++ {
				to[i+1][j] = to[i][to[i][j]]
				co[i+1][j] = min(co[i][j]+co[i][to[i][j]], N+1)
			}
		}
		res := 0
		for si := 0; si < N; si++ {
			v, cost := si, 0
			for i := 0; i < D; i++ {
				if (K>>i)&1 != 0 {
					cost += co[i][v]
					v = to[i][v]
				}
			}
			if cost > N { // cost > Nということは、ここからスタートできない＝切れ目にならないということ
				res++
			}

		}
		return res
	}

	ac, wa := 0, 0
	for i := 0; i < N; i++ {
		wa += a[i]
	}
	for ac+1 != wa {
		wj := (ac + wa) / 2
		if judge(wj) == N {
			wa = wj
		} else {
			ac = wj
		}
	}

	out(ac, judge(ac))

}
