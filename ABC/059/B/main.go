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

	A, B := getString(), getString()

	a := len(A)
	b := len(B)

	ans := "EQUAL"
	if a > b {
		ans = "GREATER"
	} else if a < b {
		ans = "LESS"
	} else {
		for i := 0; i < a; i++ {
			if A[i] > B[i] {
				ans = "GREATER"
				break
			} else if A[i] < B[i] {
				ans = "LESS"
				break
			}
		}
	}

	out(ans)
}
