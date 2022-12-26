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

type node struct {
	to []int
}

var c []int
var dp [2][100001]int

const mod = 1000000007

func dfs(col, next, prev int, n []node) int {
	if dp[col][next] != -1 {
		return dp[col][next]
	}

	r0 := 1
	r1 := 1
	for _, v := range n[next].to {
		if c[v] != 0 {
			r1 = 0
			break
		}
	}

	for _, v := range n[next].to {
		if v == prev {
			continue
		}
		c[next] = 1
		r1 *= dfs(1, v, next, n)
		r1 %= mod
		c[next] = 0
		r0 *= dfs(0, v, next, n)
		r0 %= mod
	}

	r := (r0 + r1) % mod
	dp[col][next] = r
	//	out(col, next, prev, r0+r1)

	return r
}

func main() {
	sc.Split(bufio.ScanWords)

	N := getInt()
	n := make([]node, N)
	c = make([]int, N)

	for i := 0; i < 100001; i++ {
		dp[0][i] = -1
		dp[1][i] = -1
	}
	for i := 0; i < N-1; i++ {
		from, to := getInt()-1, getInt()-1
		n[from].to = append(n[from].to, to)
		n[to].to = append(n[to].to, from)
	}
	out(dfs(0, 0, -1, n))
}
