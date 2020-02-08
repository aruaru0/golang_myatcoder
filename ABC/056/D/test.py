
N, K = map(int, input().split())
a = list(map(int, input().split()))
 
a = sorted(a)
ans = N
t = 0

for i in range(N-1, -1, -1) :
    if t+a[i] < K :
        t += a[i]
    else :
        ans = min(ans, i)
    print(i, ans, t, a[i])
 
print(ans)
