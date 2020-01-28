package main

import (
	"bufio"
	"fmt"
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

type Data struct {
	x, y, a int
}

func main() {
	sc.Split(bufio.ScanWords)

	W := getInt()
	H := getInt()
	N := getInt()
	d := make([]Data, N)
	for i := 0; i < N; i++ {
		d[i].x = getInt()
		d[i].y = getInt()
		d[i].a = getInt()
	}

	var r [100][100]int
	out(W, H, N, d)

	for _, v := range d {
		switch v.a {
		case 1: // x < xi
			out("1:", v.x)
			for y := 0; y < H; y++ {
				for x := 0; x < v.x; x++ {
					out(x, y)
					r[y][x] = 1
				}
			}
		case 2: // x > xi
			out("2:")
			for y := 0; y < H; y++ {
				for x := v.x; x < W; x++ {
					out(x, y)
					r[y][x] = 1
				}
			}
		case 3: // y < yi
			out("3:")
			for y := 0; y < v.y; y++ {
				for x := 0; x < W; x++ {
					r[y][x] = 1
				}
			}
		case 4: // y > yi
			out("4:")
			for y := v.y; y < H; y++ {
				for x := 0; x < W; x++ {
					r[y][x] = 1
				}
			}
		}
	}

	ans := 0

	for y := 0; y < H; y++ {
		for x := 0; x < W; x++ {
			if r[y][x] == 0 {
				ans++
			}
			//			fmt.Print(r[y][x])
		}
		//		fmt.Println("")
	}
	fmt.Println(ans)
}
