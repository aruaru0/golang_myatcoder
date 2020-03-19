package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
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

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func asub(a, b int) int {
	if a > b {
		return a - b
	}
	return b - a
}

func main() {
	sc.Split(bufio.ScanWords)
	sc.Buffer([]byte{}, 1000000)

	N, H := getInt(), getInt()
	a := make([]int, N)
	b := make([]int, N)
	for i := 0; i < N; i++ {
		a[i], b[i] = getInt(), getInt()
	}
	sort.Ints(a)
	sort.Ints(b)
	maxA := a[len(a)-1]
	ans := (H + maxA - 1) / maxA
	sum := 0
	pos := len(b) - 1
	cnt := 0
	for i := pos; i >= 0; i-- {
		sum += b[i]
		cnt++
		rest := max(0, H-sum)
		ans = min(ans, cnt+(rest+maxA-1)/maxA)
	}

	out(ans)
}
