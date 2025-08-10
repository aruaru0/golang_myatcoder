import sys

def solve():
    """
    問題を解くメインの関数
    """
    # 入力の高速化
    input = sys.stdin.readline

    # 入力
    N = int(input())
    PAB = [list(map(int, input().split())) for _ in range(N)]
    Q = int(input())
    queries = [int(input()) for _ in range(Q)]

    # 閾値Mを設定
    max_p = 0
    if N > 0:
        for p, _, _ in PAB:
            max_p = max(max_p, p)
    M = max_p

    # B_i の累積和（後ろから）を計算
    # S[i] = B_i + B_{i+1} + ... + B_{N-1}
    S = [0] * (N + 1)
    for i in range(N - 1, -1, -1):
        S[i] = S[i + 1] + PAB[i][2]

    # DPテーブル: memo[i][t] = プレゼントiの直前のテンションがtの時の最終テンション
    memo = [[0] * (M + 1) for _ in range(N + 1)]

    # DPのベースケース: 全てのプレゼントをもらった後
    for t in range(M + 1):
        memo[N][t] = t

    # DPテーブルを埋める: i = N-1 ... 0
    for i in range(N - 1, -1, -1):
        p, a, b = PAB[i]
        for t in range(M + 1):
            t_next = 0
            if t <= p:
                t_next = t + a
            else:
                t_next = max(0, t - b)

            if t_next <= M:
                memo[i][t] = memo[i + 1][t_next]
            else:  # t_next > M (テンションが高い状態)
                # テンションがM以下になる最初のプレゼントkを二分探索で探す
                low = i + 1
                high = N + 1

                while low+1 != high:
                    k_mid = (low + high) // 2
                    sum_b_section = S[i + 1] - S[k_mid]
                    t_at_k = t_next - sum_b_section
                    if t_at_k <= M:
                        high = k_mid
                    else:
                        low = k_mid 
                
                k_star = high

                if k_star > N:  # テンションがM以下にならない
                    final_t = t_next - S[i + 1]
                    memo[i][t] = max(0, final_t)
                else:
                    sum_b_to_k = S[i + 1] - S[k_star]
                    t_at_k = t_next - sum_b_to_k
                    # FIX: Apply the max(0, ...) rule to prevent negative index
                    final_t_at_k = max(0, t_at_k)
                    memo[i][t] = memo[k_star][int(final_t_at_k)]
    
    # クエリ処理
    for x_q in queries:
        if x_q <= M:
            print(memo[0][int(x_q)])
        else:  # x_q > M (初期テンションが高い)
            # テンションがM以下になる最初のプレゼントkを二分探索
            low = 0
            high = N + 1

            while low+1 != high:
                k_mid = (low + high) // 2
                sum_b_section = S[0] - S[k_mid]
                t_at_k = x_q - sum_b_section
                if t_at_k <= M:
                    high = k_mid
                else:
                    low = k_mid
            
            k_star = high
            
            if k_star > N:  # テンションがM以下にならない
                final_t = x_q - S[0]
                print(max(0, final_t))
            else:
                sum_b_to_k = S[0] - S[k_star]
                t_at_k = x_q - sum_b_to_k
                # FIX: Apply the max(0, ...) rule to prevent negative index
                final_t_at_k = max(0, t_at_k)
                print(memo[k_star][int(final_t_at_k)])

solve()