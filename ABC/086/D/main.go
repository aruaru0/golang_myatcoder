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

func calc(x1, y1, x2, y2 int, a [][]int) int {
	ret := a[x2][y2] + a[x1][y1] - a[x1][y2] - a[x2][y1]
	// out(x1, y1, x2, y2, ret)
	return ret
}

func main() {
	sc.Split(bufio.ScanWords)
	N, K := getInt(), getInt()
	black := make([][]int, 2*K+1)
	white := make([][]int, 2*K+1)
	for i := 0; i <= 2*K; i++ {
		black[i] = make([]int, 2*K+1)
		white[i] = make([]int, 2*K+1)
	}
	for i := 0; i < N; i++ {
		x, y, c := getInt(), getInt(), getString()
		x %= 2 * K
		y %= 2 * K
		if c == "B" {
			black[x+1][y+1]++
		} else {
			white[x+1][y+1]++
		}
	}

	// out("Black")
	// for i := 0; i <= 2*K; i++ {
	// 	for j := 0; j <= 2*K; j++ {
	// 		fmt.Print(black[j][i], " ")
	// 	}
	// 	out()
	// }
	// out("white")
	// for i := 0; i <= 2*K; i++ {
	// 	for j := 0; j <= 2*K; j++ {
	// 		fmt.Print(white[j][i], " ")
	// 	}
	// 	out()
	// }
	for i := 0; i <= 2*K; i++ {
		for j := 1; j <= 2*K; j++ {
			black[i][j] += black[i][j-1]
			white[i][j] += white[i][j-1]
		}
	}
	for i := 0; i <= 2*K; i++ {
		for j := 1; j <= 2*K; j++ {
			black[j][i] += black[j-1][i]
			white[j][i] += white[j-1][i]
		}
	}

	ans := 0
	for i := 0; i < K; i++ {
		for j := 0; j < K; j++ {
			// out(i, j, "-------")
			b0 := 0
			b1 := 0
			w0 := 0
			w1 := 0
			// out("black")
			b0 += calc(0, 0, i, j, black)
			b1 += calc(i, 0, i+K, j, black)
			b0 += calc(i+K, 0, 2*K, j, black)

			b1 += calc(0, j, i, j+K, black)
			b0 += calc(i, j, i+K, j+K, black)
			b1 += calc(i+K, j, 2*K, j+K, black)

			b0 += calc(0, j+K, i, 2*K, black)
			b1 += calc(i, j+K, i+K, 2*K, black)
			b0 += calc(i+K, j+K, K*2, K*2, black)
			// out("white")
			w0 += calc(0, 0, i, j, white)
			w1 += calc(i, 0, i+K, j, white)
			w0 += calc(i+K, 0, 2*K, j, white)

			w1 += calc(0, j, i, j+K, white)
			w0 += calc(i, j, i+K, j+K, white)
			w1 += calc(i+K, j, 2*K, j+K, white)

			w0 += calc(0, j+K, i, 2*K, white)
			w1 += calc(i, j+K, i+K, 2*K, white)
			w0 += calc(i+K, j+K, K*2, K*2, white)
			ret := max(w0+b1, w1+b0)
			// out(ret)
			ans = max(ans, ret)
		}
	}
	out(ans)
}
