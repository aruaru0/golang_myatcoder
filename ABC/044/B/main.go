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

	w := getString()

	a := make([]int, 26)
	for _, v := range w {
		a[v-'a']++
	}

	ans := "Yes"
	for _, v := range a {
		if v%2 == 1 {
			ans = "No"
		}
	}
	fmt.Println(ans)

}
