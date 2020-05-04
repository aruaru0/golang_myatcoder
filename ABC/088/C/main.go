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
	c := [4][4]int{}
	for i := 1; i <= 3; i++ {
		for j := 1; j <= 3; j++ {
			c[i][j] = getInt()
		}
	}
	ans := true
	if c[1][1]-c[1][2] != c[2][1]-c[2][2] {
		ans = false
	}
	// out(ans)
	if c[1][1]-c[1][2] != c[3][1]-c[3][2] {
		ans = false
	}
	// out(ans)

	if c[1][2]-c[1][3] != c[2][2]-c[2][3] {
		ans = false
	}
	// out(ans)
	if c[1][2]-c[1][3] != c[3][2]-c[3][3] {
		ans = false
	}
	// out(ans)

	if c[1][1]-c[2][1] != c[1][2]-c[2][2] {
		ans = false
	}
	// out(ans)
	if c[1][1]-c[2][1] != c[1][3]-c[2][3] {
		ans = false
	}
	// out(ans)

	if c[2][1]-c[3][1] != c[2][2]-c[3][2] {
		ans = false
	}
	// out(ans)
	if c[2][1]-c[3][1] != c[2][3]-c[3][3] {
		ans = false
	}
	// out(ans)

	if ans == true {
		out("Yes")
	} else {
		out("No")
	}
}
