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
	S := getS()
	t := make([][]byte, 2)
	flip := 0
	for _, e := range S {
		if e == 'R' {
			flip ^= 1
			continue
		}
		l := len(t[flip])
		if l != 0 && t[flip][l-1] == byte(e) {
			t[flip] = t[flip][:l-1]
		} else {
			t[flip] = append(t[flip], byte(e))
		}
		// out("-------")
		// out(string(t[0]))
		// out(string(t[1]))
	}
	ans := make([]byte, 0)
	if flip == 1 {
		for i := len(t[0]) - 1; i >= 0; i-- {
			ans = append(ans, t[0][i])
		}
		for i := 0; i < len(t[1]); i++ {
			if len(ans) != 0 && ans[len(ans)-1] == t[1][i] {
				ans = ans[:len(ans)-1]
			} else {
				ans = append(ans, t[1][i])
			}
		}
	} else {
		for i := len(t[1]) - 1; i >= 0; i-- {
			ans = append(ans, t[1][i])
		}
		for i := 0; i < len(t[0]); i++ {
			if len(ans) != 0 && ans[len(ans)-1] == t[0][i] {
				ans = ans[:len(ans)-1]
			} else {
				ans = append(ans, t[0][i])
			}
		}
	}
	out(string(ans))
}
