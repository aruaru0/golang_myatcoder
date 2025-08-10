import sys

sys.setrecursionlimit(1 << 25)
MOD = 998244353


class Seg:
    # lazy segment tree, range sum + range assign to constant
    def __init__(self, a):
        self.n = len(a)
        sz = 4 * self.n
        self.s = [0] * sz          # segment sums
        self.lz = [None] * sz      # pending assignment
        self._b(1, 0, self.n - 1, a)

    # build
    def _b(self, i, l, r, a):
        if l == r:
            self.s[i] = a[l] % MOD
            return
        m = (l + r) // 2
        self._b(i * 2, l, m, a)
        self._b(i * 2 + 1, m + 1, r, a)
        self.s[i] = (self.s[i * 2] + self.s[i * 2 + 1]) % MOD

    # apply assignment to node i covering [l,r]
    def _ap(self, i, l, r, v):
        self.s[i] = v * (r - l + 1) % MOD
        self.lz[i] = v

    # push lazy tag downwards
    def _ps(self, i, l, r):
        if self.lz[i] is None or l == r:
            return
        m = (l + r) // 2
        v = self.lz[i]
        self._ap(i * 2, l, m, v)
        self._ap(i * 2 + 1, m + 1, r, v)
        self.lz[i] = None

    # range sum query
    def qs(self, ql, qr):
        return self._qs(1, 0, self.n - 1, ql, qr)

    def _qs(self, i, l, r, ql, qr):
        if ql <= l and r <= qr:
            return self.s[i]
        self._ps(i, l, r)
        m = (l + r) // 2
        res = 0
        if ql <= m:
            res += self._qs(i * 2, l, m, ql, qr)
        if qr > m:
            res += self._qs(i * 2 + 1, m + 1, r, ql, qr)
        return res % MOD

    # range assign to constant v
    def qa(self, ql, qr, v):
        self._qa(1, 0, self.n - 1, ql, qr, v)

    def _qa(self, i, l, r, ql, qr, v):
        if ql <= l and r <= qr:
            self._ap(i, l, r, v)
            return
        self._ps(i, l, r)
        m = (l + r) // 2
        if ql <= m:
            self._qa(i * 2, l, m, ql, qr, v)
        if qr > m:
            self._qa(i * 2 + 1, m + 1, r, ql, qr, v)
        self.s[i] = (self.s[i * 2] + self.s[i * 2 + 1]) % MOD

    # collect final values into a list
    def get(self):
        ans = [0] * self.n
        self._gt(1, 0, self.n - 1, ans)
        return ans

    def _gt(self, i, l, r, ans):
        if l == r:
            ans[l] = self.s[i]
            return
        self._ps(i, l, r)
        m = (l + r) // 2
        self._gt(i * 2, l, m, ans)
        self._gt(i * 2 + 1, m + 1, r, ans)


def main():
    N, M = map(int, input().split())
    a = list(map(int, input().split()))

    # modular inverses of 1 … N
    iv = [0] * (N + 2)
    iv[1] = 1
    for i in range(2, N + 1):
        iv[i] = (MOD - (MOD // i) * iv[MOD % i] % MOD) % MOD

    seg = Seg(a)

    for _ in range(M):
        L, R = map(int, input().split())
        l = L - 1
        r = R - 1
        tot = seg.qs(l, r)
        ln = r - l + 1
        nv = tot * iv[ln] % MOD
        seg.qa(l, r, nv)

    ans = seg.get()
    print(' '.join(str(x) for x in ans))


if __name__ == "__main__":
    main()
