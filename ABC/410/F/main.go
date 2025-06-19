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

func calc(h0, w0, h1, w1 int, a [][]int) int {
	if h0 > h1 || w0 > w1 {
		return 0
	}
	return a[h1][w1] - a[h0][w1] - a[h1][w0] + a[h0][w0]
}

func solve() {
	h := getI()
	w := getI()
	s := make([][]byte, h)
	for i := 0; i < h; i++ {
		s[i] = []byte(getS())
	}

	if h > w {
		// 転置
		ns := make([][]byte, w)
		for i := 0; i < w; i++ {
			ns[i] = make([]byte, h)
			for j := 0; j < h; j++ {
				ns[i][j] = s[j][i]
			}
		}
		s = ns
		h, w = w, h
	}

	n := h * w
	cnt := make([]int, 2*n+1)

	ans := int64(0)
	for si := 0; si < h; si++ {
		a := make([]int, w)
		for ti := si; ti < h; ti++ {
			for j := 0; j < w; j++ {
				if s[ti][j] == '#' {
					a[j] += 1
				} else {
					a[j] += -1
				}
			}
			sum := make([]int, w+1)
			sum[0] = n
			for j := 0; j < w; j++ {
				sum[j+1] = sum[j] + a[j]
			}
			for j := 0; j < w; j++ {
				cnt[sum[j]]++
				ans += int64(cnt[sum[j+1]])
			}
			for j := 0; j < w; j++ {
				cnt[sum[j]] = 0
			}
		}
	}
	out(ans)
}

func main() {
	defer wr.Flush()
	sc.Split(bufio.ScanWords)
	sc.Buffer([]byte{}, math.MaxInt32)
	// this template is new version.
	// use getI(), getS(), getInts(), getF()
	t := getI()
	for i := 0; i < t; i++ {
		solve()
	}
}
