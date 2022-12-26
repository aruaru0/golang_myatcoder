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

	H, W := getInt(), getInt()
	a := make([]string, H)
	for i := 0; i < H; i++ {
		a[i] = getString()
	}
	// out(H, W, a)

	cnt := 0
	for y := 0; y < H; y++ {
		for x := 0; x < W; x++ {
			if a[y][x] == '#' {
				cnt++
			}
		}
	}

	if cnt == W+H-1 {
		out("Possible")
	} else {
		out("Impossible")
	}

}
