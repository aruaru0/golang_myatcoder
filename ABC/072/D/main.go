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
	p := make([]int, N+1)
	for i := 1; i <= N; i++ {
		p[i] = getInt()
	}

	ans := 0
	for i := 1; i <= N; i++ {
		if i == p[i] {
			if i == 1 {
				p[i], p[i+1] = p[i+1], p[i]
				ans++
			} else if i == N {
				p[i], p[i-1] = p[i-1], p[i]
				ans++
			} else {
				if p[i+1] == i+1 {
					p[i], p[i+1] = p[i+1], p[i]
				} else {
					p[i], p[i-1] = p[i-1], p[i]
				}
				ans++
			}

		}
	}
	out(ans)

}
