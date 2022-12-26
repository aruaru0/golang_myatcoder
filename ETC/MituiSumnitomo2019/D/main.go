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
	getInt()
	S := getString()
	// out(N, S)
	ans := 0
	for i := 0; i < 1000; i++ {
		a := make([]int, 3)
		idx := 2
		for n := i; n > 0; n /= 10 {
			a[idx] = n % 10
			idx--
		}
		idx = 0
		// out("----", idx)
		for _, v := range S {
			// out(v)
			if a[idx] == int(v-'0') {
				idx++
				// out(idx)
				if idx == 3 {
					ans++
					break
				}
			}
		}
	}
	out(ans)
}
