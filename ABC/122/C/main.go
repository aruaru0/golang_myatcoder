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
	sc.Buffer([]byte{}, 1000000)

	N, Q := getInt(), getInt()
	S := getString()
	S = S + "T"

	a := make([]int, N+1)
	b := make([]int, N+1)
	c := byte(' ')
	cnt := 0
	for i := 0; i <= N; i++ {
		if c == 'A' && S[i] == 'C' {
			cnt++
			b[i-1] = 1
			b[i] = 2
		}
		c = S[i]
		a[i] = cnt
	}
	// out(N, Q)
	// out(a)
	// out(b)

	for i := 0; i < Q; i++ {
		l, r := getInt()-1, getInt()
		n := a[r] - a[l]
		if b[r] == 2 {
			n--
		}
		// out(S[l:r])
		out(n)
	}

}
