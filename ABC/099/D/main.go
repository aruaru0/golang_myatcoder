package main

import (
	"bufio"
	"fmt"
	"os"
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

const inf = 1001001001001

func main() {
	sc.Split(bufio.ScanWords)
	N, C := getInt(), getInt()
	d := make([][]int, C)
	for i := 0; i < C; i++ {
		d[i] = make([]int, C)
		for j := 0; j < C; j++ {
			d[i][j] = getInt()
		}
	}
	c := make([][]int, N)
	for i := 0; i < N; i++ {
		c[i] = make([]int, N)
		for j := 0; j < N; j++ {
			c[i][j] = getInt() - 1
		}
	}

	var n [3][]int
	for i := 0; i < N; i++ {
		for j := 0; j < N; j++ {
			pos := ((i + 1) + (j + 1)) % 3
			n[pos] = append(n[pos], c[i][j])
		}
	}
	// out(n)

	var ans [3][30]int

	for i := 0; i < 3; i++ {
		for j := 0; j < C; j++ {
			cost := 0
			for _, v := range n[i] {
				cost += d[v][j]
			}
			ans[i][j] = cost
		}
	}
	// for i := 0; i < 3; i++ {
	// 	out(ans[i][:C])
	// }

	ret := inf
	for i := 0; i < C; i++ {
		for j := 0; j < C; j++ {
			if i == j {
				continue
			}
			for k := 0; k < C; k++ {
				if j == k || i == k {
					continue
				}
				x := ans[0][i] + ans[1][j] + ans[2][k]
				ret = min(ret, x)
			}

		}
	}
	out(ret)
}
