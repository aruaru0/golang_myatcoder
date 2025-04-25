import sys

def main():
    N, X = map(int, sys.stdin.readline().split())
    S = []
    C = []
    P = []
    for _ in range(N):
        s, c, p = map(int, sys.stdin.readline().split())
        S.append(s)
        C.append(c)
        P.append(p)

    size = 1 << N
    dp = [[0.0] * (X + 1) for _ in range(size)]

    for x_rem in range(X + 1):
        for bitmask in range(size):
            current_max = dp[bitmask][x_rem]
            for i in range(N):
                if not (bitmask & (1 << i)):
                    cost = C[i]
                    if x_rem < cost:
                        continue
                    new_x = x_rem - cost
                    prob = P[i] / 100.0
                    new_bm = bitmask | (1 << i)
                    val = prob * (S[i] + dp[new_bm][new_x])
                    val += (1 - prob) * dp[bitmask][new_x]
                    if val > current_max:
                        current_max = val
            dp[bitmask][x_rem] = current_max

    print("{0:.12f}".format(dp[0][X]))

if __name__ == "__main__":
    main()
