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

	N := getInt()
	a := make([]int, N)
	b := make([]int, N)
	c := make([]int, N)
	for i := 0; i < N; i++ {
		a[i] = getInt()
	}
	for i := 0; i < N; i++ {
		b[i] = getInt()
	}
	for i := 0; i < N; i++ {
		c[i] = getInt()
	}
	sort.Ints(a)
	sort.Ints(b)
	sort.Ints(c)
	//	out(a)
	//	out(b)
	//	out(c)
	B := make([]int, N)
	pa := 0
	for i := 0; i < N; i++ {
		for pa < N && a[pa] < b[i] {
			//			out(pa, a[pa], b[i])
			pa++
		}
		B[i] = pa
	}
	C := make([]int, N)
	pb := 0
	cnt := 0
	ans := 0
	for i := 0; i < N; i++ {
		for pb < N && b[pb] < c[i] {
			//			out(pb, b[pb], c[i])
			cnt += B[pb]
			pb++
		}
		C[i] = cnt
		ans += cnt
	}
	//	out(B)
	//	out(C)
	out(ans)
}
