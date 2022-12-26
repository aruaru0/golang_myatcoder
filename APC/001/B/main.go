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

	N := getInt()
	a := make([]int, N)
	b := make([]int, N)
	sumA := 0

	for i := 0; i < N; i++ {
		a[i] = getInt()
		sumA += a[i]
	}
	sumB := 0
	for i := 0; i < N; i++ {
		b[i] = getInt()
		sumB += b[i]
	}
	totalOps := sumB - sumA
	aCnt := 0
	bCnt := 0
	for i := 0; i < N; i++ {
		if a[i] == b[i] {
			continue
		}
		if a[i] > b[i] {
			bCnt += a[i] - b[i]
		} else {
			d := b[i] - a[i]
			if d%2 == 1 {
				bCnt++
			}
			d = (d + 1) / 2
			aCnt += d
		}
		if aCnt > totalOps || bCnt > totalOps {
			out("No")
			return
		}
	}
	out("Yes")
}
