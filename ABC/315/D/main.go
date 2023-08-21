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

type pair struct {
	f, s int
}

func main() {
	defer wr.Flush()
	sc.Split(bufio.ScanWords)
	sc.Buffer([]byte{}, math.MaxInt32)
	// this template is new version.
	// use getI(), getS(), getInts(), getF()
	H, W := getI(), getI()
	c := make([]string, H)
	for i := 0; i < H; i++ {
		c[i] = getS()
	}
	h := make([][26]int, H)
	w := make([][26]int, W)

	// 行列の文字の数を数えておく
	for i := 0; i < H; i++ {
		for j := 0; j < W; j++ {
			h[i][int(c[i][j]-'a')]++
			w[j][int(c[i][j]-'a')]++
		}
	}

	// 行、列の大きさ
	HC, WC := H, W
	// 削除されたかどうか
	fh := make([]bool, H)
	fw := make([]bool, W)
	//　行ｘ列回ループ
	for r := 0; r < H+W; r++ {
		uh := make([]pair, 0)
		uw := make([]pair, 0)
		// 消されていない行で、幅が現在の最大幅のものを探索
		for i := 0; i < H; i++ {
			if fh[i] {
				continue
			}
			for j := 0; j < 26; j++ {
				if h[i][j] == WC && WC >= 2 {
					uh = append(uh, pair{i, j})
				}
			}
		}
		// 消されていない列で、高さが現在の最大高のものを探索
		for i := 0; i < W; i++ {
			if fw[i] {
				continue
			}
			for j := 0; j < 26; j++ {
				if w[i][j] == HC && HC >= 2 {
					uw = append(uw, pair{i, j})
				}
			}
		}

		//　消去処理
		for _, p := range uh {
			fh[p.f] = true
			for i := 0; i < W; i++ {
				w[i][p.s]--
			}
			HC--
		}
		for _, p := range uw {
			fw[p.f] = true
			for i := 0; i < H; i++ {
				h[i][p.s]--
			}
			WC--
		}
	}

	out(WC * HC)

}
