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
	a := make(map[int]int)
	for i := 0; i < H; i++ {
		s := getString()
		for j := 0; j < W; j++ {
			a[int(s[j])]++
		}
	}
	// out(a)
	c1 := 0
	c2 := 0
	c4 := 0
	for _, v := range a {
		c4 += v / 4
		c2 += (v % 4) / 2
		c1 += (v % 4) % 2
	}
	// out(c1, c2, c4)

	if W%2 == 1 && H%2 == 1 {
		if c1 == 1 && c2 <= (W+H-2)/2 {
			out("Yes")
		} else {
			out("No")
		}
	} else if W%2 == 1 {
		if c1 == 0 && c2 <= H/2 {
			out("Yes")
		} else {
			out("No")
		}
	} else if H%2 == 1 {
		if c1 == 0 && c2 <= W/2 {
			out("Yes")
		} else {
			out("No")
		}
	} else {
		if c2 != 0 || c1 != 0 {
			out("No")
		} else {
			out("Yes")
		}
	}
}
