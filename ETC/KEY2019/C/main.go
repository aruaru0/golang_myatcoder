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

func main() {
	sc.Split(bufio.ScanWords)
	sc.Buffer([]byte{}, 1000000)

	N := getInt()
	a := make([]int, N)
	b := make([]int, N)
	for i := 0; i < N; i++ {
		a[i] = getInt()
	}
	for i := 0; i < N; i++ {
		b[i] = getInt()
	}

	sum := 0
	d := make([]int, N)
	for i := 0; i < N; i++ {
		d[i] = a[i] - b[i]
		sum += d[i]
	}

	// out(d, sum)

	if sum < 0 {
		out(-1)
		return
	}

	sort.Ints(d)
	sum0 := 0
	pos0 := 0
	sum1 := 0
	pos1 := len(d) - 1
	for i := 0; i < N; i++ {
		if d[i] < 0 {
			sum0 -= d[i]
			pos0++
		}
	}
	cnt := 0
	for i := 0; i < N; i++ {
		if sum1 >= sum0 {
			break
		}
		sum1 += d[pos1]
		pos1--
		cnt++
	}

	out(cnt + pos0)
}
