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

func pos(i, h, w int) (int, int) {
	x := i % w
	y := i / w
	if y%2 == 1 {
		x = w - 1 - x
	}

	return x, y
}

func main() {
	sc.Split(bufio.ScanWords)

	H, W := getInt(), getInt()
	N := getInt()
	a := make([]int, N)
	for i := 0; i < N; i++ {
		a[i] = getInt()
	}

	var mat [100][100]int
	n := a[0]
	c := 1
	for i := 0; i < H*W; i++ {
		px, py := pos(i, H, W)
		mat[px][py] = c
		n--
		if c != N && n == 0 {
			n = a[c]
			c++
		}
	}

	for y := 0; y < H; y++ {
		for x := 0; x < W; x++ {
			fmt.Printf("%d ", mat[x][y])
		}
		fmt.Println()
	}
}
