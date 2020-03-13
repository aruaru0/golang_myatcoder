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

// 再帰で書くのが思いつかずループで書いてバグ多発
// 結局、写経（考えてたのとほぼ同じだがすきっりしてる
// といっても複雑）
func rec(N, X int, t, p []int) int {
	if N == 0 {
		if X == 1 {
			return 1
		}
		return 0
	}
	if X == 1 {
		return 0
	} else if X <= t[N-1]+1 {
		return rec(N-1, X-1, t, p)
	} else if X == t[N-1]+2 {
		return p[N-1] + 1
	} else if X <= t[N-1]*2+2 {
		return rec(N-1, X-(t[N-1]+2), t, p) + p[N-1] + 1
	}
	return p[N-1]*2 + 1
}

func main() {
	sc.Split(bufio.ScanWords)

	N, X := getInt(), getInt()

	p := make([]int, 51)
	t := make([]int, 51)

	p[0] = 1
	t[0] = 1

	for i := 1; i <= N; i++ {
		t[i] = t[i-1]*2 + 3
		p[i] = p[i-1]*2 + 1
	}

	ans := rec(N, X, t, p)
	out(ans)
}
