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
	H, W := getI(), getI()
	s := make([]string, H)
	for i := 0; i < H; i++ {
		s[i] = getS()
	}

	dx := make([][]int, H)
	for i := 0; i < H; i++ {
		dx[i] = make([]int, W)
	}
	dy := make([][]int, H)
	for i := 0; i < H; i++ {
		dy[i] = make([]int, W)
	}

	for i := 0; i < H; i++ {
		cnt := 0
		for j := 0; j < W; j++ {
			if s[i][j] == '#' {
				cnt = 0
			} else {
				cnt++
			}
			dx[i][j] = cnt
		}
		for j := W - 1; j >= 0; j-- {
			if s[i][j] == '#' {
				cnt = 0
			}
			cnt = max(dx[i][j], cnt)
			dx[i][j] = cnt
		}
	}

	for i := 0; i < W; i++ {
		cnt := 0
		for j := 0; j < H; j++ {
			if s[j][i] == '#' {
				cnt = 0
			} else {
				cnt++
			}
			dy[j][i] = cnt
		}
		for j := H - 1; j >= 0; j-- {
			if s[j][i] == '#' {
				cnt = 0
			}
			cnt = max(dy[j][i], cnt)
			dy[j][i] = cnt
		}
	}

	ans := 0
	for i := 0; i < H; i++ {
		for j := 0; j < W; j++ {
			ans = max(ans, dx[i][j]+dy[i][j]-1)
		}
	}
	out(ans)

}
