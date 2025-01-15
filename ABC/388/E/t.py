def max_kagami_mochi(n, a):
    i, j = 0, n // 2  # ポインタの初期化
    count = 0         # 鏡餅の数
    
    while i < n // 2 and j < n:
        # 条件を満たす場合
        if a[i] * 2 <= a[j]:
            count += 1
            i += 1  # 小さい餅を次に進める
            j += 1  # 大きい餅も次に進める
        else:
            # 条件を満たさない場合は大きい餅を次に進める
            j += 1
    
    return count

# 入力の受け取り
n = int(input())
a = list(map(int, input().split()))

# 出力
print(max_kagami_mochi(n, a))
