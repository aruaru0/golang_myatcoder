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

	N, K := getInt(), getInt()

	ans := 0
	for a := 1; a <= N; a++ {
		kbc := K - a%K
		if (kbc+kbc)%K != 0 {
			continue
		}
		nb := (N-kbc)/K + 1
		nc := (N-kbc)/K + 1
		ans += nb * nc
	}
	out(ans)
}
