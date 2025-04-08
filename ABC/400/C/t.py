def count_unique_v(n):
    import math
    v_set = set()

    max_b = int((n // 2) ** 0.5) + 1  # なぜ2で割る？→ a >= 1 なので v >= 2*b^2
    for b in range(1, max_b):
        b2 = b * b
        v = b2 * 2
        while v <= n:
            v_set.add(v)
            v *= 2
    return len(v_set)

n = int(input())
print(count_unique_v(n))