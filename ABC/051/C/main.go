package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func out(x ...interface{}) {
	fmt.Print(x...)
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
	sx, sy, tx, ty := getInt(), getInt(), getInt(), getInt()

	// (sx, sy) to (tx, ty)
	for y := sy; y < ty; y++ {
		out("U")
	}
	for x := sx; x < tx; x++ {
		out("R")
	}
	// (tx, ty) to (sx, sy)
	for y := sy; y < ty; y++ {
		out("D")
	}
	for x := sx; x < tx; x++ {
		out("L")
	}

	// (sx, sy) to (tx, ty) 2
	out("L")
	for y := sy; y <= ty; y++ {
		out("U")
	}
	for x := sx; x <= tx; x++ {
		out("R")
	}
	out("D")

	// (tx, ty) to (sx, sy) 2
	out("R")
	for y := sy; y <= ty; y++ {
		out("D")
	}
	for x := sx; x <= tx; x++ {
		out("L")
	}
	out("U\n")

}
