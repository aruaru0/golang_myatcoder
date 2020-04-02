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

	N := getInt()
	a := make([]int, N)
	for i := 0; i < N; i++ {
		a[i] = getInt()
	}
	// out(a)
	ans := 0
	for i := 0; i < N-1; i++ {
		pair := a[i]
		if a[i+1] < a[i] {
			pair = a[i+1]
		}
		// out("pair", a[i], pair)
		if pair != 0 && (a[i]-pair)%2 == 1 {
			pair--
		}
		a[i] -= pair
		a[i+1] -= pair
		ans += pair
		x := a[i] / 2
		ans += x
		a[i] -= x * 2
		// out(pair)
	}
	ans += a[N-1] / 2
	//out(a)
	out(ans)
}
