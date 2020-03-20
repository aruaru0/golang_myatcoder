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

//------------------------------------------------
// セグメントツリー
//------------------------------------------------
const inf = 1001001001

type rmq struct {
	seg []int
	n   int
}

func (s *rmq) init(n int) {
	N := 1
	for N < n {
		N *= 2
	}
	s.seg = make([]int, N*2)
	for i := 0; i < N*2; i++ {
		s.seg[i] = inf
	}
	s.n = N
}

func (s *rmq) update(k, x int) {
	k += s.n - 1
	s.seg[k] = x
	for k > 0 {
		k = (k - 1) / 2
		s.seg[k] = min(s.seg[k*2+1], s.seg[k*2+2])
	}
}

// a, b ... 範囲
func (s *rmq) querySub(a, b, k, l, r int) int {
	if r <= a || b <= l {
		return inf
	}
	if a <= l && r <= b {
		return s.seg[k]
	}
	return min(
		s.querySub(a, b, k*2+1, l, (l+r)/2),
		s.querySub(a, b, k*2+2, (l+r)/2, r))
}

// a, b ... 範囲 a <= x < bの範囲で検索
func (s *rmq) query(a, b int) int {
	return s.querySub(a, b, 0, 0, s.n)
}

//
// initして、範囲を設定して使う内部的には2^nでバッファ確保され
// infで埋められる。updateで値を設定して、使う
// minを書き換えれば、各範囲のいろいろなものを探索できる！！
//   max, a+bなど
//
func main() {
	sc.Split(bufio.ScanWords)
	sc.Buffer([]byte{}, 1000000)

	var a rmq
	a.init(8)
	a.update(0, 2)
	a.update(1, 1)
	a.update(2, 3)
	a.update(3, 10)
	a.update(4, 1)
	a.update(5, 10)
	a.update(6, 0)
	a.update(7, 10)
	out(a.seg)
	out(a.query(2, 5))
}
