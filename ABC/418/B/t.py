# -*- coding: utf-8 -*-
import sys

def main() -> None:
    S = input().strip()
    n = len(S)

    # prefix sum of 't'
    pref = [0] * (n + 1)
    for i in range(n):
        pref[i + 1] = pref[i] + (1 if S[i] == 't' else 0)

    ans = 0.0

    for i in range(n):
        if S[i] != 't':
            continue
        for j in range(i + 2, n):          # length >= 3 -> j - i >= 2
            if S[j] != 't':
                continue
            L = j - i + 1
            x = pref[j + 1] - pref[i]
            r = (x - 2) / (L - 2)
            if r > ans:
                ans = r

    # 出力（小数点20桁まで）
    print(f"{ans:.20f}")

if __name__ == "__main__":
    main()
