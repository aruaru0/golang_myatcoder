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

const inf = 1001001001001

func f(x, y int) {
	ans := inf
	for i := 0; i < 4; i++ {
		s := x
		e := y
		cnt := 0
		if i&1 == 1 {
			s = -s
			cnt++
		}
		if i&2 == 2 {
			e = -e
			cnt++
		}
		if s <= e {
			cnt += e - s
			ans = min(ans, cnt)
		}
	}

	out(ans)
}

func main() {
	sc.Split(bufio.ScanWords)
	x, y := getInt(), getInt()

	f(x, y)
	// out("----")

	// ax := abs(x)
	// ay := abs(y)
	// if x == y {
	// 	out(0)
	// 	return
	// }
	// if ax == ay {
	// 	out(1)
	// 	return
	// }

	// if x >= 0 && y >= 0 && x > y {
	// 	if y == 0 {
	// 		out(x - y + 1)
	// 	} else {
	// 		out(x - y + 2)
	// 	}
	// 	return
	// }
	// if x >= 0 && y >= 0 && x < y {
	// 	out(y - x)
	// 	return
	// }

	// if x < 0 && y >= 0 && ax > ay {
	// 	if y == 0 {
	// 		out(ax - ay)
	// 	} else {
	// 		out(ax - ay + 1)
	// 	}
	// 	return
	// }
	// if x < 0 && y >= 0 && ax < ay {
	// 	out(ay - ax + 1)
	// 	return
	// }

	// if x >= 0 && y < 0 && ax > ay {
	// 	out(ax - ay + 1)
	// 	return
	// }
	// if x >= 0 && y < 0 && ax < ay {
	// 	out(ay - ax + 1)
	// 	return
	// }

	// if x < 0 && y < 0 && ax > ay {
	// 	out(ax - ay)
	// 	return
	// }
	// if x < 0 && y < 0 && ax < ay {
	// 	out(ay - ax + 2)
	// 	return
	// }
}
