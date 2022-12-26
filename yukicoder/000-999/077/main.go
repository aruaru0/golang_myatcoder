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

func calc(p, sz int, n [100][300]int) int {
	cnt := 0
	y := 0
	// out("----", p, sz)
	for sz > 0 {
		for x := p; x < p+sz; x++ {
			cnt += n[y][x]
		}
		sz -= 2
		p++
		y++
	}
	// out("cnt", cnt)
	return cnt
}

func main() {
	sc.Split(bufio.ScanWords)
	N := getInt()
	a := getInts(N)
	var n [100][300]int
	tot := 0
	for x := 0; x < N; x++ {
		for y := 0; y < a[x]; y++ {
			n[y][x] = 1
		}
		tot += a[x]
	}

	sz := 1
	sum := sz
	ans := tot
	for sum <= tot {
		// out("-----", sum)
		for pos := 0; pos <= N; pos++ {
			ret := calc(pos, sz, n)
			// out(pos, sz, ret, tot)
			ans = min(ans, tot-ret)
		}
		sz += 2
		sum += sz
	}
	out(ans)
}
