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

const inf = math.MaxInt64

func main() {
	sc.Split(bufio.ScanWords)
	N, A, B, C := getInt(), getInt(), getInt(), getInt()
	l := make([]int, N)
	for i := 0; i < N; i++ {
		l[i] = getInt()
	}

	n := 1
	for i := 0; i < N; i++ {
		n *= 4
	}

	ans := inf

	for i := 0; i < n; i++ {
		a := [4]int{}
		mp := 0
		x := i
		cnt := [4]int{}
		for j := 0; j < N; j++ {
			p := x % 4
			x /= 4
			// fmt.Print(p)
			if a[p] != 0 && p != 0 {
				mp += 10
			}
			cnt[p]++
			a[p] += l[j]
		}
		if cnt[1] == 0 {
			continue
		}
		if cnt[2] == 0 {
			continue
		}
		if cnt[3] == 0 {
			continue
		}
		// out(a, A, B, C)
		mp += abs(a[1] - A)
		mp += abs(a[2] - B)
		mp += abs(a[3] - C)
		ans = min(ans, mp)
		// out(a, mp, "-")
	}
	out(ans)
}
