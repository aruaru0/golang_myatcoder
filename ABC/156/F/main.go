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

func solve(n, x, m, k int, d []int) int {
	a := []int{x, 0}
	ret := 0
	for j := 1; j < n; j++ {
		idx := (j - 1) % k
		a[1] = a[0] + d[idx]
		if a[0]%m < a[1]%m {
			ret++
		}
		a[0] = a[1]
	}
	return ret
}

func solve2(n, x, m, k int, d []int) int {
	x0 := x
	xn1 := x
	eqn := 0
	for i := 0; i < k; i++ {
		num := (n - 1) / k
		if (n-1)%k > i {
			num++
		}
		xn1 += (d[i] % m) * num
		if d[i]%m == 0 {
			eqn += num
		}
	}
	ans := (n - 1) - (xn1/m - x0/m) - eqn

	return (ans)
}

func main() {
	sc.Split(bufio.ScanWords)

	k, q := getInt(), getInt()
	d := make([]int, k)
	for i := 0; i < k; i++ {
		d[i] = getInt()
	}
	n := make([]int, q)
	x := make([]int, q)
	m := make([]int, q)
	for i := 0; i < q; i++ {
		n[i], x[i], m[i] = getInt(), getInt(), getInt()
		out(solve2(n[i], x[i], m[i], k, d))
	}
}
