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

const inf = math.MaxInt32

type link struct {
	from, to, cost int
}

func main() {
	sc.Split(bufio.ScanWords)
	N, M := getInt(), getInt()

	l := make([]link, M)
	for i := 0; i < M; i++ {
		l[i] = link{getInt() - 1, getInt() - 1, getInt()}
	}

	// 自分自身のコストを０に、それ以外を無限大に設定
	var dist [100][100]int
	for i := 0; i < N; i++ {
		for j := 0; j < N; j++ {
			if i == j {
				dist[i][j] = 0
			} else {
				dist[i][j] = inf
			}
		}
	}

	// 既知の経路を設定
	for i := 0; i < M; i++ {
		dist[l[i].from][l[i].to] = l[i].cost
		dist[l[i].to][l[i].from] = l[i].cost
	}

	// 経路を更新
	for x := 0; x < N; x++ {
		for i := 0; i < N; i++ {
			for j := 0; j < N; j++ {
				if dist[i][j] > dist[i][x]+dist[x][j] {
					// 経由したほうが近い場合、更新
					dist[i][j] = dist[i][x] + dist[x][j]
				}
			}
		}
	}

	ans := 0
	for i := 0; i < M; i++ {
		if dist[l[i].from][l[i].to] < l[i].cost {
			ans++
		}
	}
	out(ans)
}
