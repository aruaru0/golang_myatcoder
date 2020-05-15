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
	a := make([][]byte, H)
	for i := 0; i < H; i++ {
		a[i] = []byte(getString())
		// out(a[i])
	}
	h := make([]bool, H)
	for i := 0; i < H; i++ {
		flg := true
		for j := 0; j < W; j++ {
			if a[i][j] == '#' {
				flg = false
				break
			}
		}
		h[i] = flg
	}

	w := make([]bool, W)
	for j := 0; j < W; j++ {
		flg := true
		for i := 0; i < H; i++ {
			if a[i][j] == '#' {
				flg = false
				break
			}
		}
		w[j] = flg
	}

	// out(w)
	for i := 0; i < H; i++ {
		if h[i] == true {
			continue
		}
		for j := 0; j < W; j++ {
			if w[j] == true {
				continue
			}
			fmt.Print(string(a[i][j]))
		}
		out()
	}
}
