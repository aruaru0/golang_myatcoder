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

type Edge struct {
	from, to, cost int
}

type Route struct {
	path []int
}

// ベルマンフォード法
func bellmanFord(N, S int, edges []Edge) ([]int, []Route, bool) {
	d := make([]int, N+1)
	for i := 0; i <= N; i++ {
		d[i] = math.MaxInt32
	}
	d[S] = 0
	r := make([]Route, N+1)
	r[S].path = []int{S}

	negativeLoop := false
	for i := 0; i <= N; i++ {
		for _, e := range edges {
			if d[e.to] > d[e.from]+e.cost {
				d[e.to] = d[e.from] + e.cost
				r[e.to].path = append(r[e.from].path, e.to)
				if i == N {
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
	edges := make([]Edge, M*2)
	for i := 0; i < M; i++ {
		f, t, c := getInt(), getInt(), getInt()
		edges[i*2] = Edge{f, t, c}
		edges[i*2+1] = Edge{t, f, c}
	}
	S, D := getInt(), getInt()

	// ここから
	d, r, negativeLoop := bellmanFord(N, S, edges)

	if negativeLoop == true {
		out("Negative Loop")
	} else {
		out("cost = ", d[D], "path=", r[D].path)
	}

}
