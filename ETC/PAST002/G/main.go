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

type str struct {
	c byte
	x int
}

func main() {
	sc.Split(bufio.ScanWords)
	Q := getInt()

	s := make([]str, 0)
	for i := 0; i < Q; i++ {
		T := getInt()
		if T == 1 {
			c, x := getString(), getInt()
			s = append(s, str{c[0], x})
		} else {
			d := getInt()
			a := make([]int, 26)
			for d > 0 {
				if len(s) == 0 {
					break
				}
				if s[0].x > d {
					s[0].x -= d
					a[int(s[0].c)-'a'] += d
					d = 0
				} else if s[0].x <= d {
					d -= s[0].x
					a[int(s[0].c)-'a'] += s[0].x
					s = s[1:]
				}
			}
			ans := 0
			for i := 0; i < 26; i++ {
				ans += a[i] * a[i]
			}
			out(ans)
		}
		// out(s)
	}
}
