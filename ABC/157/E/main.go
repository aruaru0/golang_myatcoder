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

func readLine(r *bufio.Reader) []byte {
	buf := make([]byte, 0, 1024)
	for {
		l, p, e := r.ReadLine()
		if e != nil {
			panic(e)
		}
		buf = append(buf, l...)
		if !p {
			break
		}
	}
	return buf
}

//
// セグメント木
//
type segTree []int

const segTreeDefault = 0

// マージ関数
func (s *segTree) segMerge(a, b int) int {
	if a > b {
		return a
	}
	return b
}

/*
func (s *segTree) segMerge(a,b, int) int {
	return a+b
}
*/

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

func main() {
	sc.Split(bufio.ScanWords)
	sc.Buffer([]byte{}, 1000000)

	N := getInt()
	S := []byte(getString())

	seg := make([]segTree, 26)
	for i := 0; i < 26; i++ {
		seg[i] = segTreeInit(N)
	}

	for i, v := range S {
		val := v - 'a'
		seg[val].set(1, i)
	}

	Q := getInt()

	for i := 0; i < Q; i++ {
		t := getInt()
		if t == 1 {
			i := getInt() - 1
			c := getString()
			p := S[i]
			S[i] = c[0]
			seg[p-'a'].set(0, i)
			seg[c[0]-'a'].set(1, i)
		} else {
			l, r := getInt()-1, getInt()
			ans := 0
			for i := 0; i < 26; i++ {
				ans += seg[i].query(l, r)
			}
			out(ans)
		}
	}

}
