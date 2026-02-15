import sys
from collections import deque

def solve() -> None:
    N, K = map(int, input().split())
    H = list(map(int, input().split()))

    mx = deque()          # 最大値用 (減少順)
    mn = deque()          # 最小値用 (増加順)
    ans = -10**18

    for i in range(N):
        # 最高気温のdequeを更新
        while mx and H[mx[-1]] <= H[i]:
            mx.pop()
        mx.append(i)

        # 最低気温のdequeを更新
        while mn and H[mn[-1]] >= H[i]:
            mn.pop()
        mn.append(i)

        # ウィンドウが K 以上になったら評価
        if i >= K - 1:
            left = i - K + 1

            # 左端外れたインデックスを除去
            if mx[0] < left:
                mx.popleft()
            if mn[0] < left:
                mn.popleft()

            diff = H[mx[0]] - H[mn[0]]
            if diff > ans:
                ans = diff

    print(ans)

if __name__ == "__main__":
    solve()
