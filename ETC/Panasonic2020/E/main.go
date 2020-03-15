package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

func out(x ...interface{}) {
	fmt.Println(x...)
}

var sc = bufio.NewScanner(os.Stdin)

func getInt() int {
	sc.Scan()
	i, e := strconv.Atoi(sc.Text())
	if e != nil {
		panic(e)
	}
	return i
}

func getString() string {
	sc.Scan()
	return sc.Text()
}

func match(a, b byte) bool {
	if a == 0 {
		return true
	}
	if a == b || a == '?' || b == '?' {
		return true
	}
	return false
}

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

func conv(a []byte) string {
	var s string
	for _, v := range a {
		if v != 0 {
			s += string(v)
		} else {
			s += "."
		}
	}
	return s
}

func nextPermutation(x sort.Interface) bool {
	n := x.Len() - 1
	if n < 1 {
		return false
	}
	j := n - 1
	for ; !x.Less(j, j+1); j-- {
		if j == 0 {
			return false
		}
	}
	l := n
	for !x.Less(j, l) {
		l--
	}
	x.Swap(j, l)
	for k, l := j+1, n; k < l; {
		x.Swap(k, l)
		k++
		l--
	}
	return true
}

func out2(s []string, x, y int) {
	out(s[0])
	for i := 0; i < x; i++ {
		fmt.Print(" ")
	}
	out(s[1])
	for i := 0; i < y; i++ {
		fmt.Print(" ")
	}
	out(s[2])
}

// 解説をyoutubeで見た後に、自力で作成
func main() {
	sc.Split(bufio.ScanWords)
	sc.Buffer([]byte{}, 1000000)

	S := make([]string, 3)
	for i := 0; i < 3; i++ {
		S[i] = getString()
	}

	x := []int{0, 1, 2}

	const SIZE = 2020
	var judge [3][3][SIZE]bool

	ans := 100100
	for {
		s := []string{S[x[0]], S[x[1]], S[x[2]]}
		pat := [][]int{{0, 1}, {0, 2}, {1, 2}}
		for p := 0; p < 3; p++ {
			i := pat[p][0]
			j := pat[p][1]

			for k := 0; k < len(s[i]); k++ {
				flg := true
				for ni := k; ni < len(s[i]); ni++ {
					nj := ni - k
					if nj >= len(s[j]) {
						break
					}
					if s[i][ni] == '?' || s[j][nj] == '?' {
						continue
					}
					if s[i][ni] != s[j][nj] {
						flg = false
					}
				}
				judge[i][j][k] = flg
				//out(i, j, s[i][k:], s[j], judge[i][j][k])
			}
		}

		var f = func(i, j, k int) bool {
			if k >= len(s[i]) {
				return true
			}
			return judge[i][j][k]
		}

		for i := 0; i < SIZE; i++ {
			for j := 0; j < SIZE; j++ {
				if !f(0, 1, i) {
					continue
				}
				if !f(1, 2, j) {
					continue
				}
				//out2(s, i, i+j)
				if !f(0, 2, i+j) {
					continue
				}
				//out2(s, i, i+j)
				//out(len(s[0]), len(s[1])+i, len(s[2])+i+j)
				n := max(len(s[0]), max(len(s[1])+i, len(s[2])+i+j))
				ans = min(ans, n)
			}
		}
		if nextPermutation(sort.IntSlice(x)) == false {
			break
		}
	}

	out(ans)
	/*
		s := make([]string, 3)
		for i:= 0; i < 3; i++  {
			s[i] := getString()
		}
	*/

}
