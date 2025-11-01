package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)
var wr = bufio.NewWriter(os.Stdout)

func out(x ...interface{}) {
	fmt.Fprintln(wr, x...)
}

func outSlice[T any](s []T) {
	if len(s) == 0 {
		return
	}
	for i := 0; i < len(s)-1; i++ {
		fmt.Fprint(wr, s[i], " ")
	}
	fmt.Fprintln(wr, s[len(s)-1])
}

func getI() int {
	sc.Scan()
	i, e := strconv.Atoi(sc.Text())
	if e != nil {
		panic(e)
	}
	return i
}

func getF() float64 {
	sc.Scan()
	i, e := strconv.ParseFloat(sc.Text(), 64)
	if e != nil {
		panic(e)
	}
	return i
}

func getInts(N int) []int {
	ret := make([]int, N)
	for i := 0; i < N; i++ {
		ret[i] = getI()
	}
	return ret
}

func getS() string {
	sc.Scan()
	return sc.Text()
}

func getStrings(N int) []string {
	ret := make([]string, N)
	for i := 0; i < N; i++ {
		ret[i] = getS()
	}
	return ret
}

// min, max, asub, absなど基本関数
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

// min for n entry
func nmin(a ...int) int {
	ret := a[0]
	for _, e := range a {
		ret = min(ret, e)
	}
	return ret
}

// max for n entry
func nmax(a ...int) int {
	ret := a[0]
	for _, e := range a {
		ret = max(ret, e)
	}
	return ret
}

func chmin(a *int, b int) bool {
	if *a < b {
		return false
	}
	*a = b
	return true
}

func chmax(a *int, b int) bool {
	if *a > b {
		return false
	}
	*a = b
	return true
}

func asub(a, b int) int {
	if a > b {
		return a - b
	}
	return b - a
}

func abs(a int) int {
	if a >= 0 {
		return a
	}
	return -a
}

func lowerBound(a []int, x int) int {
	idx := sort.Search(len(a), func(i int) bool {
		return a[i] >= x
	})
	return idx
}

func upperBound(a []int, x int) int {
	idx := sort.Search(len(a), func(i int) bool {
		return a[i] > x
	})
	return idx
}

// 値を圧縮した配列を返す
func compressArray(a []int) []int {
	m := make(map[int]int)
	for _, e := range a {
		m[e] = 1
	}
	b := make([]int, 0)
	for e := range m {
		b = append(b, e)
	}
	sort.Ints(b)
	for i, e := range b {
		m[e] = i
	}

	ret := make([]int, len(a))
	for i, e := range a {
		ret[i] = m[e]
	}
	return ret
}

type Node struct {
	idx, dist int
}
type Pair struct {
	cur, src int
}

func main() {
	defer wr.Flush()
	sc.Split(bufio.ScanWords)
	sc.Buffer([]byte{}, math.MaxInt32)
	// this template is new version.
	// use getI(), getS(), getInts(), getF()
	N, M := getI(), getI()
	node := make([][]int, N)
	for i := 0; i < M; i++ {
		u, v := getI()-1, getI()-1
		node[u] = append(node[u], v)
		node[v] = append(node[v], u)
	}
	s := getS()

	const inf = int(1e18)
	p1 := make([]Node, N)
	p2 := make([]Node, N)

	for i := 0; i < N; i++ {
		p1[i].idx = -1
		p1[i].dist = inf
		p2[i].idx = -1
		p2[i].dist = inf
	}

	// Sを始点としてキューに入れる（現在地、始点）
	q := []Pair{}
	for i, e := range s {
		if e == 'S' {
			p1[i].idx = i
			p1[i].dist = 0
			q = append(q, Pair{i, i})
		}
	}

	// BFS
	for len(q) != 0 {
		cur, src := q[0].cur, q[0].src
		q = q[1:]
		// 現在地までの距離
		var dist int
		if src == p1[cur].idx { // 始点が違う場合は、２番目の距離を採用
			dist = p1[cur].dist
		} else {
			dist = p2[cur].dist
		}
		for _, to := range node[cur] {
			nxt := dist + 1 // 次までは+1
			switch src {
			case p1[to].idx: // 始点と最短距離の地点が同じ
				if nxt < p1[to].dist { // 距離が短くなっていたら更新してキューに入れる
					p1[to].dist = nxt
					q = append(q, Pair{to, src})
				}
			case p2[to].idx: // 始点と２番目の距離の地点が同じ
				if nxt < p2[to].dist { // 距離が短くなっていたら更新してキューに入れる
					p2[to].dist = nxt
					q = append(q, Pair{to, src})
				}
			default: // 最短、２番目にsrcが入ってない場合
				if nxt < p1[to].dist { // 最短より短ければ、最短を更新して2位にスライド
					p2[to] = p1[to]
					p1[to] = Node{src, nxt}
					q = append(q, Pair{to, src})
				} else if nxt < p2[to].dist { // 2位より短ければ2位を入れ替え
					p2[to] = Node{src, nxt}
					q = append(q, Pair{to, src})
				}
			}
		}
	}

	for i, e := range s {
		if e == 'D' {
			out(p1[i].dist + p2[i].dist)
		}
	}

}
