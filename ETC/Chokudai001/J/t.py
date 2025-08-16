# fenwick tree
class FenwickTree:
    def __init__(self, n):
        self.n = n
        self.data = [0] * (n + 1)

    def add(self, i, x):
        while i <= self.n:
            self.data[i] += x
            i += i & -i

    def sum(self, i):
        s = 0
        while i > 0:
            s += self.data[i]
            i -= i & -i
        return s

    def range_sum(self, l, r):
        return self.sum(r) - self.sum(l - 1)
    

n = int(input())
a = list(map(int, input().split()))
max_a = max(a)
bit = FenwickTree(max_a)


cnt= 0
for i in range(n):
    cnt += bit.range_sum(a[i] + 1, max_a)
    bit.add(a[i], 1)

print(cnt)