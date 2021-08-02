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

const mod = int(1e9 + 7)

func main() {
	defer wr.Flush()
	sc.Split(bufio.ScanWords)
	sc.Buffer([]byte{}, math.MaxInt32)
	// this template is new version.
	// use getI(), getS(), getInts(), getF()

	L, R := getI(), getI()

	var dp [64][2][2]int

	nr, nl := 0, 0
	for i := 0; i < 63; i++ {
		if R&(1<<i) != 0 {
			nr = i
		}
		if L&(1<<i) != 0 {
			nl = i
		}
	}
	flagR := func(i int) bool {
		return R&(1<<i) != 0
	}
	flagL := func(i int) bool {
		return L&(1<<i) != 0
	}

	for i := nr; i >= 0; i-- {
		if nl <= i {
			x, y := 0, 0
			if i != nr {
				y = 1
			}
			if i != nl {
				x = 1
			}
			dp[i][y][x] += 1
		}
		if i == 0 {
			continue
		}
		for y := 0; y < 2; y++ {
			for x := 0; x < 2; x++ {
				for a := 0; a < 2; a++ {
					for b := 0; b < 2; b++ {
						if a < b {
							continue
						}
						ny := 1
						nx := 1
						if a == 1 && !flagR(i-1) && y == 0 {
							continue
						}
						if a == 1 && flagR(i-1) && y == 0 {
							ny = 0
						}
						if a == 0 && !flagR(i-1) && y == 0 {
							ny = 0
						}
						if b == 0 && flagL(i-1) && x == 0 {
							continue
						}
						if b == 1 && flagL(i-1) && x == 0 {
							nx = 0
						}
						if b == 0 && !flagL(i-1) && x == 0 {
							nx = 0
						}
						dp[i-1][ny][nx] += dp[i][y][x]
						dp[i-1][ny][nx] %= mod
					}
				}
			}
		}
	}
	ans := 0
	for i := 0; i < 2; i++ {
		for j := 0; j < 2; j++ {
			ans += dp[0][i][j]
			ans %= mod
		}
	}
	out(ans)
}
