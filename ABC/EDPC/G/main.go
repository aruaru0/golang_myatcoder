package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func out(x ...interface{}) {
	//	fmt.Println(x...)
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

// Node :
type Node struct {
	to []int
}

var dp [100100]int

func check(n int, node []Node) int {
	if dp[n] != -1 { // ここと、下の
		return dp[n]
	}
	res := 0
	for _, v := range node[n].to {
		res = max(res, check(v, node)+1)
	}
	dp[n] = res // ここがポイント。既に計算したものはその値を返す
	return res
}

func main() {
	sc.Split(bufio.ScanWords)

	N, M := getInt(), getInt()

	node := make([]Node, N)
	for i := 0; i < M; i++ {
		from, to := getInt()-1, getInt()-1
		node[from].to = append(node[from].to, to)
	}

	for i := 0; i < N; i++ {
		dp[i] = -1
	}

	// メモ化再帰を利用する方法
	res := 0
	for i := 0; i < N; i++ {
		res = max(res, check(i, node))
	}

	fmt.Println(res)
}
