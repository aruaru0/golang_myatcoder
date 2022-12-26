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

func check(y, x, h, w int, s []string) bool {
	dx := []int{0, -1, 1, 0}
	dy := []int{-1, 0, 0, 1}

	prev := 0
	next := 0
	for i := 0; i < 2; i++ {
		px := x + dx[i]
		py := y + dy[i]
		if px < 0 || py < 0 || px >= w || py >= h {
			continue
		}
		if s[y][x] == '#' && s[py][px] == '#' {
			prev++
		}
	}
	for i := 2; i < 4; i++ {
		nx := x + dx[i]
		ny := y + dy[i]
		if nx < 0 || ny < 0 || nx >= w || ny >= h {
			continue
		}
		if s[y][x] == '#' && s[ny][nx] == '#' {
			next++
		}
	}
	// fmt.Print(x, y, ":", prev, next, " ")

	if x == 0 && y == 0 {
		if next == 1 {
			return true
		} else {
			return false
		}
	}
	if x == w-1 && y == h-1 {
		if prev == 1 {
			return true
		} else {
			return false
		}
	}
	if prev != next {
		return false
	} else {
		return true
	}
}

func main() {
	sc.Split(bufio.ScanWords)
	sc.Buffer([]byte{}, 1000000)

	H, W := getInt(), getInt()
	a := make([]string, H)
	for i := 0; i < H; i++ {
		a[i] = getString()
	}
	// out(H, W, a)

	ans := "Possible"
	for y := 0; y < H; y++ {
		for x := 0; x < W; x++ {
			f := check(y, x, H, W, a)
			// out(f)
			if !f {
				ans = "Impossible"
				break
			}
		}
	}

	out(ans)
}
