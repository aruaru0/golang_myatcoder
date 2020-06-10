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

	H, W := getInt(), getInt()

	// 縦
	w := W / 3
	m := W % 3
	if m == 2 {
		m--
	}
	ans := (w+m)*H - w*H
	// 横
	h := H / 3
	m = H % 3
	if m == 2 {
		m--
	}
	ans = min(ans, W*(h+m)-W*h)
	// 縦→横
	w = ((H * W) / 3) / H
	a := w * H
	b := (W - w) * (H / 2)
	c := (W - w) * (H/2 + H%2)
	s := max(a, max(b, c)) - min(a, min(b, c))
	ans = min(ans, s)
	w++
	a = w * H
	b = (W - w) * (H / 2)
	c = (W - w) * (H/2 + H%2)
	s = max(a, max(b, c)) - min(a, min(b, c))
	ans = min(ans, s)
	// 横→縦
	h = ((H * W) / 3) / W
	a = W * h
	b = (H - h) * (W / 2)
	c = (H - h) * (W/2 + W%2)
	s = max(a, max(b, c)) - min(a, min(b, c))
	ans = min(ans, s)
	h++
	a = W * h
	b = (H - h) * (W / 2)
	c = (H - h) * (W/2 + W%2)
	s = max(a, max(b, c)) - min(a, min(b, c))
	ans = min(ans, s)

	out(ans)
}
