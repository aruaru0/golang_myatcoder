from heapq import heappop, heappush

N = int(input())
AB = [tuple(map(int,input().split())) for _ in range(N)]
AB.sort(key=lambda x: x[1])

print(AB)

hq = []
for _ in range(2):
    a,b = AB.pop()
    heappush(hq, a + b)
    print(a, b, a+b)

ans = 10 ** 18

print(hq)
while AB:
    a, b = AB.pop()
    x = heappop(hq)
    print(a, b, x)
    tmp = a + x + hq[0]
    ans = min(ans, tmp)
    heappush(hq, x)
    heappush(hq, a + b)
print(ans)
