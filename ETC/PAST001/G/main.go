package main

import (
	"bufio"
	"fmt"
	"math"
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

type group struct {
	to []int
	c  []int
}

func calcCost(g []int, c [][]int) int {
	ret := 0
	//out(len(g), g)
	for i := 0; i < len(g); i++ {
		for j := 0; j < len(g); j++ {
			ret += c[g[i]][g[j]]
		}
	}
	//out(ret)
	return ret
}

func main() {
	sc.Split(bufio.ScanWords)
	sc.Buffer([]byte{}, 1000000)

	N := getInt()
	g := make([][]int, N)
	for i := 0; i < N; i++ {
		g[i] = make([]int, N)
	}
	for i := 0; i < N; i++ {
		for n := 0; n < N-i-1; n++ {
			a := getInt()
			f := i
			t := n + i + 1
			g[f][t] = a
			//g[t][f] = a
		}
	}

	n := 1
	for i := 0; i < N; i++ {
		n *= 3
	}

	ans := math.MinInt64
	for i := 0; i < n; i++ {
		gr := make([][]int, 3)
		for j := 0; j < 3; j++ {
			gr[j] = make([]int, 0)
		}
		x := i
		for j := 0; j < N; j++ {
			gr[x%3] = append(gr[x%3], j)
			x /= 3
		}
		//out(gr)
		cnt := 0
		for i := 0; i < 3; i++ {
			cnt += calcCost(gr[i], g)
		}
		ans = max(ans, cnt)
	}
	out(ans)
}
