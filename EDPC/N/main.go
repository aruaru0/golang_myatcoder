package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

func out(x ...interface{}) {
	fmt.Println(x...)
}

var sc = bufio.NewScanner(os.Stdin)

func getInt() int {
	sc.Scan()
	i, e := strconv.Atoi(sc.Text())
	if e != nil {
		panic(e)
	}
	return i
}

func getString() string {
	sc.Scan()
	return sc.Text()
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

var dp [401][401]int

func solve(f, t int, a []int) int {
	if f == t {
		return 0
	}
	if dp[f][t] != 0 {
		return dp[f][t]
	}
	s := 0
	ret := math.MaxInt64
	for i := f; i < t; i++ {
		r0 := solve(f, i, a)
		r1 := solve(i+1, t, a)
		ret = min(ret, r0+r1)
		s += a[i]
	}
	ret += s + a[t]
	dp[f][t] = ret
	return ret
}

func main() {
	sc.Split(bufio.ScanWords)

	N := getInt()
	a := make([]int, N+1)
	for i := 0; i < N; i++ {
		a[i] = getInt()
	}

	ret := solve(0, N-1, a)
	out(ret)
}
