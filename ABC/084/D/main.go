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

//
// セグメント木
//
type segTree []int

const segTreeDefault = 0

// マージ関数
func (s *segTree) segMerge(a, b int) int {
	return a + b
}

func segTreeInit(N int) segTree {
	size := 1
	for N > size {
		size <<= 1
	}
	s := make(segTree, size*2)
	for i := 0; i < size*2; i++ {
		s[i] = segTreeDefault
	}
	s[0] = 0
	return s
}

func (s *segTree) set(i, pos int) {
	N := len(*s)/2 + pos
	(*s)[N] = i
	for N > 1 {
		N /= 2
		(*s)[N] = s.segMerge((*s)[2*N], (*s)[2*N+1])
	}
}

func (s *segTree) setAll(a []int, n int) {
	N := len(*s) / 2
	for i := 0; i < n; i++ {
		(*s)[N+i] = a[i]
	}
	for N > 1 {
		N >>= 1
		for i := 0; i < N; i++ {
			pos := N + i
			(*s)[pos] = s.segMerge((*s)[2*pos], (*s)[2*pos+1])
		}
	}
}

func (s *segTree) querySub(l, r, index, L, R int) int {
	if R <= l || r <= L {
		return segTreeDefault
	}
	if l <= L && R <= r {
		return (*s)[index]
	}
	med := (L + R) / 2
	lval := s.querySub(l, r, index*2, L, med)
	rval := s.querySub(l, r, index*2+1, med, R)
	return s.segMerge(lval, rval)
}

// queryの区間は[l, r)です
func (s *segTree) query(l, r int) int {
	return (*s).querySub(l, r, 1, 0, len(*s)/2)
}

/*
	seg := segTreeInit(N)

	seg.setAll(a, N)
	seg.set(val, pos)
	seg.query(l, r) // [L, R)
*/

var seg segTree

func initPrime(N int) {
	n := make([]int, N+1)
	seg = segTreeInit(N + 1)

	// 素数を求めて
	n[1] = 1
	for i := 2; i*i <= N; i++ {
		for j := i * 2; j <= N; j += i {
			n[j] = i
		}
	}

	// 素数の部分だけセグメント木に１を設定
	// セグメント木のマージ条件は数のカウント
	for i := 1; i <= N; i++ {
		if n[i] == 0 {
			p := (i + 1) / 2
			if n[p] == 0 {
				seg.set(1, i)
			}
		}
	}
}

func main() {
	sc.Split(bufio.ScanWords)
	w := bufio.NewWriter(os.Stdout)
	defer w.Flush()

	// セグメント木を作って
	initPrime(100100)

	// あとは、問い合わせるだけ
	Q := getInt()
	for i := 0; i < Q; i++ {
		l, r := getInt(), getInt()
		fmt.Fprintln(w, seg.query(l, r+1))
	}
}
