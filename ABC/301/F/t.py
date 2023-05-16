MOD = 998244353

def count_ddos_strings(S):
    n = len(S)
    dp = [[[0] * 2 for _ in range(2)] for _ in range(n + 1)]
    dp[0][0][0] = 1

    for i in range(1, n + 1):
        for j in range(2):
            for k in range(2):
                if S[i - 1].isupper():  # 英大文字の場合
                    dp[i][0][k] = (dp[i][0][k] + dp[i - 1][0][k] + dp[i - 1][1][k]) % MOD
                    dp[i][1][k] = (dp[i][1][k] + dp[i - 1][0][k - 1] + dp[i - 1][1][k - 1]) % MOD
                elif S[i - 1].islower():  # 英小文字の場合
                    dp[i][0][k] = (dp[i][0][k] + dp[i - 1][0][k] + dp[i - 1][1][k]) % MOD
                    dp[i][1][k] = (dp[i][1][k] + dp[i - 1][0][k] + dp[i - 1][1][k - 1]) % MOD
                else:  # '?' の場合
                    dp[i][j][k] = (dp[i][j][k] + dp[i - 1][j][k] + dp[i - 1][j][k - 1]) % MOD

    ans = (dp[n][0][0] + dp[n][0][1] + dp[n][1][0] + dp[n][1][1]) % MOD
    return ans

# 入力例
S = input()

# DDoS型文字列を部分列に含まないパターンの個数を求める
result = count_ddos_strings(S)
print(result)
