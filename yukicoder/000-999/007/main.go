package main

import (
	"bufio"
	"fmt"
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

func getInts(N int) []int {
	ret := make([]int, N)
	for i := 0; i < N; i++ {
		ret[i] = getInt()
	}
	return ret
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

func pnum(n int) []int {
	d := make([]int, n+1)
	ret := make([]int, 0)
	for i := 2; i*i <= n; i++ {
		for j := i * 2; j <= n; j += i {
			d[j] = 1
		}
	}
	for i := 2; i <= n; i++ {
		if d[i] == 0 {
			ret = append(ret, i)
		}
	}
	return ret
}

func main() {
	sc.Split(bufio.ScanWords)
	N := getInt()
	p := pnum(N)

	dp := make([]bool, N+1)
	dp[0] = true
	dp[1] = true
	for i := 0; i <= N; i++ {
		for _, v := range p {
			if v+i > N {
				break
			}
			dp[i+v] = dp[i+v] || !dp[i]
		}
	}
	//out(dp)
	if dp[N] {
		out("Win")
		return
	}
	out("Lose")
}
