package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func out(x ...interface{}) {
	//	fmt.Println(x...)
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

func calcout(c []int, minEven, selEven, minOdd, selOdd int) {
	fmt.Print(c)
	fmt.Print("[")
	for i := 0; i < len(c); i++ {
		if i%2 == 0 {
			fmt.Print(c[i]-selOdd, " ")
		} else {
			fmt.Print(c[i]-selEven, " ")
		}
	}
	fmt.Println("]", "even", minEven, selEven, "odd", minOdd, selOdd)
}

const inf = 1001001001

func main() {
	sc.Split(bufio.ScanWords)
	sc.Buffer([]byte{}, 1000000)

	N := getInt()
	c := make([]int, N)
	minOdd := inf
	minEven := inf
	for i := 0; i < N; i++ {
		c[i] = getInt()
		if i%2 == 0 { // 奇数
			if c[i] < minOdd {
				minOdd = c[i]
			}
		} else {
			if c[i] < minEven {
				minEven = c[i]
			}
		}
	}

	Nodd := (N + 1) / 2

	ans := 0
	Q := getInt()
	selOdd := 0
	selEven := 0
	for i := 0; i < Q; i++ {
		n := getInt()
		switch n {
		case 1:
			x, a := getInt()-1, getInt()
			out("op1", x+1, a)
			if x%2 == 0 { // 奇数
				rest := c[x] - selOdd
				if rest >= a {
					if rest == minOdd {
						minOdd -= a
					}
					c[x] -= a
					ans += a
				}
			} else {
				rest := c[x] - selEven
				if rest >= a {
					if rest == minEven {
						minEven -= a
					}
					c[x] -= a
					ans += a
				}
			}
		case 2: // sell ODD
			a := getInt()
			out("op2", a)
			if minOdd >= a {
				minOdd -= a
				selOdd += a
				ans += a * Nodd
			}
		case 3:
			a := getInt()
			out("op3", a)
			if minOdd >= a && minEven >= a {
				minOdd -= a
				minEven -= a
				selOdd += a
				selEven += a
				ans += N * a
			}

		}
		//calcout(c, minEven, selEven, minOdd, selOdd)
		out(ans)
	}

	fmt.Println(ans)
}
