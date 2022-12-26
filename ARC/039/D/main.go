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

func check(x, y, w, h int, s []string, c [][]int) bool {
	dx := []int{-1, 0, 1, -1, 0, 1, -1, 0, 1}
	dy := []int{-1, -1, -1, 0, 0, 0, 1, 1, 1}

	fill := true
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			xx := x + dx[j]
			yy := y + dy[j]
			if xx < 0 || xx >= w || yy < 0 || yy >= h {
				continue
			}
			if s[yy][xx] == '.' {
				fill = false
				break
			}
		}
	}
	if fill == true {
		c[y][x] = 1
	}

	return fill
}

func fill(x, y, w, h int, c, r [][]int) {
	dx := []int{-1, 0, 1, -1, 0, 1, -1, 0, 1}
	dy := []int{-1, -1, -1, 0, 0, 0, 1, 1, 1}

	if c[y][x] == 1 {
		for i := 0; i < 9; i++ {
			for j := 0; j < 9; j++ {
				xx := x + dx[j]
				yy := y + dy[j]
				if xx < 0 || xx >= w || yy < 0 || yy >= h {
					continue
				}
				r[yy][xx] = 1
			}
		}
	}
}

func main() {
	sc.Split(bufio.ScanWords)
	sc.Buffer([]byte{}, 1000000)

	H, W := getInt(), getInt()
	s := make([]string, H)
	c := make([][]int, H)
	r := make([][]int, H)
	for i := 0; i < H; i++ {
		s[i] = getString()
		c[i] = make([]int, W)
		r[i] = make([]int, W)
	}

	// 周辺が黒の点を選ぶ（元の推定）
	for y := 0; y < H; y++ {
		for x := 0; x < W; x++ {
			check(x, y, W, H, s, c)
		}
	}
	// 再現
	for y := 0; y < H; y++ {
		for x := 0; x < W; x++ {
			fill(x, y, W, H, c, r)
		}
	}

	// チェック
	ans := true
	for y := 0; y < H; y++ {
		for x := 0; x < W; x++ {
			if r[y][x] == 1 && s[y][x] != '#' {
				ans = false
				break
			}
			if r[y][x] == 0 && s[y][x] != '.' {
				ans = false
				break
			}
		}
	}

	if ans == false {
		out("impossible")
	} else {
		out("possible")
		for y := 0; y < H; y++ {
			for x := 0; x < W; x++ {
				if c[y][x] == 1 {
					fmt.Print("#")
				} else {
					fmt.Print(".")
				}
			}
			out()
		}
	}
}
