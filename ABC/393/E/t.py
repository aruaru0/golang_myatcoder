import sys
import math

# 入力の受け取り
N, K = map(int, sys.stdin.readline().split())
A = list(map(int, sys.stdin.readline().split()))

MAX_A = 10**6

# 数列の各値の出現回数をカウント
count = [0] * (MAX_A + 1)
for a in A:
    count[a] += 1

# 各値の累積頻度を求める（gの倍数を持つ個数）
prefix_count = [0] * (MAX_A + 1)
for g in range(1, MAX_A + 1):
    for multiple in range(g, MAX_A + 1, g):
        prefix_count[g] += count[multiple]

# 各 i について最大の GCD を求める
result = [0] * N
for i in range(N):
    a_i = A[i]
    best_gcd = 1
    
    # a_i を含む GCD の候補を大きい順に探索
    for g in range(a_i, MAX_A + 1, a_i):
        if prefix_count[g] >= K:
            best_gcd = g
    
    result[i] = best_gcd

# 結果の出力
sys.stdout.write("\n".join(map(str, result)) + "\n")
