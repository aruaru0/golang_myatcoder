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

func calc(M int, s []string, a [10][3]int) int {
	for n := 0; n < 3; n++ {
		for i := 0; i < 3*M; i++ {
			t := int(s[n][i] - '0')
			if a[t][0] == i || a[t][1] == i || a[t][2] == i {
				continue
			}
			chmin(&a[t][n], i)
		}
	}

	ans := inf
	for i := 0; i < 10; i++ {
		ans = min(ans, nmax(a[i][0], a[i][1], a[i][2]))
	}
	return ans
}

const inf = int(1e18)

func NextPermutation(x sort.Interface) bool {
	n := x.Len() - 1
	if n < 1 {
		return false
	}
	j := n - 1
	for ; !x.Less(j, j+1); j-- {
		if j == 0 {
			return false
		}
	}
	l := n
	for !x.Less(j, l) {
		l--
	}
	x.Swap(j, l)
	for k, l := j+1, n; k < l; {
		x.Swap(k, l)
		k++
		l--
	}
	return true
}

func main() {
	defer wr.Flush()
	sc.Split(bufio.ScanWords)
	sc.Buffer([]byte{}, math.MaxInt32)
	// this template is new version.
	// use getI(), getS(), getInts(), getF()
	M := getI()
	s := make([]string, 3)
	for i := 0; i < 3; i++ {
		t := getS()
		s[i] = t + t + t
	}

	var a [10][3]int
	for i := 0; i < 3; i++ {
		for j := 0; j < 10; j++ {
			a[j][i] = inf
		}
	}

	p := []int{0, 1, 2}
	ans := inf
	for {
		t := []string{s[p[0]], s[p[1]], s[p[2]]}
		ret := calc(M, t, a)
		ans = min(ans, ret)
		if NextPermutation(sort.IntSlice(p)) == false {
			break
		}
	}

	if ans == inf {
		out(-1)
	} else {
		out(ans)
	}
}
