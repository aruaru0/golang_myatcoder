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

func mex(n int) int {
	ret := 0
	switch n {
	case 0: // 000
		ret = 0
	case 1: // 001
		ret = 1
	case 2: // 010
		ret = 0
	case 3: // 011
		ret = 2
	case 4: // 100
		ret = 0
	case 5: // 101
		ret = 1
	case 6: // 110
		ret = 0
	case 7: // 111
		ret = 3
	}
	return ret
}

func main() {
	defer wr.Flush()
	sc.Split(bufio.ScanWords)
	sc.Buffer([]byte{}, math.MaxInt32)
	// this template is new version.
	// use getI(), getS(), getInts(), getF()
	N := getI()
	a := getInts(N)
	s := getS()

	// dp[i][j][bit] i番目まで見て、今MEXのj文字目で、012の使っているパターンがbitの場合のMEXの総和
	var dp [4][8]int
	ans := 0
	for i := 0; i < N; i++ {
		var tmp [4][8]int
		for j := 0; j < 4; j++ {
			for k := 0; k < 8; k++ {
				tmp[j][k] = dp[j][k]
			}
		}
		if s[i] == 'M' {
			bit := 1 << a[i]
			// out(i, bit)
			tmp[1][bit] += 1
		}
		if i != 0 && s[i] == 'E' {
			for j := 0; j < 8; j++ {
				bit := 1 << a[i]
				tmp[2][bit|j] += dp[1][j]
				// out(i, j, bit|j, dp[1][j])
			}
		}
		if i != 0 && s[i] == 'X' {
			// out(dp[2])
			for j := 0; j < 8; j++ {
				bit := j | 1<<a[i]
				sum := mex(bit)
				// out(bit, sum, dp[2][j])
				ans += sum * dp[2][j]
			}
		}
		// out(string(s[i]), a[i], i, tmp)
		dp = tmp
	}
	out(ans)
}
