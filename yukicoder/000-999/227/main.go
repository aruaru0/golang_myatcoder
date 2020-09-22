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
	a := make(map[int]int)
	for i := 0; i < 5; i++ {
		a[getInt()]++
	}
	cnt3 := 0
	cnt2 := 0
	for _, e := range a {
		if e == 2 {
			cnt2++
		}
		if e == 3 {
			cnt3++
		}
	}
	if cnt3 == 1 && cnt2 == 1 {
		out("FULL HOUSE")
		return
	}
	if cnt3 == 1 {
		out("THREE CARD")
		return
	}
	if cnt2 == 2 {
		out("TWO PAIR")
		return
	}
	if cnt2 == 1 {
		out("ONE PAIR")
		return
	}
	out("NO HAND")
}
