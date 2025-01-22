def count_squares_in_circle(R):
    # 半径を2倍して整数化
    R2 = 2 * R  # 2R
    R2_squared = R2 * R2  # (2R)^2
    
    count = 0  # 結果のカウント
    
    # i の範囲を [-R2, R2] に設定
    for i in range(-R2, R2 + 1, 2):  # i を 2 ずつ増加
        i_term = (i + 1) ** 2  # (2i+1)^2
        
        # j の範囲を二分探索で求める
        low, high = -R2, R2
        j_min, j_max = None, None
        
        # j_min を探索
        while low <= high:
            mid = (low + high) // 2
            j_term = (mid + 1) ** 2  # (2j+1)^2
            if i_term + j_term <= R2_squared:
                j_min = mid
                high = mid - 2  # 2 ずつ減少
            else:
                low = mid + 2  # 2 ずつ増加
        
        # j_max を探索
        low, high = -R2, R2
        while low <= high:
            mid = (low + high) // 2
            j_term = (mid + 1) ** 2  # (2j+1)^2
            if i_term + j_term <= R2_squared:
                j_max = mid
                low = mid + 2  # 2 ずつ増加
            else:
                high = mid - 2  # 2 ずつ減少
        
        # j_min と j_max が有効ならカウントを追加
        if j_min is not None and j_max is not None:
            count += (j_max - j_min) // 2 + 1
    
    return count

# 入力例
R = int(input())
result = count_squares_in_circle(R)
print(result)