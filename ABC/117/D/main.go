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
	N, K := getInt(), getInt()
	a := make([]int, N)
	for i := 0; i < N; i++ {
		a[i] = getInt()
		// out(strconv.FormatInt(int64(a[i]), 2))
	}

	c := make([]int, 40)
	for i := 0; i < N; i++ {
		n := a[i]
		idx := 0
		for n > 0 {
			c[idx] += n % 2
			n /= 2
			idx++
		}
	}

	cnt := 0
	k := K
	for k > 0 {
		cnt++
		k /= 2
	}
	// out(cnt)

	x := 0
	for i := cnt - 1; i >= 0; i-- {
		// out("i", i, c[i], N-c[i])
		if c[i] < N-c[i] {
			x += 1 << uint(i)
		}
		if x > K {
			x -= 1 << uint(i)
		}
	}
	// out(c)
	// out(x, strconv.FormatInt(int64(x), 2))
	ans := 0
	for i := 0; i < N; i++ {
		ans += a[i] ^ x
	}
	out(ans)
}
