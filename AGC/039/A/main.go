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

	S := []byte(getString())
	k := getInt()

	cnt := 1
	ans := 0
	for i := 1; i < len(S); i++ {
		if S[i] == S[i-1] {
			cnt++
		} else {
			ans += cnt / 2
			cnt = 1
		}
	}
	ans += cnt / 2

	if cnt == len(S) {
		out(len(S) * k / 2)
		return
	}

	c := S[0]
	left := 0
	for i := 0; i < len(S); i++ {
		if S[i] == c {
			left++
		} else {
			break
		}
	}
	right := 0
	for i := len(S) - 1; i >= 0; i-- {
		if S[i] == c {
			right++
		} else {
			break
		}
	}

	x := left/2 + right/2 - (left+right)/2

	// out(ans, x, left, right)
	out(ans*k - x*(k-1))
}
