package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"

	rbt "github.com/emirpasic/gods/trees/redblacktree"
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

func getInts(N int) []int {
	ret := make([]int, N)
	for i := 0; i < N; i++ {
		ret[i] = getI()
	}
	return ret
}

func main() {
	defer wr.Flush()
	sc.Split(bufio.ScanWords)
	sc.Buffer([]byte{}, math.MaxInt32)
	// this template is new version.
	// use getI(), getS(), getInts(), getF()
	N, D := getI(), getI()
	a := getInts(N)

	const inf = int(1e15)
	s := rbt.NewWithIntComparator()
	s.Put(-inf, -inf)
	s.Put(inf, inf)

	ans, l := 0, 0

	for r := 0; r < N; r++ {
		for {
			left, _ := s.Floor(a[r])
			right, _ := s.Ceiling(a[r])
			if right.Value.(int)-a[r] >= D && a[r]-left.Value.(int) >= D {
				break
			}
			s.Remove(a[l])
			l++
		}
		s.Put(a[r], a[r])
		ans += r - l + 1
	}
	out(ans)

}
