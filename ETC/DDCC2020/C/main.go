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

func main() {
	sc.Split(bufio.ScanWords)
	sc.Buffer([]byte{}, 1000000)

	H, W, _ := getInt(), getInt(), getInt()
	s := make([]string, H)
	for i := 0; i < H; i++ {
		s[i] = getString()
	}
	a := make([][]int, H)
	cnt := 1
	for i := 0; i < H; i++ {
		a[i] = make([]int, W)
		for j := 0; j < W; j++ {
			if s[i][j] == '#' {
				a[i][j] = cnt
				cnt++
			}
		}
	}
	/*
		for i := 0; i < H; i++ {
			out(a[i])
		}
	*/
	for y := 0; y < H; y++ {
		for x := 1; x < W; x++ {
			if a[y][x] == 0 {
				a[y][x] = a[y][x-1]
			}
		}
		for x := W - 2; x >= 0; x-- {
			if a[y][x] == 0 {
				a[y][x] = a[y][x+1]
			}
		}
	}
	/*
		out("----")
		for i := 0; i < H; i++ {
			out(a[i])
		}
	*/
	for y := 1; y < H; y++ {
		for x := 0; x < W; x++ {
			if a[y][x] == 0 {
				a[y][x] = a[y-1][x]
			}
		}
	}
	for y := H - 2; y >= 0; y-- {
		for x := 0; x < W; x++ {
			if a[y][x] == 0 {
				a[y][x] = a[y+1][x]
			}
		}
	}
	for i := 0; i < H; i++ {
		for j := 0; j < W; j++ {
			fmt.Print(a[i][j], " ")
		}
		out()
	}
}
