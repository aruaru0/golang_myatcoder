package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)

func solveComentary(n, m, q int, t, a, b, c []int) (ans []int) {
	latest := make([][2]int, n)
	for i := 0; i < n; i++ {
		latest[i] = [2]int{-1, 0}
	}
	offset := make([][]int, q)
	aid := make([]int, q)
	for i := range t {
		switch t[i] {
		case 2:
			latest[a[i]] = [2]int{i, b[i]}
		case 3:
			idx := len(ans)
			aid[i] = idx
			j, x := latest[a[i]][0], latest[a[i]][1]
			ans = append(ans, x)
			if j >= 0 {
				offset[j] = append(offset[j], i)
			}
		}
	}
	fen := NewFenwickTree(m)
	for i := range t {
		switch t[i] {
		case 1:
			fen.Add(a[i], c[i])
			fen.Add(b[i], -c[i])
		case 2:
			//sq := query.(SubstitutionQuery)
			for _, v := range offset[i] {
				ans[aid[v]] -= fen.Sum(0, b[v])
			}
		case 3:
			ans[aid[i]] += fen.Sum(0, b[i])
		}
	}
	return ans
}

func main() {
	buf := make([]byte, 1024*1024)
	sc.Buffer(buf, bufio.MaxScanTokenSize)
	sc.Split(bufio.ScanWords)

	n, m, q := nextInt(), nextInt(), nextInt()
	var t []int
	a, b, c := make([]int, q), make([]int, q), make([]int, q)
	for i := 0; i < q; i++ {
		t = append(t, nextInt())
		switch t[i] {
		case 1:
			a[i] = nextInt()
			b[i] = nextInt()
			c[i] = nextInt()
			// 0-indexed
			a[i]--
		case 2:
			a[i] = nextInt()
			b[i] = nextInt()
			// 0-indexed
			a[i]--
		case 3:
			a[i] = nextInt()
			b[i] = nextInt()
			// 0-indexed
			a[i]--
		}
	}

	ans := solveComentary(n, m, q, t, a, b, c)

	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()
	for _, v := range ans {
		fmt.Fprintln(out, v)
	}
}

func nextInt() int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return i
}

func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func Min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func Max(x, y int) int {
	if x < y {
		return y
	}
	return x
}

func Floor(x, y int) int {
	return x / y
}

func Ceil(x, y int) int {
	return (x + y - 1) / y
}

type FenwickTree struct {
	n     int
	nodes []int
}

func NewFenwickTree(n int) *FenwickTree {
	fen := new(FenwickTree)
	fen.n = n + 1
	fen.nodes = make([]int, fen.n)
	return fen
}

func (fen *FenwickTree) Add(i, v int) {
	i++
	for i <= fen.n {
		fen.nodes[i-1] += v
		i += i & -i
	}
}

func (fen *FenwickTree) Sum(l, r int) int {
	return fen.sum(r) - fen.sum(l)
}

func (fen *FenwickTree) sum(i int) int {
	res := 0
	for i > 0 {
		res += fen.nodes[i-1]
		i -= i & -i
	}
	return res
}
