MAX = int(1e5)
MOD = 998244353

fac = [0]*MAX
finv = [0]*MAX
inv = [0]*MAX
fac[0] = fac[1] = finv[0] = finv[1] = inv[1] = 1

for i in range(2,MAX):
    fac[i] = fac[i-1]*i%MOD
    inv[i] = MOD-inv[MOD%i]*(MOD//i)%MOD
    finv[i] = finv[i-1]*inv[i]%MOD
def binom(n,k):
    if n<0 or k<0 or n<k:
        return 0
    return fac[n]*finv[k]%MOD*finv[n-k]%MOD

N,M,B,W = map(int,input().split())
f1 = [[0]*(M+1) for _ in range(N+1)]
f2 = [[0]*(M+1) for _ in range(N+1)]

def init(x,f):
    for n in range(1,N+1):
        for m in range(1,M+1):
            for i in range(0,n+1):
                for j in range(0,m+1):
                    v = 1
                    if (i+j)%2:
                        v = -1
                    f[n][m] += v * binom(n,i) * binom(m,j) * binom((n-i)*(m-j),x)
                    f[n][m] %= MOD
                    print(v, binom(n,i) , binom(m,j) , binom((n-i)*(m-j),x))

init(B,f1)
init(W,f2)

# ans = 0

# for i in range(1,N+1):
#     for j in range(1,N-i+1):
#         for k in range(1,M+1):
#             for l in range(1,M-k+1):
#                 ans += binom(N,i) * binom(N-i,j) * binom(M,k) * binom(M-k,l) * f1[i][k] * f2[j][l]
#                 ans %= MOD

# print(ans)