package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

var sc = bufio.NewScanner(os.Stdin)
var wr = bufio.NewWriter(os.Stdout)

func out(x ...interface{}) {
	fmt.Fprintln(wr, x...)
}

func outSlice[T any](s []T) {
	if len(s) == 0 {
		return
	}
	for i := 0; i < len(s)-1; i++ {
		fmt.Fprint(wr, s[i], " ")
	}
	fmt.Fprintln(wr, s[len(s)-1])
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

func getStrings(N int) []string {
	ret := make([]string, N)
	for i := 0; i < N; i++ {
		ret[i] = getS()
	}
	return ret
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

// 値を圧縮した配列を返す
func compressArray(a []int) []int {
	m := make(map[int]int)
	for _, e := range a {
		m[e] = 1
	}
	b := make([]int, 0)
	for e := range m {
		b = append(b, e)
	}
	sort.Ints(b)
	for i, e := range b {
		m[e] = i
	}

	ret := make([]int, len(a))
	for i, e := range a {
		ret[i] = m[e]
	}
	return ret
}

// 平衡二分探索木
var (
	y uint = 88172645463325252
)

const inf = math.MaxInt64

func xorshift() uint {
	y ^= y << 7
	y ^= y >> 9
	return y
}

type treapNode struct {
	value    int
	priority uint
	count    int
	left     *treapNode
	right    *treapNode
}

func newTreapNode(v int) *treapNode {
	return &treapNode{v, xorshift(), 1, nil, nil}
}

func treapRotateRight(n *treapNode) *treapNode {
	l := n.left
	n.left = l.right
	l.right = n
	return l
}

func treapRotateLeft(n *treapNode) *treapNode {
	r := n.right
	n.right = r.left
	r.left = n
	return r
}

func treapInsert(n *treapNode, v int) *treapNode {
	if n == nil {
		return newTreapNode(v)
	}
	if n.value == v {
		n.count++
		return n
	}
	if n.value > v {
		n.left = treapInsert(n.left, v)
		if n.priority > n.left.priority {
			n = treapRotateRight(n)
		}
	} else {
		n.right = treapInsert(n.right, v)
		if n.priority > n.right.priority {
			n = treapRotateLeft(n)
		}
	}
	return n
}

func treapDelete(n *treapNode, v int) *treapNode {
	if n == nil {
		panic("node is not found!")
	}
	if n.value > v {
		n.left = treapDelete(n.left, v)
		return n
	}
	if n.value < v {
		n.right = treapDelete(n.right, v)
		return n
	}

	// n.value == v
	if n.count > 1 {
		n.count--
		return n
	}

	if n.left == nil && n.right == nil {
		return nil
	}

	if n.left == nil {
		n = treapRotateLeft(n)
	} else if n.right == nil {
		n = treapRotateRight(n)
	} else {
		// n.left != nil && n.right != nil
		if n.left.priority < n.right.priority {
			n = treapRotateRight(n)
		} else {
			n = treapRotateLeft(n)
		}
	}
	return treapDelete(n, v)
}

func treapCount(n *treapNode) int {
	if n == nil {
		return 0
	}
	return n.count + treapCount(n.left) + treapCount(n.right)
}

func treapString(n *treapNode) string {
	if n == nil {
		return ""
	}
	result := make([]string, 0)
	if n.left != nil {
		result = append(result, treapString(n.left))
	}
	result = append(result, fmt.Sprintf("%d:%d", n.value, n.count))
	if n.right != nil {
		result = append(result, treapString(n.right))
	}
	return strings.Join(result, " ")
}

func treapMin(n *treapNode) int {
	if n.left != nil {
		return treapMin(n.left)
	}
	return n.value
}

func treapLEMax(n *treapNode, v int) int {
	ret := -inf
	t := n
	for t != nil {
		if t.value <= v {
			ret = max(ret, t.value)
		}
		if t.value > v {
			t = t.left
		} else {
			t = t.right
		}
	}

	return ret
}

func treapGEMin(n *treapNode, v int) int {
	ret := inf
	t := n
	for t != nil {
		if t.value >= v {
			ret = min(ret, t.value)
		}
		if t.value < v {
			t = t.right
		} else {
			t = t.left
		}
	}

	return ret
}

func treapMax(n *treapNode) int {
	if n.right != nil {
		return treapMax(n.right)
	}
	return n.value
}

type treap struct {
	root *treapNode
	size int
}

func (t *treap) Insert(v int) {
	t.root = treapInsert(t.root, v)
	t.size++
}

func (t *treap) Delete(v int) {
	t.root = treapDelete(t.root, v)
	t.size--
}

func (t *treap) String() string {
	return treapString(t.root)
}

func (t *treap) Count() int {
	return t.size
}

func (t *treap) Min() int {
	return treapMin(t.root)
}

func (t *treap) LEMax(v int) int {
	return treapLEMax(t.root, v)
}

func (t *treap) Max() int {
	return treapMax(t.root)
}

func (t *treap) GEMin(v int) int {
	return treapGEMin(t.root, v)
}

func main() {
	defer wr.Flush()
	sc.Split(bufio.ScanWords)
	sc.Buffer([]byte{}, math.MaxInt32)
	// this template is new version.
	// use getI(), getS(), getInts(), getF()
	N, Q := getI(), getI()

	t := &treap{}
	mR := make(map[int]int)
	blkCnt := 0

	for i := 0; i < Q; i++ {
		L, R := getI(), getI()
		newL, newR := L, R

		// newL以下の最大の開始位置を取得
		l_start := t.LEMax(newL)

		// l_startが存在し (-infでない)、かつ その区間が今回の区間と接する/重なる場合
		if l_start != -inf {
			l_end := mR[l_start]
			// 接する条件: 既存の右端 >= 今回の左端 - 1
			if l_end >= newL-1 {
				// マージする
				newL = min(newL, l_start)
				newR = max(newR, l_end)
				// 既存の区間を削除する準備
				blkCnt -= (l_end - l_start + 1)
				t.Delete(l_start)
				delete(mR, l_start)
			}
		}

		// 現在の newL 以上の開始位置を持つ区間を順に見ていく
		for {
			// newL以上の最小の開始位置を取得
			// ※前のステップで削除を行っているため、GEMinは常に「次に処理すべき区間」を見つける
			from := t.GEMin(newL)

			// 区間が存在しない、または次の区間が今回の区間の右端よりさらに右にある場合
			if from == inf {
				break
			}
			if from > newR+1 {
				break
			}

			// マージ処理
			to := mR[from]
			newR = max(newR, to)
			// 左端は newL のまま（from >= newL なので）

			// 既存区間の削除
			blkCnt -= (to - from + 1)
			t.Delete(from)
			delete(mR, from)
		}

		// 3. マージ後の新しい区間を登録
		t.Insert(newL)
		mR[newL] = newR
		blkCnt += (newR - newL + 1)

		// 出力 (白マスの数 = N - 黒マスの数)
		out(N - blkCnt)
	}

}
