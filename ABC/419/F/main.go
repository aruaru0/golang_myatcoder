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

func main() {
	defer wr.Flush()
	sc.Split(bufio.ScanWords)
	sc.Buffer([]byte{}, math.MaxInt32)
	// this template is new version.
	// use getI(), getS(), getInts(), getF()
	n, l := getI(), getI()
	S := getStrings(n)

	// sunukeさんのコードをGoに変換しました。

	// 1. すべての接頭辞を抽出し、状態として管理
	id := make(map[string]int)
	id[""] = 0 // 空文字列も状態の一つ
	for i := 0; i < n; i++ {
		for j := 0; j < len(S[i]); j++ {
			prefix := S[i][:j+1]
			id[prefix] = 0
		}
	}

	// 2. 接頭辞をソートし、IDを確定させる
	// (Goのmapは順序を保証しないため、C++のstd::mapの挙動を再現するためにソートが必要)
	keys := make([]string, 0, len(id))
	for k := range id {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	m := 0 // 状態の総数
	sufs := make([]string, len(keys))
	for _, k := range keys {
		id[k] = m
		sufs[m] = k
		m++
	}

	// 3. 各状態(接頭辞)がどのSの末尾と一致するかをビットマスクで記録
	mask := make([]int, m)
	for i := 0; i < m; i++ {
		t := sufs[i]
		for j := 0; j < n; j++ {
			if strings.HasSuffix(t, S[j]) {
				mask[i] |= (1 << j)
			}
		}
	}

	// 4. 状態遷移テーブル `to[j][c]` を作成
	// `to[j][c]` は状態jのときに文字cを追加した場合の遷移先の状態IDを示す
	to := make([][]int, m)
	for j := 0; j < m; j++ {
		to[j] = make([]int, 26)
		for c := 0; c < 26; c++ {
			t := sufs[j] + string('a'+c)
			// tの接頭辞を削りながら、idに存在する最長の接尾辞を探す
			for {
				if _, ok := id[t]; ok {
					break
				}
				if len(t) > 0 {
					t = t[1:]
				}
			}
			to[j][c] = id[t]
		}
	}

	// === 動的計画法 (DP) ===
	var dp [101][1 << 8][85]int
	mod := 998244353

	n2 := 1 << n
	dp[0][0][0] = 1 // 初期状態: 長さ0, 達成マスク0, 状態0(空文字列) の場合の数は1

	for i := 0; i < l; i++ {
		for s := 0; s < n2; s++ {
			for j := 0; j < m; j++ {
				now := dp[i][s][j]
				if now == 0 {
					continue
				}
				// 次の文字 'a' から 'z' を試す
				for c := 0; c < 26; c++ {
					ni := i + 1
					nj := to[j][c]
					ns := s | mask[nj]
					dp[ni][ns][nj] += now
					dp[ni][ns][nj] %= mod
				}
			}
		}
	}

	// === 結果の集計 ===
	ans := 0

	targetMask := n2 - 1 // すべての文字列が達成された状態のマスク
	for j := 0; j < m; j++ {
		ans += dp[l][targetMask][j]
		ans %= mod
	}

	out(ans)
}
