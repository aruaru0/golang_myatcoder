package main

import (
	"bufio"
	"fmt"
	"math/rand"
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

func getInts(N int) []int {
	ret := make([]int, N)
	for i := 0; i < N; i++ {
		ret[i] = getInt()
	}
	return ret
}

func getString() string {
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

func calcScore(day int, score, c, last []int) (int, []int) {
	loss := make([]int, 26)
	for i := 0; i < 26; i++ {
		loss[i] = c[i] * (day - last[i])
	}

	ret := 0
	new := make([]int, 26)
	for j := 0; j < 26; j++ {
		new[j] = score[j] - loss[j]
		ret += new[j]
	}

	return ret, new
}

func calcAllScore(D int, days []int, c []int, s [][]int) int {
	last := make([]int, 26)
	score := make([]int, 26)

	ans := 0
	for i := 0; i < D; i++ {
		t := days[i]
		score[t] += s[i][t]
		last[t] = i + 1
		ret, nw := calcScore(i+1, score, c, last)
		score = nw
		ans = ret
		// out(score)
	}
	return ans
}

var tab [366][26]int
var sum [366][26]int
var exe [366]int
var D int
var c []int
var s [][]int

func calc(typ int) {
	score := 0
	last := 0
	for i := 1; i <= D; i++ {
		if tab[i][typ] == 1 {
			score += s[i][typ]
			last = i
		}
		score += c[typ] * (last - i)
		sum[i][typ] = score
	}
}

var last [26]int

func calcNow(day int) {
	for i := 0; i < 26; i++ {
		sum[day][i] = sum[day-1][i] + c[i]*(last[i]-day)
	}
}

type pair struct {
	day, typ int
}

func main() {
	sc.Split(bufio.ScanWords)
	D = getInt()
	c = getInts(26)
	s = make([][]int, D+1)
	for i := 0; i < D; i++ {
		s[i+1] = getInts(26)
	}

	for i := 1; i <= D; i++ {
		calcNow(i)
		// out(sum[i-1])
		// out(sum[i])
		tot := 0
		for j := 0; j < 26; j++ {
			tot += sum[i][j]
		}
		maxS := sum[i][0]
		idx := 0
		for j := 0; j < 26; j++ {
			if s[i][j] > maxS {
				maxS = s[i][j]
				idx = j
			}
		}
		tab[i][idx] = 1
		exe[i] = idx + 1
		sum[i][idx] += s[i][idx]
		last[idx] = i
	}

	// 	t := 0
	// 	tab[i][t] = 1
	// 	exe[i] = t + 1
	// }
	for i := 0; i < 26; i++ {
		calc(i)
	}
	ans := 0
	for i := 0; i < 26; i++ {
		ans += sum[D][i]
	}

	var day, typ int
	for k := 0; k < 600000; k++ {
		sel := rand.Intn(2)
		if sel == 0 {
			day = k%365 + 1 //rand.Intn(364) + 1
			typ = (typ + k) % 26
		} else {
			day = rand.Intn(364) + 1
			typ = rand.Intn(26)
		}

		old := exe[day] - 1
		tab[day][typ] = 1
		tab[day][old] = 0
		osA := sum[D][typ]
		osB := sum[D][old]
		// for i := 0; i < 26; i++ {
		// 	calc(i)
		// }
		calc(typ)
		calc(old)
		tot := 0
		for i := 0; i < 26; i++ {
			tot += sum[D][i]
		}
		if tot > ans {
			// out(tot, ans)
			ans = tot
			exe[day] = typ + 1
		} else {
			tab[day][typ] = 0
			tab[day][old] = 1
			sum[D][typ] = osA
			sum[D][old] = osB
		}
	}

	// for i := 0; i < 365; i++ {
	// 	out(tab[i])
	// }

	// ans := 0
	// for i := 0; i < 26; i++ {
	// 	ans += sum[D][i]
	// }
	// out(ans)

	wr := bufio.NewWriter(os.Stdout)
	defer wr.Flush()
	for i := 1; i <= D; i++ {
		fmt.Fprintln(wr, exe[i])
	}
}
