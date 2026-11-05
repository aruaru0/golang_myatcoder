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

// PfsMap : 素因数分解し、マップを作成
func PfsMap(n int) map[int]int {
	pfs := make(map[int]int)
	for n%2 == 0 {
		pfs[2] = pfs[2] + 1
		n = n / 2
	}

	for i := 3; i*i <= n; i = i + 2 {
		for n%i == 0 {
			pfs[i] = pfs[i] + 1
			n = n / i
		}
	}

	if n > 2 {
		pfs[n] = pfs[n] + 1
	}

	return pfs
}

// Pfs :　素因数分解し、スライスを作成
func Pfs(n int) (pfs []int) {
	for n%2 == 0 {
		pfs = append(pfs, 2)
		n = n / 2
	}

	for i := 3; i*i <= n; i = i + 2 {
		for n%i == 0 {
			pfs = append(pfs, i)
			n = n / i
		}
	}

	if n > 2 {
		pfs = append(pfs, n)
	}

	return
}

func main() {
	defer wr.Flush()
	sc.Split(bufio.ScanWords)
	sc.Buffer([]byte{}, math.MaxInt32)
	// this template is new version.
	// use getI(), getS(), getInts(), getF()
	N, M := getI(), getI()

	// Pfs 関数で N の素因数スライスを取得し、mod M を取っておく
	rawP := Pfs(N)
	P := make([]int, len(rawP))
	for i, v := range rawP {
		P[i] = v % M
	}

	L := len(P)
	ans := make([]int, M)
	for i := 0; i < M; i++ {
		ans[i] = -1
	}

	// 状態数: (L + 1) * M
	totalStates := (L + 1) * M
	visited := make([]uint64, (totalStates+63)/64)

	isVisited := func(s, u int) bool {
		idx := s*M + u
		return (visited[idx>>6] & (uint64(1) << (idx & 63))) != 0
	}

	setVisited := func(s, u int) {
		idx := s*M + u
		visited[idx>>6] |= (uint64(1) << (idx & 63))
	}

	// 初期状態: t = 1, x = 1 (mod M)
	startU := 1 % M
	setVisited(0, startU)
	ans[startU] = 1
	remaining := M - 1
	if M == 1 {
		remaining = 0
	}

	qCur := make([]int, 0, M)
	qNext := make([]int, 0, M)
	qNextRound := make([]int, 0, M)

	qCur = append(qCur, startU)
	t := 1

	for len(qCur) > 0 && remaining > 0 {
		// ステージ 0 -> 1 -> ... -> L の遷移 (コスト 0)
		for s := 0; s < L; s++ {
			p := P[s]
			qNext = qNext[:0]

			for _, u := range qCur {
				// 遷移 1: 素因数 P[s] を掛ける
				nu := int((int64(u) * int64(p)) % int64(M))
				if !isVisited(s+1, nu) {
					setVisited(s+1, nu)
					if ans[nu] == -1 {
						ans[nu] = t
						remaining--
					}
					qNext = append(qNext, nu)
				}

				// 遷移 2: スキップする
				if !isVisited(s+1, u) {
					setVisited(s+1, u)
					if ans[u] == -1 {
						ans[u] = t
						remaining--
					}
					qNext = append(qNext, u)
				}
			}

			qCur, qNext = qNext, qCur
		}

		// ステージ L から次のラウンドのステージ 0 へ遷移 (コスト 1)
		t++
		qNextRound = qNextRound[:0]
		for _, u := range qCur {
			if !isVisited(0, u) {
				setVisited(0, u)
				qNextRound = append(qNextRound, u)
			}
		}

		qCur = append(qCur[:0], qNextRound...)
	}

	for k := 0; k < M; k++ {
		out(ans[k])
	}
}
