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

type move struct {
	xs, ys, xe, ye int
}

func main() {
	sc.Split(bufio.ScanWords)
	sc.Buffer([]byte{}, 1000000)

	H, W := getInt(), getInt()
	a := make([][]int, H)
	odd := 0
	for i := 0; i < H; i++ {
		a[i] = make([]int, W)
		for j := 0; j < W; j++ {
			a[i][j] = getInt()
			if a[i][j]%2 == 1 {
				odd++
			}
		}
	}

	x := 0
	y := 0
	cnt := 0
	dx := 1
	m := make([]move, 0)
	for i := 0; i < H*W; i++ {
		if a[y][x]%2 == 1 {
			cnt++
		}
		nx := x
		ny := y
		nx += dx
		if nx == W {
			nx--
			dx = -1
			ny++
		}
		if nx < 0 {
			nx++
			dx = 1
			ny++
		}
		if cnt%2 == 1 {
			if cnt < odd {
				m = append(m, move{x, y, nx, ny})
			}
		}
		x = nx
		y = ny
	}

	w := bufio.NewWriter(os.Stdout)
	defer w.Flush()
	fmt.Fprintln(w, len(m))
	for _, v := range m {
		fmt.Fprintf(w, "%d %d %d %d\n", v.ys+1, v.xs+1, v.ye+1, v.xe+1)
	}
}
