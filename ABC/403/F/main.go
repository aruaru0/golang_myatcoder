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
	if len(s) == 0 {
		return
	}
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

func chminS(s *string, t string) {
	if t == "" {
		return
	}
	if *s == "" || len(t) < len(*s) {
		*s = t
	}
}

func main() {
	defer wr.Flush()
	sc.Split(bufio.ScanWords)
	sc.Buffer([]byte{}, math.MaxInt32)
	// this template is new version.
	// use getI(), getS(), getInts(), getF()
	n := getI()
	dp_expr := make([]string, n+1)
	dp_term := make([]string, n+1)
	dp_factor := make([]string, n+1)

	for i := 1; i <= n; i = i*10 + 1 {
		str := strconv.Itoa(i)
		dp_factor[i] = str
	}

	for i := 1; i <= n; i++ {
		for k := 0; k < 2; k++ {
			chminS(&dp_expr[i], dp_term[i])
			for j := 1; j < i; j++ {
				if dp_expr[j] == "" {
					continue
				}
				if dp_term[i-j] == "" {
					continue
				}
				chminS(&dp_expr[i], dp_expr[j]+"+"+dp_term[i-j])
			}
			chminS(&dp_term[i], dp_factor[i])
			for j := 1; j <= i; j++ {
				if i%j != 0 {
					continue
				}
				if dp_term[j] == "" {
					continue
				}
				if dp_factor[i/j] == "" {
					continue
				}
				chminS(&dp_term[i], dp_term[j]+"*"+dp_factor[i/j])
			}
			if dp_expr[i] != "" {
				chminS(&dp_factor[i], "("+dp_expr[i]+")")
			}
		}
	}
	out(dp_expr[n])
}
