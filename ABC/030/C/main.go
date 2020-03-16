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

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func asub(a, b int) int {
	if a > b {
		return a - b
	}
	return b - a
}

func find(t int, a []int) int {
	if len(a) == 1 {
		if a[0] < t {
			return -1
		}
		return a[0]
	}

	L := 0
	R := len(a) - 1

	for L+1 != R {
		mid := (L + R) / 2
		if a[mid] <= t {
			L = mid
		} else {
			R = mid
		}
	}

	ans := a[L]
	if ans < t {
		ans = a[R]
	}
	if ans < t {
		ans = -1
	}

	return ans
}

func main() {
	sc.Split(bufio.ScanWords)
	sc.Buffer([]byte{}, 1000000)

	N, M := getInt(), getInt()
	X, Y := getInt(), getInt()
	a := make([]int, N)
	for i := 0; i < N; i++ {
		a[i] = getInt()
	}
	b := make([]int, M)
	for i := 0; i < M; i++ {
		b[i] = getInt()
	}

	T := 0
	ans := 0
	for {
		f := find(T, a)
		if f == -1 {
			break
		}
		//		out("A->B", f, f+X)
		T = f + X
		t := find(T, b)
		if t == -1 {
			break
		}
		//		out("B->A", t, t+Y)
		T = t + Y
		ans++
	}

	out(ans)
}
