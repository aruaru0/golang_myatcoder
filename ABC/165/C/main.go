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

type item struct {
	a, b, c, d int
}

var ans int

func f(n, N, M int, a []int, nn []item, cc chan []int) {
	if len(a) == N {
		cc <- a
		return
	}

	for i := n; i <= M; i++ {
		f(i, N, M, append(a, i), nn, cc)
	}
}

func ff(n, N, M int, a []int, nn []item, cc chan []int) {
	f(n, N, M, a, nn, cc)
	close(cc)
}

func main() {
	sc.Split(bufio.ScanWords)
	N, M, Q := getInt(), getInt(), getInt()
	nn := make([]item, Q)
	for i := 0; i < Q; i++ {
		a, b, c, d := getInt(), getInt(), getInt(), getInt()
		nn[i] = item{a, b, c, d}
	}

	a := make([]int, 0)
	ch := make(chan []int)

	go ff(1, N, M, a, nn, ch)

	for v := range ch {
		cost := 0
		Q := len(nn)
		for j := 0; j < Q; j++ {
			f := nn[j].a - 1
			t := nn[j].b - 1
			if v[t]-v[f] == nn[j].c {
				cost += nn[j].d
			}
		}
		// out(cost)
		ans = max(ans, cost)
	}

	out(ans)
}
