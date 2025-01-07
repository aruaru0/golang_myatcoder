# https://atcoder.jp/contests/abc387/tasks/abc387_c
import sys; input: lambda _: sys.stdin.readline().rstrip()
# import pypyjit; pypyjit.set_param('max_unroll_recursion=-1')
sys.setrecursionlimit(10001000)
int1=lambda x: int(x) - 1
L, R = map(int, input().split())
L -= 1

def f(x):
    sx = str(x)
    n = len(sx)
    dp = [[[0] * 2 for _ in range(10)] for __ in range(n+1)]
    for j in range(int(sx[0])):
        dp[1][j][1] = 1
    dp[1][int(sx[0])][0] = 1

    for i in range(1, n):
        s = int(sx[i])
        for j in range(10):
            dp[i+1][j][1] = dp[i][j][1] * j + 1
            dp[i+1][j][1] += dp[i][j][0] * min(j, s)
            if j > s:
                dp[i+1][j][0] = dp[i][j][0]
    ret = 0
    for j in range(10):
        ret += sum(dp[-1][j])
    return ret

# print(f(R))
# print(f(L))
print(f(R) - f(L))