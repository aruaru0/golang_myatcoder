package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
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

func getInts(N int) []int {
	ret := make([]int, N)
	for i := 0; i < N; i++ {
		ret[i] = getInt()
	}
	return ret
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

func lowerBound(a []int, x int) int {
	idx := sort.Search(len(a), func(i int) bool {
		return a[i] >= x
	})
	return idx
}

func upperBound(a []int, x int) int {
	idx := sort.Search(len(a), func(i int) bool {
		return a[i] > x
	})
	return idx
}

const inf = int(1e15)

func solve(X, Y, Z int) {
	l := make([][2]int, 0)
	l = append(l, [2]int{1, 0})
	l = append(l, [2]int{0, 1})

	for i := 2; i < 100; i++ {
		x := [2]int{0, 0}
		for j := 1; j <= 2; j++ {
			for k := 0; k < 2; k++ {
				x[k] += l[i-j][k]
			}
		}
		l = append(l, x)
		if x[1] > int(1e9) {
			break
		}
	}

	// if X == Y && Y == Z {
	// 	A, B := inf, inf
	// 	for i := 1; i < len(l); i++ {
	// 		a := 1
	// 		u := l[i][0] * a
	// 		d := X - u
	// 		if d <= 0 {
	// 			break
	// 		}
	// 		// out(d, "a", a, l[i])
	// 		if d%l[i][1] == 0 {
	// 			b := d / l[i][1]
	// 			if a < A {
	// 				A, B = a, b
	// 			} else if a == A && b < B {
	// 				A, B = a, b
	// 			}
	// 			// out("SPECIAL", a, b)
	// 		}
	// 	}
	// 	out(A, B)
	// 	return
	// }

	if X == Y && Y == Z {
		X = 1
	} else if X == Y {
		Y, Z = Z, Y
	}

	A, B := inf, inf

	for i := 0; i < len(l); i++ {
		for j := 0; j < len(l); j++ {
			if i == j {
				continue
			}
			// out(i, j)
			u := X*l[j][0] - Y*l[i][0]
			d := l[i][1]*l[j][0] - l[j][1]*l[i][0]
			// out("a", u, d)
			if u%d != 0 {
				continue
			}
			b := u / d
			u = X*l[j][1] - Y*l[i][1]
			d = l[i][0]*l[j][1] - l[j][0]*l[i][1]
			// out("b", u, d)
			if u%d != 0 {
				continue
			}
			a := u / d
			// out(a, b)
			if a <= 0 || b <= 0 {
				continue
			}
			for k := 0; k < len(l); k++ {
				z := a*l[k][0] + b*l[k][1]
				if z == Z {
					if a < A {
						A, B = a, b
					} else if a == A && b < B {
						A, B = a, b
					}
				}
			}
		}
	}

	if A == inf {
		out(-1)
		return
	}
	out(A, B)
}

func main() {
	sc.Split(bufio.ScanWords)
	a := getInts(3)
	sort.Ints(a)

	solve(a[0], a[1], a[2])
}
