def count_integers(N):
    import math
    
    # 集合を使って重複を避ける
    seen = set()
    
    for a in range(2, int(math.isqrt(N)) + 1):
        b = 2
        power = a ** b
        while power < N:
            seen.add(power)
            b += 1
            power = a ** b
    
    # 最後に 1^b = 1 を追加
    seen.add(1)
    
    return len(seen)

N = 10**18
result = count_integers(N)
print(result)