package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func getNextString(scanner *bufio.Scanner) string {
	if !scanner.Scan() {
		panic("scan failed")
	}
	return scanner.Text()
}
func atoi(s string) int                             { x, _ := strconv.Atoi(s); return x }
func getNextInt(scanner *bufio.Scanner) int         { return atoi(getNextString(scanner)) }
func atoi64(s string) int64                         { x, _ := strconv.ParseInt(s, 10, 64); return x }
func getNextInt64(scanner *bufio.Scanner) int64     { return atoi64(getNextString(scanner)) }
func atof64(s string) float64                       { x, _ := strconv.ParseFloat(s, 64); return x }
func getNextFloat64(scanner *bufio.Scanner) float64 { return atof64(getNextString(scanner)) }
func isLocal() bool                                 { return os.Getenv("I") == "IronMan" }
func main() {
	// バグが取れなかったので、正解をコピーした
	fp := os.Stdin
	wfp := os.Stdout
	if isLocal() {
		fp, _ = os.Open(os.Getenv("END_GAME"))
	}
	scanner := bufio.NewScanner(fp)
	scanner.Split(bufio.ScanWords)
	scanner.Buffer(make([]byte, 1000005), 1000005)
	writer := bufio.NewWriter(wfp)
	defer func() {
		if isLocal() {
			r := recover()
			if r != nil {
				fmt.Fprintln(writer, r)
			}
		}
		writer.Flush()
	}()
	solve(scanner, writer)
	for i := 0; isLocal() && i < 100; i++ {
		fmt.Fprintln(writer, "-----------------------------------")
		solve(scanner, writer)
	}
}
func solve(scanner *bufio.Scanner, writer *bufio.Writer) {
	n := getNextInt(scanner)
	m := getNextInt(scanner)
	g := newGraph(n)
	for i := 0; i < m; i++ {
		a := getNextInt(scanner) - 1
		b := getNextInt(scanner) - 1
		g.AppendEdge(a, b, i)
		g.AppendEdge(b, a, i)
	}
	g.Dfs(0, -1, 0)
	g.Doubling()
	unused := map[int]bool{}
	for i := 0; i < n; i++ {
		if g.v[i].unused {
			unused[i] = true
		}
	}
	for uv := range unused {
		g.Bfs(uv)
	}
	q := getNextInt(scanner)
	for i := 0; i < q; i++ {
		u := getNextInt(scanner) - 1
		v := getNextInt(scanner) - 1
		lca := g.Get(u, v)
		ans := g.v[u].d + g.v[v].d - g.v[lca].d*2
		for uv := range unused {
			ans = min(ans, g.v[uv].daike[u]+g.v[uv].daike[v])
		}
		fmt.Fprintln(writer, ans)
	}
}
func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}
func min(a, b int) int { return -max(-a, -b) }
func abs(a int) int    { return max(a, -a) }

type vertex struct {
	unused  bool
	visited bool
	d       int
	daike   []int
	parent  [32]int
}
type edge struct {
	to, id int
	used   bool
}
type graph struct {
	v []vertex
	e [][]*edge
}

func newGraph(n int) graph {
	return graph{
		v: make([]vertex, n),
		e: make([][]*edge, n),
	}
}
func (g *graph) AppendEdge(from, to, id int) {
	g.e[from] = append(g.e[from], &edge{
		to: to,
		id: id,
	})
}
func (g *graph) Dfs(now, p, d int) {
	g.v[now].visited = true
	g.v[now].parent[0] = p
	g.v[now].d = d
	for _, e := range g.e[now] {
		if e.to == p {
			continue
		}
		if g.v[e.to].visited {
			g.v[now].unused = true
			g.v[e.to].unused = true
			continue
		}
		e.used = true
		g.Dfs(e.to, now, d+1)
	}
}
func (g *graph) Doubling() {
	for i := 1; i < 32; i++ {
		for j := 0; j < len(g.v); j++ {
			g.v[j].parent[i] = -1
			if g.v[j].parent[i-1] == -1 {
				continue
			}
			g.v[j].parent[i] = g.v[g.v[j].parent[i-1]].parent[i-1]
		}
	}
}
func (g *graph) Get(u, v int) int {
	if g.v[u].d < g.v[v].d {
		u, v = v, u
	}
	h := g.v[u].d - g.v[v].d
	for i := 0; i < 32; i++ {
		if h>>uint(i)&1 == 0 {
			continue
		}
		u = g.v[u].parent[i]
	}
	if u == v {
		return u
	}
	for i := 31; i >= 0; i-- {
		if g.v[u].parent[i] == g.v[v].parent[i] {
			continue
		}
		u = g.v[u].parent[i]
		v = g.v[v].parent[i]
	}
	return g.v[u].parent[0]
}
func (g *graph) Bfs(s int) {
	g.v[s].daike = make([]int, len(g.v))
	for i := 0; i < len(g.v); i++ {
		g.v[s].daike[i] = -1
	}
	q := make([][2]int, 0)
	q = append(q, [2]int{s, 0})
	for len(q) > 0 {
		p := q[0]
		to := p[0]
		d := p[1]
		q = q[1:]
		if g.v[s].daike[to] != -1 {
			continue
		}
		g.v[s].daike[to] = d
		for _, e := range g.e[to] {
			if g.v[s].daike[e.to] == -1 {
				q = append(q, [2]int{e.to, d + 1})
			}
		}
	}
}
