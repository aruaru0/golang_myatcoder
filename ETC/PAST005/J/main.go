package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"unicode"
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

func solve(S string, X int) {
	a := make([]int, len(S)+1)
	idx := 0
	for i, v := range S {
		if !unicode.IsNumber(v) {
			a[i+1] = a[i] + 1
		} else {
			n := int(v - '0')
			a[i+1] = a[i] * (n + 1)
		}
		if a[i+1] > int(1e17) {
			break
		}
		idx++
	}

	x := X
	for i := idx; i > 0; i-- {
		// out(x, a[i], i)
		if unicode.IsNumber(rune(S[i-1])) {
			n := int(S[i-1]-'0') + 1
			x = x % (a[i] / n)
			if x == 0 {
				x += a[i] / n
			}
		} else {
			if x == a[i] {
				fmt.Print(string(S[i-1]))
				break
			}
		}
	}
}

func main() {
	defer wr.Flush()
	sc.Split(bufio.ScanWords)
	sc.Buffer([]byte{}, 1000000)
	// this template is new version.
	// use getI(), getS(), getInts(), getF()
	S := getS()
	X := getI()

	solve(S, X)

	// for i := 1; i <= X; i++ {
	// 	solve(S, i)
	// }
	fmt.Println()
}
