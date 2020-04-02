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

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

var col []int

func dfs(u, c int, n [][]int) bool {
	col[u] = c
	for i, v := range n[u] {
		if v == 0 || v == inf {
			continue
		}
		var d int
		if c == 1 {
			d = 2
		} else {
			d = 1
		}
		if col[i] != 0 {
			if col[i] != d {
				return false
			}
			continue
		}
		if !dfs(i, d, n) {
			return false
		}
	}
	return true
}

const inf = 1001001001

func main() {
	sc.Split(bufio.ScanWords)
	sc.Buffer([]byte{}, 1000000)

	N := getInt()
	n := make([][]int, N)
	for i := 0; i < N; i++ {
		n[i] = make([]int, N)
		for j := 0; j < N; j++ {
			if i == j {
				continue
			}
			n[i][j] = inf
		}
	}
	for i := 0; i < N; i++ {
		s := getString()
		for j, v := range s {
			if v == '1' {
				n[i][j] = 1
			}
		}
	}

	col = make([]int, N)
	check := dfs(0, 1, n)

	if !check {
		out(-1)
		return
	}

	for k := 0; k < N; k++ {
		for i := 0; i < N; i++ {
			for j := 0; j < N; j++ {
				if n[i][j] > n[i][k]+n[k][j] {
					n[i][j] = n[i][k] + n[k][j]
				}
			}
		}
	}

	ans := 0
	for i := 0; i < N; i++ {
		for j := 0; j < N; j++ {
			ans = max(ans, n[i][j]+1)
		}
	}
	out(ans)
}
