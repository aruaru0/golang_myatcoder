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

	N, M := getInt(), getInt()
	a := make([]int, N)
	b := make([]int, M)
	for i := 0; i < N; i++ {
		a[i] = getInt()
	}
	for i := 0; i < M; i++ {
		b[i] = getInt()
	}
	sort.Ints(a)
	sort.Ints(b)

	n := 0
	cnt := 0
	for i := 0; i < M; i++ {
		for n < N {
			if b[i] <= a[n] {
				cnt++
				n++
				break
			}
			n++
		}
	}
	if cnt != M {
		out("NO")
	} else {
		out("YES")
	}
}
