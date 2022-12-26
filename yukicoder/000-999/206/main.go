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

func main() {
	sc.Split(bufio.ScanWords)
	L, M, N := getInt(), getInt(), getInt()
	n := N/64 + 1
	a := make([]uint, n)
	b := make([]uint, n)
	for i := 0; i < L; i++ {
		v := getInt()
		pos := v / 64
		bit := 63 - v%64
		a[pos] |= 1 << bit
	}
	for i := 0; i < M; i++ {
		v := getInt()
		pos := v / 64
		bit := 63 - v%64
		b[pos] |= 1 << bit
	}

	// for i := 0; i < n; i++ {
	// 	fmt.Printf("%64.64b ", uint(a[i]))
	// }
	// out()
	// for i := 0; i < n; i++ {
	// 	fmt.Printf("%64.64b ", uint(b[i]))
	// }
	// out()

	w := bufio.NewWriter(os.Stdout)
	defer w.Flush()

	Q := getInt()
	for i := 0; i < Q; i++ {
		cnt := 0
		pos := i / 64
		bit := i % 64
		pre := uint(0)

		// out(i, ":")
		for j, k := pos, 0; j < n; j++ {
			y := (b[k] >> bit) | pre
			// fmt.Printf("b %64.64b --> \n", uint(b[k])>>bit)
			// fmt.Printf("b %64.64b\n", uint(y))
			// fmt.Printf("a %64.64b\n", uint(a[j]))
			cnt += bits.OnesCount(uint(a[j] & y))
			pre = b[k] << (64 - bit)
			// fmt.Printf("-->%64.64b\n", uint(pre))
			k++
		}
		fmt.Fprintln(w, cnt)
	}
}
