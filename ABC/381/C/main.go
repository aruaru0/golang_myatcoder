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

func change_st(c byte, st int) int {
	ret := 0
	switch st {
	case 0:
		switch c {
		case '1':
			ret = 1
			cnt_1++
		case '2':
			ret = 0
		case '/':
			ret = 0
		}
	case 1: // '1'
		switch c {
		case '1':
			ret = 1
			cnt_1++
		case '2':
			ret = 0
			cnt_1 = 0
		case '/':
			ret = 2
		}
	case 2: // '/'
		switch c {
		case '1':
			ret = 1
			cnt_1 = 1
		case '2':
			ret = 3
			cnt_2++
		case '/':
			ret = 0
			cnt_1 = 0
		}
	case 3: // '2'
		switch c {
		case '1':
			// out("state:3 ", cnt_1, cnt_2)
			ans = max(ans, min(cnt_1, cnt_2)*2+1)
			// check
			ret = 1
			cnt_2 = 0
			cnt_1 = 1
		case '2':
			ret = 3
			cnt_2++
		case '/':
			// out("state:3 ", cnt_1, cnt_2)
			ans = max(ans, min(cnt_1, cnt_2)*2+1)
			ret = 0
			cnt_2 = 0
			cnt_1 = 0
		}
	}

	return ret
}

var cnt_1, cnt_2 int
var ans int

func main() {
	defer wr.Flush()
	sc.Split(bufio.ScanWords)
	sc.Buffer([]byte{}, math.MaxInt32)
	// this template is new version.
	// use getI(), getS(), getInts(), getF()
	N := getI()
	s := getS()

	state := 0

	ans = 0
	for i := 0; i < N; i++ {
		state = change_st(s[i], state)
		if s[i] == '/' {
			ans = max(ans, 1)
		}
		// out("in", string(s[i]), "state", state, cnt_1, cnt_2)
	}
	if state == 3 {
		ans = max(ans, min(cnt_1, cnt_2)*2+1)
	}

	out(ans)
}
