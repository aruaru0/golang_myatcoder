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

type pair struct {
	pos, cost int
}

func main() {
	defer wr.Flush()
	sc.Split(bufio.ScanWords)
	sc.Buffer([]byte{}, math.MaxInt32)
	// this template is new version.
	// use getI(), getS(), getInts(), getF()
	N, K := getI(), getI()
	p := getInts(N)
	c := getInts(N)

	ans := -int(1e18)

	for i := 0; i < N; i++ {
		MAX := -int(1e18)
		used := make([]bool, N)
		pos := i
		cost := 0
		rt := []int{i}
		tot := []int{0}
		for used[pos] == false {
			used[pos] = true
			pos = p[pos] - 1
			cost += c[pos]
			rt = append(rt, pos)
			tot = append(tot, cost)
		}

		if K < len(tot) { // Kが小さい場合は、そこで終了
			for j := 1; j <= K; j++ {
				MAX = max(MAX, tot[j])
			}
		} else if cost < 0 { // ループする必要なし
			for _, e := range tot {
				MAX = max(MAX, e)
			}
		} else {
			loop := K / (len(tot) - 1)
			if loop > 0 {
				loop--
				sum := loop * tot[len(tot)-1]
				for _, e := range tot {
					MAX = max(MAX, sum+e)
				}
			}
			loop = K / (len(tot) - 1)
			sum := loop * tot[len(tot)-1]
			rest := K % (len(tot) - 1)
			MAX = max(MAX, sum)
			for i := 1; i <= rest; i++ {
				MAX = max(MAX, sum+tot[i])
			}
		}
		ans = max(ans, MAX)
	}

	out(ans)
}
