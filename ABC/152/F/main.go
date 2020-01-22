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

// Link :
type Link struct {
	id int
	to int
}

// Node :
type Node struct {
	link []Link
}

var node []Node
var route []int = make([]int, 0)

// DFS :
func DFS(from, to, prev int) bool {
	if from == to {
		return true
	}
	for _, v := range node[from].link {
		if v.to == prev {
			continue
		}
		if DFS(v.to, to, from) {
			route = append(route, v.id)
			return true
		}
	}
	return false
}

func main() {
	// read datas
	sc.Split(bufio.ScanWords)
	N := getInt()
	node = make([]Node, N)
	for i := 0; i < N-1; i++ {
		a := getInt() - 1
		b := getInt() - 1
		node[a].link = append(node[a].link, Link{i, b})
		node[b].link = append(node[b].link, Link{i, a})
	}

	M := getInt()
	u := make([]int, M)
	v := make([]int, M)
	for i := 0; i < M; i++ {
		u[i] = getInt() - 1
		v[i] = getInt() - 1
	}

	// calc route
	x := make([][]int, M)
	for i := 0; i < M; i++ {
		route = make([]int, 0)
		DFS(u[i], v[i], -1)
		//out("from", u[i], "to", v[i], "route", route)
		x[i] = make([]int, N-1)
		for _, v := range route {
			x[i][v] = 1
		}
	}

	// change route to bit
	bits := make([]int, M)
	for i := 0; i < M; i++ {
		for j := 0; j < N-1; j++ {
			bits[i] <<= 1
			if x[i][j] == 1 {
				bits[i] |= 1
			}
		}
	}

	// calc number of path
	ans := 1 << uint(N-1)
	n := 1 << uint(M)
	for i := 1; i < n; i++ {
		l := 0
		idx := 0
		cnt := 0
		for j := i; j != 0; j >>= 1 {
			if j%2 == 1 {
				l |= bits[idx]
				cnt++
			}
			idx++
		}
		//		fmt.Printf("bits %d = %b cnt %d\n", i, l, cnt)
		flg := 0
		for j := l; j != 0; j >>= 1 {
			if j%2 == 1 {
				flg++
			}
		}
		flg = N - 1 - flg
		//		out(flg)
		if cnt%2 == 1 {
			ans -= 1 << uint(flg)
		} else {
			ans += 1 << uint(flg)
		}
	}
	out(ans)
}
