import sys

def main():
    """
    問題を解くメインの関数
    """
    # 標準入力の高速化
    input = sys.stdin.readline

    try:
        N, Q = map(int, input().split())
    except ValueError:
        # ローカル環境でのテスト時に、入力の最後に空行があるとエラーになるのを防ぐ
        return

    # データ構造の初期化
    # nodes[node_id] = (parent_id, added_string)
    # 根ノード(ID:0)は空文字列を表し、親はいない(ID:-1)
    nodes = {0: (-1, "")}
    next_node_id = 1
    
    # サーバーと各PCが現在どのノードを参照しているかを示すポインタ
    server_ptr = 0
    pc_ptrs = [0] * (N + 1) # PC番号は1-indexedのため N+1

    # Q個のクエリを処理
    for _ in range(Q):
        query = input().split()
        query_type = int(query[0])

        if query_type == 1:
            # クエリタイプ 1: PC p <- Server
            # PC p のポインタをサーバーのポインタで上書きする
            p = int(query[1])
            pc_ptrs[p] = server_ptr
        
        elif query_type == 2:
            # クエリタイプ 2: PC p <- PC p + s
            # PC p が参照するノードを親として新しいノードを作成し、
            # PC p のポインタをその新しいノードに向ける
            p = int(query[1])
            s = query[2]
            
            parent_id = pc_ptrs[p]
            nodes[next_node_id] = (parent_id, s)
            pc_ptrs[p] = next_node_id
            next_node_id += 1
            
        elif query_type == 3:
            # クエリタイプ 3: Server <- PC p
            # サーバーのポインタを PC p のポインタで上書きする
            p = int(query[1])
            server_ptr = pc_ptrs[p]

    # 最終的なサーバーの文字列を復元
    result_parts = []
    current_ptr = server_ptr
    
    # サーバーが参照するノードから根ノード(-1)に到達するまで親を遡る
    while current_ptr != -1:
        parent_id, s = nodes[current_ptr]
        result_parts.append(s)
        current_ptr = parent_id
        
    # 遡って集めた文字列の断片は逆順になっているため、
    # さらに逆順にして結合し、元の順序の文字列を生成する
    final_string = "".join(reversed(result_parts))
    
    # 結果を出力
    print(final_string)

if __name__ == "__main__":
    main()