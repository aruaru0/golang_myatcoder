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

const inf = 10000

func main() {
	sc.Split(bufio.ScanWords)
	sc.Buffer([]byte{}, 10000)

	H, W := getInt(), getInt()
	dist := make([][]int, H)
	a := make([]string, H)
	for i := 0; i < H; i++ {
		a[i] = getString()
		dist[i] = make([]int, W)
	}

	for y := 0; y < H; y++ {
		t := inf
		for x := 0; x < W; x++ {
			if a[y][x] == '#' {
				t = 0
			}
			dist[y][x] = t
			t++
		}
		t = dist[y][W-1]
		for x := W - 1; x >= 0; x-- {
			if a[y][x] == '#' {
				t = 0
			}
			dist[y][x] = min(dist[y][x], t)
			t++
		}
	}

	// for y := 0; y < H; y++ {
	// 	out(dist[y])
	// }
	// out("----")

	for x := 0; x < W; x++ {
		t := dist[0][x]
		for y := 0; y < H; y++ {
			if a[y][x] == '#' {
				t = 0
			}
			t = min(dist[y][x], t)
			dist[y][x] = t
			t++
		}
		t = dist[H-1][x]
		for y := H - 1; y >= 0; y-- {
			if a[y][x] == '#' {
				t = 0
			}
			t = min(dist[y][x], t)
			dist[y][x] = t
			t++
		}
	}

	// for y := 0; y < H; y++ {
	// 	out(dist[y])
	// }
	ans := 0
	for y := 0; y < H; y++ {
		for x := 0; x < W; x++ {
			ans = max(ans, dist[y][x])
		}
	}
	out(ans)
}
