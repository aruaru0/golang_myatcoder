package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func out(x ...interface{}) {
	//	fmt.Println(x...)
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

func bitCount(bits int) int {

	bits = (bits & 0x55555555) + (bits >> 1 & 0x55555555)
	bits = (bits & 0x33333333) + (bits >> 2 & 0x33333333)
	bits = (bits & 0x0f0f0f0f) + (bits >> 4 & 0x0f0f0f0f)
	bits = (bits & 0x00ff00ff) + (bits >> 8 & 0x00ff00ff)
	return (bits & 0x0000ffff) + (bits >> 16 & 0x0000ffff)
}

type say struct {
	x, y int
}

type person struct {
	n int
	s []say
}

func check(tbl []int, N int, A []person) bool {
	for i := 0; i < N; i++ {
		for _, v := range A[i].s {
			if tbl[i] == 1 && tbl[v.x] != v.y {
				return false
			}
			// 問題をよく読む 0の人は真偽不明！！（嘘を言っているわけではない）
			//			if tbl[i] == 0 && tbl[v.x] == v.y {
			//				return false
			//			}
		}
	}
	return true
}

func main() {
	sc.Split(bufio.ScanWords)

	N := getInt()
	A := make([]person, N)
	for i := 0; i < N; i++ {
		n := getInt()
		A[i].s = make([]say, n)
		for j := 0; j < n; j++ {
			A[i].s[j] = say{getInt() - 1, getInt()}
		}
	}

	out(N)
	for i := 0; i < N; i++ {
		out(A[i])
	}

	ans := 0
	for i := 0; i < 1<<uint(N); i++ {
		tbl := make([]int, N)
		for j := 0; j < N; j++ {
			tbl[j] = (i >> uint(j)) % 2
		}
		res := check(tbl, N, A)
		n := bitCount(i)
		if res == true && ans < n {
			ans = n
		}
		out(tbl, res, ans)
	}
	fmt.Println(ans)
}
