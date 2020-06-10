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

type pair struct {
	x, y int
}

func main() {
	sc.Split(bufio.ScanWords)
	sc.Buffer([]byte{}, 1000000)
	H, W, D := getInt(), getInt(), getInt()
	a := make([]pair, H*W)
	for i := 0; i < H; i++ {
		for j := 0; j < W; j++ {
			x := getInt() - 1
			a[x] = pair{i, j}
		}
	}

	N := H * W
	t := make([]int, N)
	for i := 0; i < D; i++ {
		for j := i + D; j < N; j += D {
			d := abs(a[j].x-a[j-D].x) + abs(a[j].y-a[j-D].y)
			t[j] = d + t[j-D]
		}
	}
	// out(t)

	Q := getInt()
	w := bufio.NewWriter(os.Stdout)
	defer w.Flush()
	for i := 0; i < Q; i++ {
		l, r := getInt()-1, getInt()-1
		fmt.Fprintln(w, t[r]-t[l])
	}
}
