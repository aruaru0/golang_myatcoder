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

type person struct {
	dn []int
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func dfs(i int, p []person) int {
	if len(p[i].dn) == 0 {
		return 1
	}
	mi := 1001001001
	ma := -1
	for _, v := range p[i].dn {
		ret := dfs(v, p)
		mi = min(mi, ret)
		ma = max(ma, ret)
	}
	return mi + ma + 1
}

func main() {
	sc.Split(bufio.ScanWords)

	N := getInt()
	p := make([]person, N)
	for i := 1; i < N; i++ {
		b := getInt() - 1
		p[b].dn = append(p[b].dn, i)
	}
	out(dfs(0, p))
}
