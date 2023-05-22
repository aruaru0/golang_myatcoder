def max_gift_value(N, M, A, B, D):
    # 青木君への贈り物の価値を昇順にソート
    A.sort()
    # すぬけ君への贈り物の価値を昇順にソート
    B.sort()

    max_value = -1  # 贈り物の価値の和の最大値
    j = 0  # すぬけ君への贈り物の現在のインデックス

    for i in range(N):
        # 青木君への贈り物の価値
        a = A[i]

        # すぬけ君への贈り物の価値の上限
        b_limit = a + D

        # 条件を満たすすぬけ君への贈り物を選ぶ
        while j < M and B[j] < a:
            j += 1

        # すぬけ君への贈り物の価値の差が D 以下の範囲で最大のものを選ぶ
        if j < M and B[j] <= b_limit:
            max_value = max(max_value, a + B[j])

    return max_value


# 入力の受け取り
N, M, D = map(int, input().split())
A = list(map(int, input().split()))
B = list(map(int, input().split()))

# 関数呼び出しと結果の出力
result = max_gift_value(N, M, A, B, D)
print(result)
