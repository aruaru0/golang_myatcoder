import sys
sys.setrecursionlimit(10**6)
MOD = 998244353


def dfs(cur=0, pre=-1):
    dp0, dp1 = [1, 0], [0, 1]

    for nxt in adj[cur]:
        if nxt == pre:  # not go back
            continue

        rs0, rs1 = dfs(nxt, cur)
        n, m = len(dp0), len(rs0)
        nx0, nx1 = [0]*(n+m-1), [0]*(n+m-1)

        for i in range(n):
            for j in range(m-1, 0, -1):  # j:[m-1,1]
                nx0[i+j] += dp0[i] * (rs0[j] + rs1[j])
                nx1[i+j] += dp1[i] * rs0[j]
                nx0[i+j] %= MOD
                nx1[i+j] %= MOD
                nx1[i+j-1] += dp1[i] * rs1[j]
            nx0[i] += dp0[i] * (rs0[0] + rs1[0])  # j=0
            nx1[i] += dp1[i] * rs0[0]  # j=0
            nx0[i] %= MOD
            nx1[i] %= MOD

        dp0, dp1 = nx0, nx1

    return dp0, dp1


N = int(input())
adj = [[] for _ in range(N)]
for _ in range(N-1):
    a, b = [int(x)-1 for x in input().split()]
    adj[a].append(b)
    adj[b].append(a)


#print(*map(lambda p: sum(p) % MOD, [*zip(*dfs())][1:]), sep='\n')
print([*zip(*dfs())][1:])
