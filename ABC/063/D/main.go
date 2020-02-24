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

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	sc.Split(bufio.ScanWords)

	N, A, B := getInt(), getInt(), getInt()
	vmax := 0
	h := make([]int, N)
	for i := 0; i < N; i++ {
		h[i] = getInt()
		if vmax < h[i] {
			vmax = h[i]
		}
	}

	r := (vmax + B - 1) / B
	l := 0
	diff := A - B
	for l+1 < r {
		mid := (l + r) / 2
		dmg := B * mid
		ex := mid
		f := true
		for i := 0; i < N; i++ {
			if h[i] > dmg {
				d := h[i] - dmg
				ex -= (d + diff - 1) / diff
			}
			if ex < 0 {
				f = false
				break
			}
		}
		if f {
			r = mid
		} else {
			l = mid
		}
	}

	out(r)
}
