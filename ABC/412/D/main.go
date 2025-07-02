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

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func combinations(list []int, k, buf int) (c chan []int) {
	c = make(chan []int, buf)
	n := len(list)
	pattern := make([]int, k)

	var body func(pos, begin int)
	body = func(pos, begin int) {
		if pos == k {
			t := make([]int, k)
			copy(t, pattern)
			c <- t
			return
		}
		for num := begin; num < n+pos-k+1; num++ {
			pattern[pos] = list[num]
			body(pos+1, num+1)
		}
	}
	go func() {
		defer close(c)
		body(0, 0)
	}()
	return
}

func generatePartitions(n int) [][]int {
	var res [][]int
	var dfs func(remaining int, path []int)
	dfs = func(remaining int, path []int) {
		if remaining == 0 {
			copyPath := make([]int, len(path))
			copy(copyPath, path)
			res = append(res, copyPath)
			return
		}
		start := 3
		if len(path) > 0 {
			start = max(start, path[len(path)-1])
		}
		for i := start; i <= remaining; i++ {
			path = append(path, i)
			dfs(remaining-i, path)
			path = path[:len(path)-1]
		}
	}
	dfs(n, []int{})
	return res
}

func generateGroupings(nodes []int, partition []int) [][][]int {
	if len(partition) == 0 {
		return [][][]int{{}}
	}
	first := partition[0]
	rest := partition[1:]
	var groupings [][][]int
	for combo := range combinations(nodes, first, 100000) {
		nodeMap := make(map[int]bool)
		for _, v := range combo {
			nodeMap[v] = true
		}
		var remaining []int
		for _, x := range nodes {
			if !nodeMap[x] {
				remaining = append(remaining, x)
			}
		}
		subGroupings := generateGroupings(remaining, rest)
		for _, sg := range subGroupings {
			group := [][]int{append([]int{}, combo...)}
			group = append(group, sg...)
			groupings = append(groupings, group)
		}
	}
	return groupings
}

func generateCycles(group []int) []map[[2]int]struct{} {
	sort.Ints(group)
	var result = make(map[string]map[[2]int]struct{})
	perm := append([]int{}, group...)
	sortPerms := func(a []int) string {
		s := make([][2]int, len(a))
		for i := range a {
			u, v := a[i], a[(i+1)%len(a)]
			if u > v {
				u, v = v, u
			}
			s[i] = [2]int{u, v}
		}
		sort.Slice(s, func(i, j int) bool {
			if s[i][0] == s[j][0] {
				return s[i][1] < s[j][1]
			}
			return s[i][0] < s[j][0]
		})
		key := ""
		m := make(map[[2]int]struct{})
		for _, e := range s {
			key += fmt.Sprintf("%d-%d,", e[0], e[1])
			m[e] = struct{}{}
		}
		return key
	}

	for {
		key := sortPerms(perm)
		if _, ok := result[key]; !ok {
			m := make(map[[2]int]struct{})
			for i := 0; i < len(perm); i++ {
				u := perm[i]
				v := perm[(i+1)%len(perm)]
				if u > v {
					u, v = v, u
				}
				m[[2]int{u, v}] = struct{}{}
			}
			result[key] = m
		}
		if !nextPermutation(sort.IntSlice(perm)) {
			break
		}
	}
	var ret []map[[2]int]struct{}
	for _, v := range result {
		ret = append(ret, v)
	}
	return ret
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

func main() {
	defer wr.Flush()
	sc.Split(bufio.ScanWords)
	sc.Buffer([]byte{}, math.MaxInt32)

	N := getI()
	M := getI()
	original := make(map[[2]int]struct{})
	for i := 0; i < M; i++ {
		a, b := getI(), getI()
		if a > b {
			a, b = b, a
		}
		original[[2]int{a, b}] = struct{}{}
	}

	nodes := make([]int, N)
	for i := 0; i < N; i++ {
		nodes[i] = i + 1
	}

	minOps := math.MaxInt32
	for _, part := range generatePartitions(N) {
		groupings := generateGroupings(nodes, part)
		for _, grouping := range groupings {
			var cycleOptions [][]map[[2]int]struct{}
			for _, group := range grouping {
				cycleOptions = append(cycleOptions, generateCycles(group))
			}
			var recur func(int, map[[2]int]struct{})
			recur = func(pos int, acc map[[2]int]struct{}) {
				if pos == len(cycleOptions) {
					ops := 0
					for k := range acc {
						if _, ok := original[k]; !ok {
							ops++
						}
					}
					for k := range original {
						if _, ok := acc[k]; !ok {
							ops++
						}
					}
					if ops < minOps {
						minOps = ops
					}
					return
				}
				for _, cycle := range cycleOptions[pos] {
					next := make(map[[2]int]struct{})
					for k := range acc {
						next[k] = struct{}{}
					}
					for k := range cycle {
						next[k] = struct{}{}
					}
					recur(pos+1, next)
				}
			}
			recur(0, map[[2]int]struct{}{})
		}
	}
	out(minOps)
}
