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

	H, W, A, B := getInt(), getInt(), getInt(), getInt()

	if H < B || W < A {
		out("No")
		return
	}

	w := bufio.NewWriter(os.Stdout)
	defer w.Flush()

	//	out(H, W, A, B)
	flgY := false
	for y := 0; y < H; y++ {
		if y == B {
			flgY = !flgY
		}
		flgX := false
		for x := 0; x < W; x++ {
			if x == A {
				flgX = !flgX
			}
			if flgX != flgY {
				fmt.Fprintf(w, "1")
			} else {
				fmt.Fprintf(w, "0")
			}
		}
		fmt.Fprintln(w)
	}

}
