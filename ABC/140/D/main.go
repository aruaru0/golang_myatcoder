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

func main() {
	sc.Split(bufio.ScanWords)
	sc.Buffer([]byte{}, 1000000)

	N, K := getInt(), getInt()
	s := getString()

	lr := 0
	rl := 0
	rr := 0
	ll := 0
	l := 0
	r := 0
	for i := 1; i < N; i++ {
		if s[i] == 'R' {
			r++
		}
		if s[i] == 'L' {
			l++
		}
		if s[i-1] == 'L' && s[i] == 'R' {
			lr++
		}
		if s[i-1] == 'R' && s[i] == 'L' {
			rl++
		}
		if s[i-1] == 'R' && s[i] == 'R' {
			rr++
		}
		if s[i-1] == 'L' && s[i] == 'L' {
			ll++
		}
	}
	ans := min(K*2, lr+rl)
	ans += ll + rr
	ans = min(N-1, ans)
	out(ans)
}
