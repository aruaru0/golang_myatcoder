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

func nextInt() int {
	sc.Scan()
	i, e := strconv.Atoi(sc.Text())
	if e != nil {
		panic(e)
	}
	return i
}

func nextString() string {
	sc.Scan()
	return sc.Text()
}

func main() {
	sc.Split(bufio.ScanWords)
	a := nextInt()
	b := nextInt()
	c := nextInt()

	a5 := 0
	a7 := 0
	if a == 5 {
		a5++
	} else if a == 7 {
		a7++
	}

	if b == 5 {
		a5++
	} else if b == 7 {
		a7++
	}

	if c == 5 {
		a5++
	} else if c == 7 {
		a7++
	}

	if (a5 == 2) && (a7 == 1) {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}

}
