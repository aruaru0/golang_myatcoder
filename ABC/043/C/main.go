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
	sum := 0
	for i := 0; i < N; i++ {
		a[i] = getInt()
		sum += a[i]
	}

	//	x := sum / N
	//	y := sum % N
	//	if y > N/2 {
	//		x++
	//	}

	x := 0
	if sum > 0 {
		x = int(float64(sum)/float64(N) + 0.5)
	} else {
		x = int(float64(sum)/float64(N) - 0.5)
	}

	ans := 0
	for i := 0; i < N; i++ {
		ans += (a[i] - x) * (a[i] - x)
	}
	fmt.Println(ans)
}
