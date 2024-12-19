import sys
input = sys.stdin.readline

N = int(input())
A = list(map(int, input().split()))

# v2(x): xを2で割れる最大回数

max_k = 24  # 十分大きな値（実際はlog2(10^7)<24）
sumO = [0]*(max_k+1)
count = [0]*(max_k+1)

# A_iごとにO_i, k_iを求めて集計
for x in A:
    k = (x & (-x)).bit_length() - 1
    O = x >> k
    print(k, O)
    if k > max_k:
        # 実際ここに来ることはないはず
        # 念のため拡張処理が必要ならここで対応
        pass
    sumO[k] += O
    count[k] += 1

# 同一k内計算
same_k_sum = 0
for k in range(max_k+1):
    if count[k] > 0:
        S = sumO[k]
        n = count[k]
        # 同一k内合計: S*(n+1)/2
        same_k_sum += S*(n+1)//2  # n+1, Sともに整数なので//2可能

# 異なるk間計算
diff_k_sum = 0
for k1 in range(max_k+1):
    if count[k1] == 0:
        continue
    for k2 in range(k1+1, max_k+1):
        if count[k2] == 0:
            continue
        power = 1 << (k2 - k1)
        diff_k_sum += count[k2]*sumO[k1] + power*count[k1]*sumO[k2]

ans = same_k_sum + diff_k_sum
print(ans)
