def can_match_with_two_swaps(N, A, B):
    from collections import Counter

    # 要素の出現数が一致しなければ一致不可能
    if Counter(A) != Counter(B):
        return "No"

    # 差分があるインデックスだけ見る
    diff_indices = [i for i in range(N) if A[i] != B[i]]

    # 差分が0: すでに一致 → swap不可
    # 差分が2以下: swap1回で一致する可能性しかない → No
    # 差分が4以下ならswap2回で一致する可能性あり
    if len(diff_indices) > 4:
        return "No"

    from itertools import combinations

    # 差分部分のインデックスに絞って、2回のswapを全て試す
    for i in range(N - 1):
        for j in range(i + 1, N - 1):
            A_copy = A[:]
            # swap1
            A_copy[i], A_copy[i + 1] = A_copy[i + 1], A_copy[i]
            # swap2
            A_copy[j], A_copy[j + 1] = A_copy[j + 1], A_copy[j]
            if A_copy == B:
                return "Yes"

    return "No"


# 入力
N = int(input())
A = list(map(int, input().split()))
B = list(map(int, input().split()))

# 出力
print(can_match_with_two_swaps(N, A, B))
