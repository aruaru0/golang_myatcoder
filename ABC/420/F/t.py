import sys

input = sys.stdin.readline
N, M, K = map(int, input().split())
S = [input().strip() for _ in range(N)]

# 前処理: 各マスから上方向に連続する '.' の数を数える
up = [[0] * M for _ in range(N)]
for j in range(M):
    for i in range(N):
        if S[i][j] == '.':
            if i == 0:
                up[i][j] = 1
            else:
                up[i][j] = up[i-1][j] + 1

ans = 0
# 各行を下端として固定
for i in range(N):
    h = up[i]
    
    # スタックには (高さ, 開始インデックス) を格納
    st = []
    # `current_sum` は、現在の j を右端とする長方形の総数
    current_sum = 0
    
    for j in range(M):
        # 新しいバー h[j] を追加する
        # h[j] より高いバーをスタックから取り除く
        p_new = j
        while st and st[-1][0] > h[j]:
            v, p = st.pop()
            width = j - p
            
            # 取り除いたバーの貢献を current_sum から引く
            # 高さ v, 幅 1..width の長方形の貢献
            max_h = K // width
            if v > max_h:
                pass 

            current_sum -= v * width
            p_new = p
        
        # h[j] をスタックに追加
        width_new = j - p_new + 1
        current_sum += h[j] * width_new
        st.append((h[j], p_new))
        
for i in range(N):
    h = up[i]
    for j in range(M):
        if S[i][j] == '#':
            continue
        
        min_h = h[j]
        for l in range(j, -1, -1):
            min_h = min(min_h, h[l])
            if min_h == 0:
                break
            
            width = j - l + 1
            max_h_by_area = K // width
            ans += min(min_h, max_h_by_area)

print(ans)
