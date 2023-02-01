package main

import (
	"bufio"
	"fmt"
	"os"
)

const MOD int = 998244353

func components(n int, edges [][2]int) []int {
	adjList := make([][]int, n)
	for i := 0; i < n-1; i++ {
		u, v := edges[i][0], edges[i][1]
		adjList[u] = append(adjList[u], v)
		adjList[v] = append(adjList[v], u)
	}

	subSize := make([]int, n)
	dp := make([][][2]int, n)

	var dfs func(cur, pre int)
	dfs = func(cur, pre int) {
		subSize[cur] = 1
		dp[cur] = make([][2]int, 2)
		dp[cur][0][0] = 1
		dp[cur][1][1] = 1

		for _, next := range adjList[cur] {
			if next == pre {
				continue
			}

			dfs(next, cur)
			merged := make([][2]int, subSize[cur]+subSize[next]+1) // 当前不选/当前选
			for i := 0; i <= subSize[cur]; i++ {
				for j := 0; j <= subSize[next]; j++ {
					merged[i+j][0] += dp[cur][i][0] * (dp[next][j][0] + dp[next][j][1])
					merged[i+j][0] %= MOD

					merged[i+j][1] += dp[cur][i][1] * dp[next][j][0]
					merged[i+j][1] %= MOD

					if i+j-1 >= 0 {
						merged[i+j-1][1] += dp[cur][i][1] * dp[next][j][1]
						merged[i+j-1][1] %= MOD
					}
				}
			}

			subSize[cur] += subSize[next]
			dp[cur] = merged
		}
	}
	dfs(0, -1)

	res := make([]int, n+1)
	for i := 1; i <= n; i++ {
		res[i] = (dp[0][i][0] + dp[0][i][1]) % MOD
	}
	return res
}

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var n int
	fmt.Fscan(in, &n)
	edges := make([][2]int, n-1)
	for i := 0; i < n-1; i++ {
		fmt.Fscan(in, &edges[i][0], &edges[i][1])
		edges[i][0]--
		edges[i][1]--
	}

	res := components(n, edges)
	for i := 1; i <= n; i++ {
		fmt.Fprintln(out, res[i])
	}
}
