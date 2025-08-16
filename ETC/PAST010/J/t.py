import sys

MOD = 998244353

def main():
    import sys
    input = sys.stdin.read().split()
    idx = 0
    N = int(input[idx])
    idx += 1
    K = int(input[idx])
    idx += 1
    A = list(map(int, input[idx:idx+N]))
    
    if K == 1:
        print(0)
        return
    
    A.sort()
    k = K - 2
    
    max_n = N
    fact = [1] * (max_n + 1)
    for i in range(1, max_n + 1):
        fact[i] = fact[i-1] * i % MOD
    
    inv_fact = [1] * (max_n + 1)
    inv_fact[max_n] = pow(fact[max_n], MOD-2, MOD)
    for i in range(max_n - 1, -1, -1):
        inv_fact[i] = inv_fact[i+1] * (i+1) % MOD
    
    def comb(n, c):
        if n < 0 or c < 0 or n < c:
            return 0
        return fact[n] * inv_fact[c] % MOD * inv_fact[n - c] % MOD
    
    B = [0] * N
    for x in range(N):
        B[x] = comb(x, k)
    
    def ntt(a, invert):
        MOD = 998244353
        n = len(a)
        assert (n & (n-1)) == 0, "Length must be a power of two"
        root = pow(3, (MOD-1) // n, MOD)
        rev = [0] * n
        for i in range(n):
            rev[i] = rev[i >> 1] >> 1
            if i & 1:
                rev[i] |= n >> 1
        for i in range(n):
            if i < rev[i]:
                a[i], a[rev[i]] = a[rev[i]], a[i]
        for s in range(1, n.bit_length()):
            m = 1 << s
            w_m = pow(root, m >> 1, MOD)
            for i in range(0, n, m):
                w = 1
                for j in range(i, i + (m >> 1)):
                    x = a[j]
                    y = a[j + (m >> 1)] * w % MOD
                    a[j] = (x + y) % MOD
                    a[j + (m >> 1)] = (x - y) % MOD
                    w = w * w_m % MOD
        if invert:
            inv_n = pow(n, MOD-2, MOD)
            for i in range(n):
                a[i] = a[i] * inv_n % MOD
        return a
    
    def convolve(a, b):
        len_a = len(a)
        len_b = len(b)
        n = 1
        while n < len_a + len_b:
            n <<= 1
        a += [0] * (n - len_a)
        b += [0] * (n - len_b)
        ntt(a, False)
        ntt(b, False)
        c = [ (a[i] * b[i]) % MOD for i in range(n)]
        ntt(c, True)
        return c[:len_a + len_b -1]
    
    conv = convolve(A, B)
    
    total_sum = 0
    for j in range(N):
        c_j = comb(j, k + 1)
        sum1 = c_j
        if j == 0:
            sum2 = 0
        else:
            if (j-1) < len(conv):
                sum2 = conv[j-1]
            else:
                sum2 = 0
        contribution = (A[j] * sum1 - sum2) % MOD
        total_sum = (total_sum + contribution) % MOD
    
    denominator = comb(N, K)
    inv_denominator = pow(denominator, MOD-2, MOD)
    result = (total_sum * inv_denominator) % MOD
    print(result)

if __name__ == "__main__":
    main()
