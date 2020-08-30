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

func main() {
	sc.Split(bufio.ScanWords)
	sc.Buffer([]byte{}, 1000000)
	N, L := getInt(), getInt()*60
	S := make([]int, N)
	maxS := 0
	sum := 0
	for i := 0; i < N; i++ {
		s := getString()
		var mm, ss int
		fmt.Sscanf(s, "%d:%d", &mm, &ss)
		S[i] = mm*60 + ss
		maxS = max(maxS, S[i])
		sum += S[i]
	}
	if sum <= L {
		out(N)
		return
	}
	dp := make([][]int, N+1)
	for i := 0; i <= N; i++ {
		dp[i] = make([]int, L+maxS)
	}
	dp[0][0] = 1
	for i := 0; i < N; i++ {
		for j := N - 1; j >= 0; j-- {
			for k := L - 1; k >= 0; k-- {
				dp[j+1][k+S[i]] += dp[j][k]
			}
		}
	}
	acc := 0.0
	fact := make([]float64, N+1)
	fact[0] = 1
	for i := 1; i <= N; i++ {
		fact[i] = float64(i) * fact[i-1]
	}
	prv := make([][]int, N+1)
	for i := 0; i <= N; i++ {
		prv[i] = make([]int, L+maxS)
	}
	for i := 0; i < N; i++ {
		for j := 0; j <= N; j++ {
			for k := 0; k <= L; k++ {
				prv[j][k] = dp[j][k]
			}
		}
		// prv := dp
		for j := 0; j < N; j++ {
			for k := 0; k < L; k++ {
				prv[j+1][k+S[i]] -= prv[j][k]
			}
		}
		for j := 0; j < N; j++ {
			for k := max(0, L-S[i]); k < L; k++ {
				acc += float64(j+1) * fact[j] * float64(prv[j][k]) * fact[N-j-1]
			}
		}
	}
	out(acc / fact[N])
}
