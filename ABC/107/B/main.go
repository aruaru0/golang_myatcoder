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
	}

	hflg := make([]bool, H)
	for i := 0; i < H; i++ {
		del := true
		for j := 0; j < W; j++ {
			if a[i][j] != '.' {
				del = false
				break
			}
		}
		if del {
			hflg[i] = true
		}
	}

	wflg := make([]bool, W)
	for i := 0; i < W; i++ {
		del := true
		for j := 0; j < H; j++ {
			if a[j][i] != '.' {
				del = false
				break
			}
		}
		if del {
			wflg[i] = true
		}
	}

	for i := 0; i < H; i++ {
		if hflg[i] {
			continue
		}
		for j := 0; j < W; j++ {
			if wflg[j] != true {
				fmt.Print(string(a[i][j]))
			}
		}
		fmt.Println()
	}

}
