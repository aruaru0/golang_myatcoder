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
	N := getI()

	hold := make([][]int, N)
	get := make([]bool, N)
	use := make([]bool, N)
	T := make([]bool, N)

	ok := true
	for i := 0; i < N; i++ {
		t, x := getI(), getI()-1
		if t == 1 {
			T[i] = true
			hold[x] = append(hold[x], i)
		} else {
			if len(hold[x]) == 0 {
				ok = false
			} else {
				l := len(hold[x])
				get[hold[x][l-1]] = true
				hold[x] = hold[x][:l-1]
				use[i] = true
			}
		}
		// out("---------")
		// out(t, x)
		// out(hold)
		// out(get)
		// out(use)
	}

	if !ok {
		out(-1)
		return
	}

	ans := 0
	tot := 0
	for i := 0; i < N; i++ {
		if get[i] == true {
			tot++
		}
		if use[i] == true {
			tot--
		}
		ans = max(ans, tot)
	}

	out(ans)
	for i := 0; i < N; i++ {
		if T[i] == true {
			if get[i] == true {
				fmt.Fprint(wr, 1, " ")
			} else {
				fmt.Fprint(wr, 0, " ")
			}
		}
	}
	out()
}
