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
	w := make([]int, N)
	for i := 0; i < N; i++ {
		w[i] = getInt()
	}

	now := make([]int, N)
	ans := 0
	for i := N - 1; i >= 0; i-- {
		ok := false
		for j := 0; j < ans; j++ {
			//			out(i, j, now, w)
			if now[j] <= w[i] {
				now[j] = w[i]
				ok = true
				break
			}
		}
		if ok == false {
			//			out("new line")
			now[ans] = w[i]
			ans++
		}
		//		out(now)
	}
	fmt.Println(ans)
}
