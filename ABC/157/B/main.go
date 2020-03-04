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

	var a [3][3]int
	var c [3][3]int

	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			a[i][j] = getInt()
		}
	}

	N := getInt()
	for k := 0; k < N; k++ {
		b := getInt()
		for i := 0; i < 3; i++ {
			for j := 0; j < 3; j++ {
				if a[i][j] == b {
					c[i][j] = 1
				}
			}
		}
	}

	ans := "No"
	for i := 0; i < 3; i++ {
		if c[0][i]+c[1][i]+c[2][i] == 3 {
			ans = "Yes"
		}
	}
	for i := 0; i < 3; i++ {
		if c[i][0]+c[i][1]+c[i][2] == 3 {
			ans = "Yes"
		}
	}
	if c[0][0]+c[1][1]+c[2][2] == 3 {
		ans = "Yes"
	}
	if c[0][2]+c[1][1]+c[2][0] == 3 {
		ans = "Yes"
	}
	out(ans)
}
