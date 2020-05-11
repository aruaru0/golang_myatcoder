package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
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

const inf = math.MaxInt64 / 4

func main() {
	sc.Split(bufio.ScanWords)
	N, M, X := getInt(), getInt(), getInt()
	c := make([]int, N)
	a := make([][]int, N)
	for i := 0; i < N; i++ {
		c[i] = getInt()
		a[i] = make([]int, M)
		for j := 0; j < M; j++ {
			a[i][j] = getInt()
		}
	}

	// out(c)
	// out(a)
	n := 1 << uint(N)
	ans := inf
	for i := 0; i < n; i++ {
		b := make([]int, N)
		cnt := i
		idx := N - 1
		for cnt > 0 {
			b[idx] = cnt % 2
			cnt /= 2
			idx--
		}
		// out(b)

		cash := 0
		sum := make([]int, M)
		for j := 0; j < N; j++ {
			if b[j] == 1 {
				cash += c[j]
				for k := 0; k < M; k++ {
					// out(j, k)
					sum[k] += a[j][k]
				}
			}
		}
		flg := true
		for j := 0; j < M; j++ {
			if sum[j] < X {
				flg = false
				break
			}
		}
		if flg {
			ans = min(ans, cash)
		}
	}

	if ans == inf {
		out(-1)
		return
	}
	out(ans)

}
