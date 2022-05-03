from collections import defaultdict

N, X = map(int, input().split())

p = [[] for _ in range(N)]
for i in range(N) :
    l = list(map(int, input().split()))
    p[i] = l[1:]


x = defaultdict(int)
for e in p[0] :
    x[e]+=1

for i in range(1, N) :
    t = defaultdict(int)
    for e in p[i] :
        for v in x :
            t[e*v] += x[v]
    x,t = t,x

print(x[X])