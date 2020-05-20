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
	h, w := getInt(), getInt()
	a := make([][]int, h)
	b := make([][]int, h)
	for i := 0; i < h; i++ {
		a[i] = make([]int, w)
		for j := 0; j < w; j++ {
			a[i][j] = getInt()
		}
	}
	for i := 0; i < h; i++ {
		b[i] = make([]int, w)
		for j := 0; j < w; j++ {
			b[i][j] = getInt()
		}
	}
	var dp [81][81][20000]bool

	dp[0][0][abs(a[0][0]-b[0][0])] = true
	for i := 0; i < h; i++ {
		for j := 0; j < w; j++ {
			for k := 0; k < 20000; k++ {
				if dp[i][j][k] == true {
					if j+1 < w {
						d := abs(a[i][j+1] - b[i][j+1])
						dp[i][j+1][abs(k-d)] = true
						dp[i][j+1][abs(k+d)] = true
					}
					if i+1 < h {
						d := abs(a[i+1][j] - b[i+1][j])
						dp[i+1][j][abs(k-d)] = true
						dp[i+1][j][abs(k+d)] = true
					}
				}
			}
		}
	}

	// out(dp[h-1][w-1][:10])
	//ans := math.MaxInt32
	for i := 0; i < 20000; i++ {
		if dp[h-1][w-1][i] == true {
			out(i)
			return
		}
	}
}
