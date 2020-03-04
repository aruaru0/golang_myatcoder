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
	for i := 0; i < N; i++ {
		a[i] = getInt()
	}
	sort.Sort(sort.Reverse(sort.IntSlice(a)))

	sel := []int{0, 0}

	selected := 0
	cnt := 1
	for i := 1; i < len(a); i++ {
		if a[i] == a[i-1] {
			cnt++
		} else {
			cnt = 1
		}
		if cnt == 2 {
			sel[selected] = a[i-1]
			selected++
			cnt = 0
		}
		if selected >= 2 {
			break
		}
	}

	if selected == 2 {
		out(sel[0] * sel[1])
	} else {
		out(0)
	}
}
