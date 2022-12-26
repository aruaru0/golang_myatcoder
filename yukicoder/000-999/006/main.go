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

func prime(K, N int) []int {
	n := make([]int, N+1)
	n[1] = 1
	for i := 2; i*i <= N; i++ {
		if n[i] != 0 {
			continue
		}
		for j := i * 2; j <= N; j += i {
			n[j] = 1
		}
	}
	ret := make([]int, 0)
	for i := K; i <= N; i++ {
		if n[i] == 0 {
			ret = append(ret, i)
		}
	}
	return ret
}

const inf = 1001001001001

func main() {
	sc.Split(bufio.ScanWords)
	K, N := getInt(), getInt()
	x := prime(K, N)
	a := make([]int, 0)
	for _, v := range x {
		if v < K {
			continue
		}
		n := v
		for n >= 10 {
			x := 0
			for n > 0 {
				x += n % 10
				n /= 10
			}
			n = x
		}
		a = append(a, n)
	}

	// 尺取り法で実装
	memo := make([]int, 10)
	cnt := 0
	r := 0
	ma := 0
	num := 0
	for l := 0; l < len(a); l++ {
		for r < len(a) && memo[a[r]] == 0 {
			memo[a[r]] = 1
			r++
		}
		cnt = r - l
		if cnt >= ma {
			ma = cnt
			num = x[l]
		}
		memo[a[l]] = 0
	}
	out(num)

	// 単純ループでも間に合ったぽい。以下単純ループ
	// ma := 0
	// ans := 0
	// for i := 0; i < len(a); i++ {
	// 	m := make([]int, 10)
	// 	cnt := 0
	// 	for j := i; j < len(a); j++ {
	// 		if m[a[j]] != 0 {
	// 			break
	// 		}
	// 		m[a[j]] = 1
	// 		cnt++
	// 	}
	// 	if cnt >= ma {
	// 		ma = cnt
	// 		ans = x[i]
	// 	}
	// }
	// out(ans)
}
