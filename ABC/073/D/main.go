package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
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

const inf = math.MaxInt32

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

// リストからidx番目の要素を抜かしたものを取得
func delete(idx int, L []int) (result []int) {
	result = append(result, L[:idx]...)
	result = append(result, L[idx+1:]...)
	return
}

// sliceの全組み合わせを返す
func permute(L []int) (result [][]int) {

	var inner func(LL []int)

	inner = func(LL []int) {
		if len(LL) == 0 {
			result = append(result, []int{}) //
		}
		for idx, r := range L {
			for _, t := range permute(delete(idx, L)) {
				result = append(result, append([]int{r}, t...))
			}
		}
	}
	inner(L)
	return

}

func main() {
	sc.Split(bufio.ScanWords)

	N, M, R := getInt(), getInt(), getInt()
	var D [201][201]int
	for i := 0; i <= N; i++ {
		for j := 0; j <= N; j++ {
			D[i][j] = inf
		}
		D[i][i] = 0
	}
	r := make([]int, R)
	for i := 0; i < R; i++ {
		r[i] = getInt()
	}
	for i := 0; i < M; i++ {
		from, to, cost := getInt(), getInt(), getInt()
		D[from][to] = cost
		D[to][from] = cost
	}

	for k := 1; k <= N; k++ {
		for i := 1; i <= N; i++ {
			for j := 1; j <= N; j++ {
				if D[i][k]+D[k][j] < D[i][j] {
					D[i][j] = D[i][k] + D[k][j]
				}
			}
		}
	}

	ans := inf
	for _, l := range permute(r) {
		sum := 0
		from := l[0]
		for i := 1; i < len(l); i++ {
			to := l[i]
			sum += D[from][to]
			from = to
		}
		ans = min(ans, sum)
	}
	out(ans)

}
