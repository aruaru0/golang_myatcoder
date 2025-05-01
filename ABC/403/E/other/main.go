package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
)

/* ---------- I/O ユーティリティ ---------- */

var sc = bufio.NewScanner(os.Stdin)
var wr = bufio.NewWriter(os.Stdout)

func out(x ...interface{}) { fmt.Fprintln(wr, x...) }

func getI() int {
	sc.Scan()
	v, _ := strconv.Atoi(sc.Text())
	return v
}

func getS() string { sc.Scan(); return sc.Text() }

/* ---------- 汎用 ---------- */

func lowerBound(a []int, x int) int {
	return sort.Search(len(a), func(i int) bool { return a[i] >= x })
}

/* ----------------------------------------------------------------
   Trie（ポインタ版）
---------------------------------------------------------------- */

type Node struct {
	ch   map[rune]*Node // 子ノード
	ng   bool           // X で”封印”されたら true
	numY int            // ここに Y が何回置かれたか
}

type Trie struct {
	root *Node
	ans  int
}

func NewTrie() *Trie { return &Trie{root: &Node{ch: make(map[rune]*Node)}} }

// 文字列を挿入し、終端ノードへのポインタを返す
func (t *Trie) Add(s string) *Node {
	v := t.root
	for _, c := range s {
		if v.ch[c] == nil {
			v.ch[c] = &Node{ch: make(map[rune]*Node)}
		}
		v = v.ch[c]
	}
	return v
}

func (t *Trie) Init() { t.ans = 0 } // 追加直後は全ノードが未使用なので 0 で十分

// X クエリ
func (t *Trie) AddX(v *Node) {
	if v.ng {
		return
	}
	v.ng = true
	t.ans -= v.numY
	for _, u := range v.ch {
		t.AddX(u)
	}
}

// Y クエリ
func (t *Trie) AddY(v *Node) {
	if v.ng {
		return
	}
	t.ans++
	v.numY++
}

/* ---------- main ---------- */

func main() {
	defer wr.Flush()
	sc.Split(bufio.ScanWords)
	sc.Buffer(make([]byte, 0, 1<<20), math.MaxInt32)

	q := getI()
	tr := NewTrie()

	// クエリを一旦ためる（典型的 AtCoder テク）
	type query struct {
		typ int
		v   *Node
	}
	qs := make([]query, q)

	for i := 0; i < q; i++ {
		typ := getI()
		s := getS()
		v := tr.Add(s)
		qs[i] = query{typ, v}
	}
	tr.Init()

	for _, qu := range qs {
		if qu.typ == 1 {
			tr.AddX(qu.v)
		} else {
			tr.AddY(qu.v)
		}
		out(tr.ans)
	}
}
