package main

import (
	"bufio"
	"fmt"
	"math/bits"
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

func getf() float64 {
	sc.Scan()
	f, _ := strconv.ParseFloat(sc.Text(), 64)
	return f
}

var probA [20][20]float64
var probB [20][20]float64

func main() {
	sc.Split(bufio.ScanWords)
	N := getInt()
	pa, pb := getf(), getf()
	a := getInts(N)
	b := getInts(N)
	sort.Ints(a)
	sort.Ints(b)

	dp := make([]float64, 1<<N)
	dp[0] = 1

	for i := 0; i < 1<<N; i++ {
		count := bits.OnesCount(uint(i))
		first := true
		firstP := 1.0
		if count != N-1 {
			firstP = pa
		}
		secondP := 1 - firstP
		if count != N-1 {
			secondP /= float64(N - 1 - count)
		}
		for j := 0; j < N; j++ {
			if (i>>j)%2 == 1 {
				continue
			}
			p := secondP
			if first {
				p = firstP
				first = false
			}
			dp[i+(1<<j)] += dp[i] * p
			probA[count][j] += dp[i] * p
		}
	}

	dp = make([]float64, 1<<N)
	dp[0] = 1

	for i := 0; i < 1<<N; i++ {
		count := bits.OnesCount(uint(i))
		first := true
		firstP := 1.0
		if count != N-1 {
			firstP = pb
		}
		secondP := 1 - firstP
		if count != N-1 {
			secondP /= float64(N - 1 - count)
		}
		for j := 0; j < N; j++ {
			if (i>>j)%2 == 1 {
				continue
			}
			p := secondP
			if first {
				p = firstP
				first = false
			}
			dp[i+(1<<j)] += dp[i] * p
			probB[count][j] += dp[i] * p
		}
	}

	ans := 0.0
	for i := 0; i < N; i++ {
		for j := 0; j < N; j++ {
			move := 0.0
			if a[i] > b[j] {
				move = float64(a[i] + b[j])
			}
			for k := 0; k < N; k++ {
				ans += move * probA[k][i] * probB[k][j]
			}
		}
	}
	out(ans)
}
