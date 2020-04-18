package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
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

// Data :
type Data struct {
	X, N int
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
	return p[i].X > p[j].X
}

func main() {
	sc.Split(bufio.ScanWords)

	N, M := getInt(), getInt()
	x := make(map[int]int)
	for i := 0; i < N; i++ {
		a := getInt()
		x[a]++
	}
	for i := 0; i < M; i++ {
		b, c := getInt(), getInt()
		x[c] += b
	}

	y := make(Datas, 0)
	for i, v := range x {
		y = append(y, Data{i, v})
	}

	sort.Sort(y)

	ans := 0
	cnt := 0
	for _, v := range y {
		n := min(v.N, N-cnt)
		ans += n * v.X
		cnt += n
		if cnt == N {
			break
		}
	}
	// out(y)
	out(ans)

}
