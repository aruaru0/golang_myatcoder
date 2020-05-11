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

	H, W := getInt(), getInt()
	N := getInt()
	a := make([]int, N)
	for i := 0; i < N; i++ {
		a[i] = getInt()
	}

	dirc := true
	idx := 0
	cnt := 0
	for y := 0; y < H; y++ {
		b := make([]int, W)
		for x := 0; x < W; x++ {
			xx := x
			if !dirc {
				xx = W - 1 - x
			}
			b[xx] = idx + 1
			cnt++
			if a[idx] == cnt {
				cnt = 0
				idx++
			}
		}
		for x := 0; x < W; x++ {
			fmt.Print(b[x], " ")
		}
		out()
		if dirc {
			dirc = false
		} else {
			dirc = true
		}
	}

}
