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
	N := getInt()
	L := make([]float64, N)
	for i := 0; i < N; i++ {
		L[i] = float64(getInt())
	}
	K := getInt()
	l := float64(0.0)
	r := float64(10e9)
	prev := 1.0
	for i := 0; i < 100; i++ { // 無限ループ回避のため１００回以下しかループさせない
		cnt := 0
		m := (l + r) / 2
		for i := 0; i < N; i++ {
			cnt += int(L[i] / m)
		}
		if cnt < K {
			r = m
		} else {
			l = m
		}
		delta := math.Abs(prev - m)
		if delta < 1e-9 {
			break
		}
		prev = m
	}
	out(l)
}
