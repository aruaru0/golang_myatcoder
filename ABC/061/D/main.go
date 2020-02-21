package main

import (
	"bufio"
	"fmt"
	"math"
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

// Edge :
type Edge struct {
	from, to, cost int
}

// Route :
type Route struct {
	path []int
}

const inf = math.MaxInt64 >> 10

// ベルマンフォード法
func bellmanFord(N, S int, edges []Edge) ([]int, []Route, bool) {
	d := make([]int, N+1)
	for i := 0; i <= N; i++ {
		d[i] = inf
	}
	d[S] = 0
	r := make([]Route, N+1)
	r[S].path = []int{S}

	negativeLoop := false
	for i := 0; i <= N*2; i++ {
		for _, e := range edges {
			if d[e.to] > d[e.from]+e.cost {
				d[e.to] = d[e.from] + e.cost
				r[e.to].path = append(r[e.from].path, e.to)
				// i==Nを超えて、頂点Nに到達、つまり２回到達
				if i > N && (e.to == N || e.from == N) {
					negativeLoop = true
				}
			}
		}
	}

	return d, r, negativeLoop
}

func main() {
	sc.Split(bufio.ScanWords)

	N, M := getInt(), getInt()
	edges := make([]Edge, M)
	for i := 0; i < M; i++ {
		f, t, c := getInt(), getInt(), getInt()
		edges[i] = Edge{f, t, -c}
	}
	S, D := 1, N

	// ここから
	d, _, negativeLoop := bellmanFord(N, S, edges)

	if negativeLoop == true {
		out("inf")
	} else {
		out(-d[D])
	}

}
