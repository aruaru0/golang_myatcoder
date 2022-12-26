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

type pair struct {
	x, y int
}

func main() {
	sc.Split(bufio.ScanWords)
	sc.Buffer([]byte{}, 1000000)
	s := make([]string, 10)
	a := make([][]int, 10)
	for i := 0; i < 10; i++ {
		s[i] = getString()
		a[i] = make([]int, 10)
	}

	dx := []int{-1, 1, 0, 0}
	dy := []int{0, 0, -1, 1}
	var bsf = func(x, y, label int) {
		q := make([]pair, 0)
		q = append(q, pair{x, y})
		a[y][x] = label
		for len(q) != 0 {
			cx := q[0].x
			cy := q[0].y
			q = q[1:]
			for i := 0; i < 4; i++ {
				px := cx + dx[i]
				py := cy + dy[i]
				if px < 0 || px >= 10 || py < 0 || py >= 10 {
					continue
				}
				if a[py][px] != 0 {
					continue
				}
				if s[py][px] == 'o' {
					a[py][px] = label
					q = append(q, pair{px, py})
				}
			}
		}
	}

	label := 1
	for y := 0; y < 10; y++ {
		for x := 0; x < 10; x++ {
			if s[y][x] == 'o' && a[y][x] == 0 {
				bsf(x, y, label)
				label++
			}
		}
	}

	if label <= 2 {
		out("YES")
		return
	}

	if label > 5 {
		out("NO")
		return
	}

	// out(label)
	// for y := 0; y < 10; y++ {
	// 	out(a[y])
	// }

	for y := 0; y < 10; y++ {
		for x := 0; x < 10; x++ {
			if a[y][x] == 0 {
				flg := 0
				for i := 0; i < 4; i++ {
					px := x + dx[i]
					py := y + dy[i]
					if px < 0 || px >= 10 || py < 0 || py >= 10 {
						continue
					}
					if a[py][px] == 0 {
						continue
					}
					flg |= 1 << (a[py][px] - 1)
				}
				if label == 3 && flg == 0x03 {
					out("YES")
					return
				}
				if label == 4 && flg == 0x07 {
					out("YES")
					return
				}
				if label == 5 && flg == 0x0f {
					out("YES")
					return
				}
			}
		}
	}
	out("NO")
	return
}
