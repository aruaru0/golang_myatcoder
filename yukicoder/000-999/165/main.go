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

func main() {
	sc.Split(bufio.ScanWords)
	N, B := getInt(), getInt()
	x := make([]int, N)
	y := make([]int, N)
	p := make([]int, N)
	mx := make(map[int]int)
	my := make(map[int]int)
	for i := 0; i < N; i++ {
		x[i], y[i], p[i] = getInt(), getInt(), getInt()
		mx[x[i]]++
		my[y[i]]++
	}
	// get xy
	px := make([]int, 0, len(mx))
	py := make([]int, 0, len(my))
	for xx := range mx {
		px = append(px, xx)
	}
	for yy := range my {
		py = append(py, yy)
	}
	sort.Ints(px)
	sort.Ints(py)
	mx = make(map[int]int)
	my = make(map[int]int)
	for i, xx := range px {
		mx[xx] = i + 1
	}
	for i, yy := range py {
		my[yy] = i + 1
	}

	// make cost table
	ny := len(my)
	nx := len(mx)
	c := make([][]int, ny+1)
	num := make([][]int, ny+1)
	for yy := 0; yy <= ny; yy++ {
		c[yy] = make([]int, nx+1)
		num[yy] = make([]int, nx+1)
	}
	for i := 0; i < N; i++ {
		xpos := mx[x[i]]
		ypos := my[y[i]]
		c[ypos][xpos] = p[i]
		num[ypos][xpos] = 1
	}
	// for yy := 0; yy <= ny; yy++ {
	// 	out(c[yy])
	// }
	// out("-------------")
	// for yy := 0; yy <= ny; yy++ {
	// 	out(num[yy])
	// }
	// out("-------------")

	for yy := 0; yy <= ny; yy++ {
		for xx := 1; xx <= nx; xx++ {
			c[yy][xx] += c[yy][xx-1]
			num[yy][xx] += num[yy][xx-1]
		}
	}
	for xx := 0; xx <= nx; xx++ {
		for yy := 1; yy <= ny; yy++ {
			c[yy][xx] += c[yy-1][xx]
			num[yy][xx] += num[yy-1][xx]
		}
	}
	// for yy := 0; yy <= ny; yy++ {
	// 	out(c[yy])
	// }
	// out("-------------")
	// for yy := 0; yy <= ny; yy++ {
	// 	out(num[yy])
	// }
	// 積分画像作成完了C
	cnt := 0
	for sy := 1; sy <= ny; sy++ {
		for ey := sy; ey <= ny; ey++ {
			sx, ex := 1, 0
			// out("Y", sy, ey, nx, ny)
			for sx <= nx && ex <= nx {
				cost := 0
				d := 0
				for cost <= B {
					ex++
					if ex > nx {
						break
					}
					cost = c[ey][ex] - c[sy-1][ex] - c[ey][sx-1] + c[sy-1][sx-1]
					d = num[ey][ex] - num[sy-1][ex] - num[ey][sx-1] + num[sy-1][sx-1]
					// out("EX", sx, ex, nx, cost, d)
					if cost <= B {
						cnt = max(cnt, d)
					}
				}
				for cost > B {
					sx++
					if ex > nx {
						break
					}
					// out("SX", sx, ex, nx)
					cost = c[ey][ex] - c[sy-1][ex] - c[ey][sx-1] + c[sy-1][sx-1]
					d = num[ey][ex] - num[sy-1][ex] - num[ey][sx-1] + num[sy-1][sx-1]
					if cost <= B {
						cnt = max(cnt, d)
					}
				}
				// out(sx, ex)
			}
		}
	}

	out(cnt)
}
