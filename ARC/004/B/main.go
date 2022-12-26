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
	d := make([]int, N)
	max := -1
	sum := 0
	for i := 0; i < N; i++ {
		d[i] = getInt()
		if max < d[i] {
			max = d[i]
		}
		sum += d[i]
	}

	out(sum)
	if max*2 > sum {
		out(max - (sum - max))
	} else {
		out(0)
	}

}
