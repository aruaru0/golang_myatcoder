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

type pos struct {
	x, y, h int
}

func check(cx, cy int, p []pos) (int, bool) {
	h := p[0].h + abs(p[0].x-cx) + abs(p[0].y-cy)
	// out(cx, cy, h, p[0])
	for _, v := range p {
		x := max(h-abs(v.x-cx)-abs(v.y-cy), 0)
		if x != v.h {
			return 0, false
		}
	}
	return h, true
}

func main() {
	sc.Split(bufio.ScanWords)

	N := getInt()
	p := make([]pos, N+1)
	for i := 1; i <= N; i++ {
		p[i].x, p[i].y, p[i].h = getInt(), getInt(), getInt()
		if p[i].h > 0 {
			p[0] = p[i]
		}
	}

	for y := 0; y <= 100; y++ {
		for x := 0; x <= 100; x++ {
			h, ok := check(x, y, p)
			if ok {
				out(x, y, h)
				return
			}
		}
	}
}
