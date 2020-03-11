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

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	sc.Split(bufio.ScanWords)
	sc.Buffer([]byte{}, 1000000)

	H, W := getInt(), getInt()
	X := make([][]int, H)
	Y := make([][]int, H)

	for y := 0; y < H; y++ {
		X[y] = make([]int, W)
		Y[y] = make([]int, W)
		s := getString()
		for x := 0; x < W; x++ {
			if s[x] == '#' {
				X[y][x] = 0
				Y[y][x] = 0
			} else {
				if x == 0 {
					X[y][x] = 1
				} else {
					X[y][x] = X[y][x-1] + 1
				}
				if y == 0 {
					Y[y][x] = 1
				} else {
					Y[y][x] = Y[y-1][x] + 1
				}
			}
		}
	}

	ymax := make([]int, W)
	for x := 0; x < W; x++ {
		ymax[x] = -1
	}
	for y := H - 1; y >= 0; y-- {
		xmax := -1
		for x := W - 1; x >= 0; x-- {
			if X[y][x] == 0 || xmax < X[y][x] {
				xmax = X[y][x]
			} else {
				X[y][x] = xmax
			}
			if Y[y][x] == 0 || ymax[x] < Y[y][x] {
				ymax[x] = Y[y][x]
			} else {
				Y[y][x] = ymax[x]
			}
		}
	}

	ans := 0
	for y := 0; y < H; y++ {
		for x := 0; x < W; x++ {
			ans = max(ans, Y[y][x]+X[y][x])
		}
	}
	out(ans - 1)
}
