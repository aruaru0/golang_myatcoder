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
	N := getI()
	a := getInts(N)

	b := make([][20]int, N)
	for i := 0; i < N; i++ {
		for j := 0; j < 20; j++ {
			b[i][j] = (a[i] >> j) & 1
		}
	}
	// for i := 0; i < N; i++ {
	// 	out(b[i])
	// }
	x := make([]int, 20)
	cnt := 0
	l := 0
	ans := 0
	for i := 0; i < N; i++ {
		ok := true
		for j := 0; j < 20; j++ {
			x[j] += b[i][j]
			if x[j] == 2 {
				ok = false
			}
		}
		if ok == false {
			ans += cnt * (cnt + 1) / 2
			// out("-----", cnt, l, ans)
			for {
				ok := true
				for j := 0; j < 20; j++ {
					x[j] -= b[l][j]
					if x[j] == 2 {
						ok = false
					}
				}
				l++
				cnt--
				if ok == true {
					break
				}
			}
			rest := i - l
			ans -= rest * (rest + 1) / 2
		}
		cnt++
		// out(cnt, a[i])
	}
	ans += cnt * (cnt + 1) / 2
	out(ans)
}
