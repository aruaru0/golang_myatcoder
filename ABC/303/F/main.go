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

func bs(ok, ng int, f func(int) bool) int {
	if !f(ok) {
		return -1
	}
	if f(ng) {
		return ng
	}
	for abs(ok-ng) > 1 {
		mid := (ok + ng) / 2

		if f(mid) {
			ok = mid
		} else {
			ng = mid
		}
	}

	return ok
}

const inf = int(1e18)

func main() {
	// とりあえず、写経。。。。
	defer wr.Flush()
	sc.Split(bufio.ScanWords)
	sc.Buffer([]byte{}, math.MaxInt32)
	// this template is new version.
	// use getI(), getS(), getInts(), getF()
	n, h := getI(), getI()

	type magic struct {
		t int
		d int
	}
	magics := make([]magic, n)
	for i := 0; i < n; i++ {
		t, d := getI(), getI()
		magics[i] = magic{t: t, d: d}
	}

	// 効果の大きい順、同じ場合はターンが長い順にソート
	sort.Slice(magics, func(i, j int) bool {
		if magics[i].d == magics[j].d {
			return magics[i].t > magics[j].t
		}
		return magics[i].d > magics[j].d
	})

	// 前から順に魔法を見ていって、あと何ターン残っていたらどの魔法を使うかをメモする
	// あと　turn[i] 以上のターンが残っていたら、mag[i]番目の魔法を使う
	turn := []int{}
	mag := []int{}
	max := 0
	for i := 1; i < n; i++ {
		if magics[max].d*magics[max].t < magics[i].d*magics[i].t {
			// ある程度のターンからは、iのほうが効果が高くなる
			turn = append(turn, magics[max].d*magics[max].t/magics[i].d)
			mag = append(mag, max)
			max = i
		}
	}
	mag = append(mag, max)
	turn = append(turn, inf)

	// 二分探索を行う
	// magic[0]しか使えないとしても、h/(magic[0].d*magic[0].t)+magic[0].t-1 ターンで倒せる
	maxturn := h/(magics[mag[0]].d*magics[mag[0]].t) + magics[mag[0]].t

	res := bs(maxturn, 0, func(x int) bool {
		// xターン目までに倒せるかどうか
		// 残りターン数をiとする
		i := 0
		j := 0
		res := 0
		for i < x {
			// 今使う魔法はmag[j]
			// 残りターン数がmin(turn[j], magics[mag[j]].t)を超えるまではこの魔法の効果が高まる
			k := min(min(turn[j], magics[mag[j]].t), x)
			if i < k {
				// res がオーバーフローしそうなら、ここで終了
				if (h-res+magics[mag[j]].d-1)/magics[mag[j]].d <= ((k+1)*k/2 - (i+1)*i/2) {
					return true
				}
				res += magics[mag[j]].d * ((k+1)*k/2 - (i+1)*i/2)
				i = k
			}
			// 残りターン数がmin(turn[j], x)を超えるまではこの魔法を使い続ける
			l := min(turn[j], x)
			if i < l {
				// res がオーバーフローしそうなら、ここで終了
				if (h-res+l-i-1)/(l-i) <= magics[mag[j]].d*magics[mag[j]].t {
					return true
				}
				res += magics[mag[j]].d * magics[mag[j]].t * (l - i)
				i = l
			}
			j++
		}
		if res >= h {
			return true
		} else {
			return false
		}
	})

	fmt.Println(res)

}
