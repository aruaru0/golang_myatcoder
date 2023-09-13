package main

import (
	"bufio"
	"container/heap"
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

type pq struct {
	arr   []interface{}
	comps []compFunc
}

type compFunc func(p, q interface{}) int

func newpq(comps []compFunc) *pq {
	return &pq{
		comps: comps,
	}
}

func (pq pq) Len() int {
	return len(pq.arr)
}

func (pq pq) Swap(i, j int) {
	pq.arr[i], pq.arr[j] = pq.arr[j], pq.arr[i]
}

func (pq pq) Less(i, j int) bool {
	for _, comp := range pq.comps {
		result := comp(pq.arr[i], pq.arr[j])
		switch result {
		case -1:
			return true
		case 1:
			return false
		case 0:
			continue
		}
	}
	return true
}

func (pq *pq) Push(x interface{}) {
	pq.arr = append(pq.arr, x)
}

func (pq *pq) Pop() interface{} {
	n := pq.Len()
	item := pq.arr[n-1]
	pq.arr = pq.arr[:n-1]
	return item
}

func (pq *pq) IsEmpty() bool {
	return pq.Len() == 0
}

// pq.GetRoot().(edge)
func (pq *pq) GetRoot() interface{} {
	return pq.arr[0]
}

// 今回は、写経しました。オレンジ問題のため、解けないと判断
// 動画解説後に確認予定
func main() {
	defer wr.Flush()
	sc.Split(bufio.ScanWords)
	sc.Buffer([]byte{}, math.MaxInt32)
	// this template is new version.
	// use getI(), getS(), getInts(), getF()
	n := getI()
	G := make([][]int, n) //隣接グラフ
	type v struct {
		i, t, s, g int
	}
	vs := make([]v, n) // 頂点の情報
	md := []int{}      // 薬の頂点
	for i := 1; i < n; i++ {
		p := getI() - 1
		G[p] = append(G[p], i)
		vs[i] = v{i: i, t: getI(), s: getI(), g: getI()}
		if vs[i].t == 2 {
			md = append(md, i)
		}
	}

	// dp[S] = S (bit) の薬を使った時のstrの最大値
	M := len(md)
	dp := make([]int, 1<<M)
	// next[S] = S (bit) の薬を使った時、次に行ける頂点のリスト
	next := make([][]int, 1<<M)

	// 今の強さとアクセスできる頂点、使える薬iをもとに、次の強さとアクセスできる頂点を返す
	solver := func(str, m int, vset []int) (int, []int) {
		enemy := newpq([]compFunc{
			func(p, q interface{}) int {
				if vs[p.(int)].s < vs[q.(int)].s {
					return -1
				} else if vs[p.(int)].s > vs[q.(int)].s {
					return 1
				} else {
					return 0
				}
			}})
		medicine := []int{}
		for _, v := range vset {
			if vs[v].t == 1 {
				heap.Push(enemy, v)
			} else {
				medicine = append(medicine, v)
			}
		}
		// 薬mを使う
		if m > 0 {
			str *= vs[m].g
			if str > 2<<30 {
				str = 2 << 30
			}
			for _, v := range G[m] {
				if vs[v].t == 1 {
					heap.Push(enemy, v)
				} else {
					medicine = append(medicine, v)
				}
			}
			for i := 0; i < len(medicine)-1; i++ {
				if medicine[i] == m {
					copy(medicine[i:], medicine[i+1:])
					break
				}
			}
			medicine = medicine[:len(medicine)-1]
		}
		// 倒せる敵を全部倒す
		for !enemy.IsEmpty() {
			e := heap.Pop(enemy).(int)
			if vs[e].s <= str {
				str += vs[e].g
				if str > 2<<30 {
					str = 2 << 30
				}
				// 次のノードを追加
				for _, next := range G[e] {
					if vs[next].t == 1 {
						heap.Push(enemy, next)
					} else {
						medicine = append(medicine, next)
					}
				}
			} else {
				heap.Push(enemy, e)
				break
			}
		}
		// 行ける頂点をメモ
		accesible := make([]int, len(medicine))
		copy(accesible, medicine)
		for !enemy.IsEmpty() {
			e := heap.Pop(enemy).(int)
			accesible = append(accesible, e)
		}
		return str, accesible
	}

	// init: s=0の場合を解く
	dp[0], next[0] = solver(1, 0, G[0])

	for s := 0; s < 1<<M-1; s++ {
		// 配るDP
		medicine := []int{}
		for _, u := range next[s] {
			if vs[u].t == 2 {
				medicine = append(medicine, u)
			}
		}
		// それぞれの薬を使う場合でDPを更新
		for i := 0; i < M; i++ {
			// Slice md の中に i を含むかどうか判定
			ok := false
			for j := 0; j < len(medicine); j++ {
				if medicine[j] == md[i] {
					ok = true
					break
				}
			}
			if !ok {
				continue
			}

			ns := s | (1 << i)
			str, accesible := solver(dp[s], md[i], next[s])
			if dp[ns] < str {
				// DPを更新
				dp[ns] = str
				// 行ける頂点をメモ
				next[ns] = accesible
			}
		}
	}

	// 最終的な強さで、全ての敵を倒せるか
	for i := 0; i < n; i++ {
		if vs[i].t == 1 && vs[i].s > dp[1<<M-1] {
			fmt.Println("No")
			return
		}
	}
	fmt.Println("Yes")
}
