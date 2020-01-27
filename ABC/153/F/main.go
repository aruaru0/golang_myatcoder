package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

// Data :
type Data struct {
	X int
	H int
}

// Datas :
type Datas []Data

func (p Datas) Len() int {
	return len(p)
}

func (p Datas) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

func (p Datas) Less(i, j int) bool {
	return p[i].X < p[j].X
}

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

type queue struct {
	pos int
	dmg int
}

func main() {
	sc.Split(bufio.ScanWords)

	N := getInt()
	D := getInt()
	A := getInt()
	m := make(Datas, N)
	for i := 0; i < N; i++ {
		m[i].X = getInt()
		m[i].H = getInt()
	}

	sort.Sort(m)

	q := make([]queue, 0, N)

	ans := 0
	D = 2 * D
	T := 0
	for i := 0; i < N; i++ {
		pos := m[i].X
		//		out("pos", pos, len(q))
		for len(q) != 0 && q[0].pos < pos {
			//			out("del", q[0], "qpos", q[0].pos, "pos", pos)
			T -= q[0].dmg
			q = q[1:]
		}
		x := m[i].H
		x -= T
		//		out("attack", m[i].X, m[i].H, T, pos, x)

		if x > 0 {
			num := (x + A - 1) / A
			ans += num
			dmg := num * A
			T += dmg
			q = append(q, queue{pos + D, dmg})
			//			out(q)
		}

	}
	fmt.Println(ans)
}
