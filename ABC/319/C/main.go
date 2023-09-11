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

func check(x, y, c int, a [3][3]int) bool {
	h := make([][10]int, 3)
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			h[i][a[i][j]]++
		}
	}
	w := make([][10]int, 3)
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			w[j][a[i][j]]++
		}
	}

	lr := make([][10]int, 2)
	lr[0][a[0][0]]++
	lr[0][a[1][1]]++
	lr[0][a[2][2]]++

	lr[1][a[0][2]]++
	lr[1][a[1][1]]++
	lr[1][a[2][0]]++

	// out(lr, h, w)

	ret := false
	for i := 1; i < 10; i++ {
		if h[y][i] == 2 {
			ret = true
		}
	}
	for i := 1; i < 10; i++ {
		if w[x][i] == 2 {
			ret = true
		}
	}

	for i := 0; i < 3; i++ {
		if x == i && y == i {
			for j := 1; j < 10; j++ {
				if lr[0][j] == 2 {
					ret = true
				}
			}
		}
		if x == 2-i && y == i {
			for j := 1; j < 10; j++ {
				if lr[1][j] == 2 {
					ret = true
				}
			}
		}
	}

	return ret
}

func calc_xy(n int) (int, int) {
	return n / 3, n % 3
}

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
	c := make([][]int, 3)
	for i := 0; i < 3; i++ {
		c[i] = getInts(3)
	}

	p := []int{0, 1, 2, 3, 4, 5, 6, 7, 8}

	tot := 0.0
	cnt := 0.0
	for {
		// out(x)
		var a [3][3]int
		ok := true
		for _, e := range p {
			y, x := calc_xy(e)
			ret := check(x, y, c[y][x], a)
			// out(a, ret)
			if ret == true {
				ok = false
				break
			}
			a[y][x] = c[y][x]
		}
		// out("OK=", ok)
		if ok {
			cnt++
		}
		tot++
		if NextPermutation(sort.IntSlice(p)) == false {
			break
		}
	}

	// out(cnt, tot)
	out(cnt / tot)
}
