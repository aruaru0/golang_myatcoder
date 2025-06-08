from atcoder import dsu

n, m = map(int, input().split())
a = [()] * m
for i in range(m):
    u, v, w = map(int, input().split())
    a[i] = (u - 1, v - 1, w)
x = (1 << 30) - 1
for k in range(29, -1, -1):
    x ^= 1 << k
    d = dsu.DSU(n)
    for u, v, w in a:
        if (x | w) == x:
            d.merge(u, v)
    if not d.same(0, n - 1):
        x |= 1 << k
print(x)
