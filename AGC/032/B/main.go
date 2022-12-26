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

func sel(a []int, u []bool, n int) []int {
	x := make([]int, 0)
	sum := 0
	idx := 0
	for sum != n/3 {
		if u[idx] != true && sum+a[idx] <= n/3 {
			sum += a[idx]
			x = append(x, a[idx])
			u[idx] = true
		}
		idx++
	}
	return x
}

func solve(N int) {
	g := make([][2]int, N/2)
	total := 0
	for i := 0; i < N/2; i++ {
		g[i][0] = i + 1
		g[i][1] = N - i
		total += i
	}
	out(total * 4)
	// out(g)
	len := N / 2
	for i := 0; i < len-1; i++ {
		for j := i + 1; j < len; j++ {
			out(g[i][0], g[j][0])
			out(g[i][0], g[j][1])
			out(g[i][1], g[j][0])
			out(g[i][1], g[j][1])
		}
	}
}

type ij struct {
	i, j int
}

func solveOdd(N int) {
	n := N/2 + 1
	g := make([][]int, n)
	a := make([]int, N)
	u := make([]bool, N)
	for i := 0; i < N; i++ {
		a[i] = N - i
	}
	// out(a)
	for i := 0; i < n; i++ {
		sum := 0
		idx := 0
		for sum != N {
			// out(a[idx], sum)
			if u[idx] != true && sum+a[idx] <= N {
				sum += a[idx]
				u[idx] = true
				g[i] = append(g[i], a[idx])
			}
			idx++
		}
	}

	cnt := 0
	ans := make([]ij, 0)
	for i := 0; i < n-1; i++ {
		for j := i + 1; j < n; j++ {
			for _, xx := range g[i] {
				for _, yy := range g[j] {
					ans = append(ans, ij{xx, yy})
					cnt++
				}
			}
		}
	}

	out(cnt)
	for _, v := range ans {
		out(v.i, v.j)
	}
}

func main() {
	sc.Split(bufio.ScanWords)
	N := getInt()

	if N == 3 {
		out(2)
		out(1, 3)
		out(2, 3)
		return
	}

	if N%2 == 0 {
		solve(N)
		return
	}

	solveOdd(N)

}
