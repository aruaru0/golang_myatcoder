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

func main() {
	sc.Split(bufio.ScanWords)
	N, K := getInt(), getInt()
	a := make([]int, N+1)
	for i := 0; i < N; i++ {
		a[i] = getInt()
	}
	s := 0
	e := 0
	sum := 0
	ans := 0
	for {
		// out(s, e, sum, a)
		if sum >= K {
			// out("*", sum)
			sum -= a[s]
			s++
			ans += N + 1 - e
		} else {
			sum += a[e]
			e++
		}
		if e > N {
			break
		}
	}
	out(ans)
}
