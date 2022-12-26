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
	N, X, M := getI(), getI(), getI()

	used := make([]bool, M)
	a := []int{X}
	for used[a[len(a)-1]] != true {
		last := len(a) - 1
		used[a[last]] = true
		a = append(a, a[last]*a[last]%M)
	}

	if len(a) >= N {
		tot := 0
		for i := 0; i < N; i++ {
			tot += a[i]
		}
		out(tot)
		return
	}

	loopP := 0
	loopV := a[len(a)-1]
	preSum := 0
	a = a[:len(a)-1]
	for i := 0; i < len(a); i++ {
		if a[i] == loopV {
			loopP = i
			break
		}
		preSum += a[i]
	}

	loopSum := 0
	loopCnt := 0
	for i := loopP; i < len(a); i++ {
		loopSum += a[i]
		loopCnt++
	}

	ans := preSum
	N -= loopP
	n := N / loopCnt
	m := N % loopCnt
	ans += n * loopSum
	for i := 0; i < m; i++ {
		ans += a[loopP+i]
	}
	out(ans)
}
