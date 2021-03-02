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

func calc(s string, a int) int {
	x := make([]int, 10)
	x[a]++
	for i := 0; i < len(s)-1; i++ {
		x[int(s[i]-'0')]++
	}
	score := 0
	for i := 1; i < 10; i++ {
		c := 1
		for j := 0; j < x[i]; j++ {
			c *= 10
		}
		score += c * i
	}
	return score
}

func main() {
	defer wr.Flush()
	sc.Split(bufio.ScanWords)
	sc.Buffer([]byte{}, math.MaxInt32)
	// this template is new version.
	// use getI(), getS(), getInts(), getF()
	K, S, T := getI(), getS(), getS()

	num := make([]int, 10)
	for i := 0; i < 10; i++ {
		num[i] = K
	}
	for i := 0; i < len(S)-1; i++ {
		num[int(S[i]-'0')]--
	}
	for i := 0; i < len(T)-1; i++ {
		num[int(T[i]-'0')]--
	}

	tscore := make([]int, 10)
	ascore := make([]int, 10)
	for i := 1; i < 10; i++ {
		tscore[i] = calc(S, i)
		ascore[i] = calc(T, i)
	}
	// out(tscore)
	// out(ascore)
	// out(num)
	tot := float64(9*K - 8)
	cnt := 0.0
	for i := 1; i < 10; i++ {
		for j := 1; j < 10; j++ {
			if tscore[i] > ascore[j] {
				if i == j {
					if num[j] >= 2 {
						// out("win", i, j, num[i], num[j])
						cnt += float64(num[i]) / tot * float64(num[i]-1) / (tot - 1)
					}
				} else {
					if num[i] > 0 && num[j] > 0 {
						// out("win", i, j, num[i], num[j])
						cnt += float64(num[i]) / tot * float64(num[j]) / (tot - 1)
					}
				}
			}
		}
	}
	out(cnt)
}
