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
	N, D, K := getInt(), getInt(), getInt()

	mi := K * (K + 1) / 2
	ma := N*(N+1)/2 - (N-K)*(N-K+1)/2

	if K > N {
		out(-1)
		return
	}
	if D < mi || ma < D {
		out(-1)
		return
	}

	res := make([]int, 0)
	for i := 0; i < K-1; i++ {
		res = append(res, i+1)
		D -= i + 1
	}
	res = append(res, D)

	for i := K - 1; i > 0; i-- {
		if res[i] > N {
			res[i-1] += res[i] - N
			res[i] = N
			N--
		} else {
			break
		}
	}
	for _, v := range res {
		fmt.Print(v, " ")
	}
	out()
}
