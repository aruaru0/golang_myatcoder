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

type minheap struct {
	buf  []int
	less func(int, int) bool
}

func Newminheap(f func(int, int) bool) *minheap { buf := make([]int, 0); return &minheap{buf, f} }
func (q *minheap) IsEmpty() bool                { return len(q.buf) == 0 }
func (q *minheap) Clear()                       { q.buf = q.buf[:0] }
func (q *minheap) Len() int                     { return len(q.buf) }
func (q *minheap) Push(v int)                   { q.buf = append(q.buf, v); q.siftdown(0, len(q.buf)-1) }
func (q *minheap) Head() int                    { return q.buf[0] }
func (q *minheap) Pop() int {
	v1 := q.buf[0]
	l := len(q.buf)
	if l == 1 {
		q.buf = q.buf[:0]
	} else {
		l--
		q.buf[0] = q.buf[l]
		q.buf = q.buf[:l]
		q.siftup(0)
	}
	return v1
}
func (q *minheap) Heapify(pri []int) {
	q.buf = append(q.buf, pri...)
	n := len(q.buf)
	for i := n/2 - 1; i >= 0; i-- {
		q.siftup(i)
	}
}
func (q *minheap) siftdown(startpos, pos int) {
	newitem := q.buf[pos]
	for pos > startpos {
		ppos := (pos - 1) >> 1
		p := q.buf[ppos]
		if !q.less(newitem, p) {
			break
		}
		q.buf[pos], pos = p, ppos
	}
	q.buf[pos] = newitem
}
func (q *minheap) siftup(pos int) {
	endpos, startpos, newitem, chpos := len(q.buf), pos, q.buf[pos], 2*pos+1
	for chpos < endpos {
		rtpos := chpos + 1
		if rtpos < endpos && !q.less(q.buf[chpos], q.buf[rtpos]) {
			chpos = rtpos
		}
		q.buf[pos], pos = q.buf[chpos], chpos
		chpos = 2*pos + 1
	}
	q.buf[pos] = newitem
	q.siftdown(startpos, pos)
}

func main() {
	defer wr.Flush()
	sc.Split(bufio.ScanWords)
	sc.Buffer([]byte{}, math.MaxInt32)
	// this template is new version.
	// use getI(), getS(), getInts(), getF()
	N := getI()
	T := make([]int, N)
	D := make([]int, N)
	X := make([]int, N)
	for i := 0; i < N; i++ {
		T[i] = getI()
		D[i] = getI()
		X[i] = getI()
	}
	iless := func(a, b int) bool { return a < b }
	igreater := func(a, b int) bool { return a > b }
	tmp := make([]int, N+5)
	mhleft := Newminheap(igreater)
	mhleft.Heapify(tmp)
	mhright := Newminheap(iless)
	mhright.Heapify(tmp)
	minval, addL, addR, mytime := 0, 0, 0, 0
	pushleft := func(x int) { mhleft.Push(x - addL) }
	pushright := func(x int) { mhright.Push(x - addR) }
	topleft := func() int { return mhleft.Head() + addL }
	topright := func() int { return mhright.Head() + addR }
	popleft := func() int { v := topleft(); mhleft.Pop(); return v }
	popright := func() int { v := topright(); mhright.Pop(); return v }
	addRightDamage := func(x, minval int) int {
		minval += max(0, topleft()-x)
		pushleft(x)
		pushright(popleft())
		return minval
	}
	addLeftDamage := func(x, minval int) int {
		minval += max(0, x-topright())
		pushright(x)
		pushleft(popright())
		return minval
	}
	for i := 0; i < N; i++ {
		t, d, x := T[i], D[i], X[i]
		addL -= t - mytime
		addR += (t - mytime)
		mytime = t
		if d == 0 {
			minval = addLeftDamage(x, minval)
		} else {
			minval = addRightDamage(x, minval)
		}
	}
	fmt.Println(minval)
}
