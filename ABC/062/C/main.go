package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

func out(x ...interface{}) {
	//	fmt.Println(x...)
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

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}

func solve(H, W int) int {
	res := math.MaxInt64 >> 4

	for i := 0; i < H; i++ {
		sel0 := i * W
		// 水平に３つ
		N := H - i
		n := N / 2
		sel1 := n * W
		sel2 := (N - n) * W
		d := max(sel0, max(sel1, sel2)) - min(sel0, min(sel1, sel2))
		if d < res {
			res = d
		}
		// 水平に１つ、垂直に２つ
		h := H - i
		w := W / 2
		sel1 = h * w
		sel2 = (W - w) * h
		d = max(sel0, max(sel1, sel2)) - min(sel0, min(sel1, sel2))
		if d < res {
			res = d
		}
	}
	return res
}

func main() {
	sc.Split(bufio.ScanWords)

	H, W := getInt(), getInt()

	res0 := solve(H, W)
	res1 := solve(W, H)
	fmt.Println(min(res0, res1))
}
