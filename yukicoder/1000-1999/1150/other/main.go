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
	N, s, t := getInt(), getInt()-1, getInt()-1
	a := make([]int, N)
	tot := 0
	for i := 0; i < N; i++ {
		a[i] = getInt()
		tot += a[i]
	}
	// ２連結して簡単化
	a = append(a, a...)
	// Ｓの右側を取り出す
	b := make([]int, 0, N)
	// Ｓの左側を取り出す
	c := make([]int, 0, N)
	if s < t {
		for i := s + N + 1; i < t+N; i++ {
			b = append(b, a[i])
		}
		for i := s + N - 1; i > t; i-- {
			c = append(c, a[i])
		}
	} else {
		for i := s + 1; i < t+N; i++ {
			b = append(b, a[i])
		}
		for i := s - 1; i > t; i-- {
			c = append(c, a[i])
		}
	}
	// 累積和に変換
	for i := 1; i < len(b); i++ {
		b[i] += b[i-1]
	}
	for i := 1; i < len(c); i++ {
		c[i] += c[i-1]
	}
	// 長さ
	n := (N - 2 + 1) / 2
	ans := -int(1e15)
	// 自力で行ける最大以上をカット
	b = b[:(len(b)+1)/2]
	c = c[:(len(c)+1)/2]
	// 簡易化のために０を挿入（合わせてｎを１足す）
	b = append([]int{0}, b...)
	c = append([]int{0}, c...)
	n++
	for i := 0; i < n; i++ {
		l := i
		r := n - 1 - i
		if l >= len(b) || r >= len(c) {
			continue
		}
		// 左右からＮ個取った時の最大を調べる
		v := b[l] + c[r] + a[s]
		ans = max(ans, v-(tot-v))
	}
	out(ans)
}
