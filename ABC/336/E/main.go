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

func reverse(s []int) []int {
	t := []int{}
	for i := len(s) - 1; i >= 0; i-- {
		t = append(t, s[i])
	}
	return t
}

func main() {
	defer wr.Flush()
	sc.Split(bufio.ScanWords)
	sc.Buffer([]byte{}, math.MaxInt32)
	// this template is new version.
	// use getI(), getS(), getInts(), getF()
	n := getI() + 1
	digit := make([]int, 0)
	for n > 0 {
		digit = append(digit, n%10)
		n /= 10
	}
	digit = reverse(digit)
	m := len(digit)

	var dp [16][2][130][130]int
	ans := 0
	for k := 1; k <= 126; k++ {
		for i := 0; i < m+1; i++ {
			for j := 0; j < 2; j++ {
				for s := 0; s < k+1; s++ {
					for r := 0; r < k; r++ {
						dp[i][j][s][r] = 0
					}
				}
			}
		}
		dp[0][0][0][0] = 1
		for i := 0; i < m; i++ {
			for j := 0; j < 2; j++ {
				for s := 0; s < k+1; s++ {
					for r := 0; r < k; r++ {
						for d := 0; d < 10; d++ {
							ni := i + 1
							nj := j
							ns := s + d
							nr := (r*10 + d) % k
							if ns > k {
								continue
							}
							if j == 0 {
								if digit[i] < d {
									continue
								}
								if digit[i] > d {
									nj = 1
								}
							}
							dp[ni][nj][ns][nr] += dp[i][j][s][r]
						}
					}
				}
			}
		}
		ans += dp[m][1][k][0]
	}
	out(ans)
}
