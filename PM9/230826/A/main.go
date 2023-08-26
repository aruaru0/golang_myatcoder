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

type Point struct {
	x, y int
}

// Max
func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// Min
func Min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// MinMaxCrossCheck 区間 p1,p2 と　p3,p4に重なりがあるか
func MinMaxCrossCheck(p1, p2, p3, p4 int) bool {
	min_ab, max_ab := Min(p1, p2), Max(p1, p2)
	min_cd, max_cd := Min(p3, p4), Max(p3, p4)

	if min_ab > max_cd || max_ab < min_cd {
		return false
	}
	return true
}

// CrossJudge 線分ABと線分CDが交差しているか判定
func CrossJudge(a, b, c, d Point) bool {
	tc1 := (a.x-b.x)*(c.y-a.y) + (a.y-b.y)*(a.x-c.x)
	tc2 := (a.x-b.x)*(d.y-a.y) + (a.y-b.y)*(a.x-d.x)
	td1 := (c.x-d.x)*(a.y-c.y) + (c.y-d.y)*(c.x-a.x)
	td2 := (c.x-d.x)*(b.y-c.y) + (c.y-d.y)*(c.x-b.x)

	ab := false
	cd := false
	if tc1 >= 0 && tc2 <= 0 {
		ab = true
	}
	if tc1 <= 0 && tc2 >= 0 {
		ab = true
	}
	if td1 >= 0 && td2 <= 0 {
		cd = true
	}
	if td1 <= 0 && td2 >= 0 {
		cd = true
	}
	return ab && cd
}

// CrossJudge2 線分ABと線分CDが交差しているか判定（線分ABとCDの一部が一直線に重なっている場合も交差と判定）
func CrossJudge2(a, b, c, d Point) bool {
	// xによる重なり判定
	if MinMaxCrossCheck(a.x, b.x, c.x, d.x) == false {
		return false
	}
	// yによる重なり判定
	if MinMaxCrossCheck(a.y, b.y, c.y, d.y) == false {
		return false
	}

	tc1 := (a.x-b.x)*(c.y-a.y) + (a.y-b.y)*(a.x-c.x)
	tc2 := (a.x-b.x)*(d.y-a.y) + (a.y-b.y)*(a.x-d.x)
	td1 := (c.x-d.x)*(a.y-c.y) + (c.y-d.y)*(c.x-a.x)
	td2 := (c.x-d.x)*(b.y-c.y) + (c.y-d.y)*(c.x-b.x)

	ab := false
	cd := false
	if tc1 >= 0 && tc2 <= 0 {
		ab = true
	}
	if tc1 <= 0 && tc2 >= 0 {
		ab = true
	}
	if td1 >= 0 && td2 <= 0 {
		cd = true
	}
	if td1 <= 0 && td2 >= 0 {
		cd = true
	}
	return ab && cd
}

func main() {
	defer wr.Flush()
	sc.Split(bufio.ScanWords)
	sc.Buffer([]byte{}, math.MaxInt32)
	// this template is new version.
	// use getI(), getS(), getInts(), getF()
	a := Point{getI(), getI()}
	b := Point{getI(), getI()}
	c := Point{getI(), getI()}
	d := Point{getI(), getI()}

	if CrossJudge2(a, b, c, d) {
		out("Yes")
	} else {
		out("No")
	}
}
