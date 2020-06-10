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

func next(x, y, H, W int) (int, int) {
	if y%2 == 1 {
		x--
		if x < 0 {
			x = 0
			y++
		}
	} else {
		x++
		if x == W {
			x--
			y++
		}
	}
	return x, y
}

func main() {
	sc.Split(bufio.ScanWords)
	H, W := getInt(), getInt()
	a := make([][]int, H)
	for i := 0; i < H; i++ {
		a[i] = make([]int, W)
		for j := 0; j < W; j++ {
			a[i][j] = getInt()
		}
	}

	x := 0
	y := 0
	rest := 0
	cnt := 0
	b := make([][4]int, 0)
	for i := 0; i < H*W-1; i++ {
		a[y][x] += rest
		// out(y, x)
		if a[y][x]%2 == 0 {
			rest = 0
		} else {
			rest = 1
			cnt++
			a[y][x]--
			xx, yy := next(x, y, H, W)
			b = append(b, [4]int{y + 1, x + 1, yy + 1, xx + 1})
		}
		x, y = next(x, y, H, W)
	}
	if rest == 1 {
		a[H-1][W-1]++
	}

	out(cnt)
	w := bufio.NewWriter(os.Stdout)
	defer w.Flush()
	for _, v := range b {
		fmt.Fprintln(w, v[0], v[1], v[2], v[3])
	}

}
