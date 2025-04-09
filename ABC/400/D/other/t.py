from collections import deque
import sys

def solve():
    input = sys.stdin.readline
    H, W = map(int, input().split())
    S = [input().rstrip('\n') for _ in range(H)]
    A, B, C, D = map(int, input().split())
    # 0-based index に変換
    A -= 1
    B -= 1
    C -= 1
    D -= 1

    INF = 10**9
    dist = [[INF]*W for _ in range(H)]

    # 移動方向ベクトル
    directions = [(-1,0), (1,0), (0,-1), (0,1)]

    # 0-1 BFS 用デック
    dq = deque()
    dist[A][B] = 0
    dq.append((A,B))

    while dq:
        r, c = dq.popleft()
        dcur = dist[r][c]
        # もし既に最適コストより大きいならスキップ
        if dcur > dist[r][c]:
            continue

        # 1) コスト 0 で行ける隣接道
        for dr, dc in directions:
            nr, nc = r+dr, c+dc
            if 0 <= nr < H and 0 <= nc < W:
                if S[nr][nc] == '.':
                    # 壁ではなく道ならコスト 0
                    if dist[nr][nc] > dcur:
                        dist[nr][nc] = dcur
                        dq.appendleft((nr,nc))  # 0 コストなので前に

        # 2) コスト 1 で壊せる方向探し
        for dr, dc in directions:
            # その方向へ最大2つまで壁を壊せる
            wall_count = 0
            nr, nc = r, c
            while True:
                nr += dr
                nc += dc
                if not (0 <= nr < H and 0 <= nc < W):
                    # 町の外に出れば打ち切り
                    break
                if S[nr][nc] == '#':
                    wall_count += 1
                    if wall_count > 2:
                        # 壁が 3 つ目になったらこれ以上は進めない
                        break
                else:
                    # 壁でない(道)マスが現れた時点で，
                    # 「そこに到達可能になる」＝コスト +1 の候補
                    # （ただし既により小さいコストで到達済みなら更新しない）
                    if dist[nr][nc] > dcur + 1:
                        dist[nr][nc] = dcur + 1
                        dq.append((nr,nc))  # コスト 1 なので後ろ
                    # 壁を壊しつつさらに奥を見ることはできるので続行
                    # （道だったので wall_count はそのまま）
                # さらにその先へ進んでみる(whileループ継続)

    # 答え
    print(dist[C][D])
