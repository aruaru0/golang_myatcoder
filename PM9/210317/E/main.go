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
	score, p, c int
}

func main() {
	defer wr.Flush()
	sc.Split(bufio.ScanWords)
	sc.Buffer([]byte{}, math.MaxInt32)
	// this template is new version.
	// use getI(), getS(), getInts(), getF()
	D, G := getI(), getI()
	x := make([]pair, D)
	for i := 0; i < D; i++ {
		x[i] = pair{(i + 1) * 100, getI(), getI()}
	}

	bit := 1 << D
	ans := int(1e18)
	for i := 0; i < bit; i++ {
		a := make([]int, D)
		tot := 0
		cnt := 0
		for j := 0; j < D; j++ {
			if (i>>j)%2 == 1 {
				tot += x[j].score*x[j].p + x[j].c
				cnt += x[j].p
			} else {
				a[j] = 1
			}
		}
		for j := D - 1; j >= 0; j-- {
			if a[j] == 0 {
				continue
			}
			if tot >= G {
				break
			}
			rest := G - tot
			n := min(x[j].p, (rest+x[j].score-1)/x[j].score)
			cnt += n
			tot += n * x[j].score
		}
		if tot >= G {
			ans = min(ans, cnt)
		}
	}
	out(ans)
}
