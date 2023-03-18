from sys import stdin; input = stdin.readline
#from collections import deque
INF = float('inf')

N, M = map(int, input().split())
A = list(map(int, input().split()))
XYZ = [list(map(int, input().split())) for _ in range(M)]

p = 0
for a in A[::-1]: p <<= 1; p |= a

print(p)

G = [[] for _ in range(1<<N)]
for i in range(M):
	x, y, z = XYZ[i]
	x, y, z = x-1, y-1, z-1
	q = 1<<x | 1<<y | 1<<z
	for i in range(1<<N): G[i].append(i^q)

print(G)

dist = [INF]*(1<<N); dist[p] = 0
#deq = deque([p])
#while deq:
deq = [p]
for here in deq:
	print(here)
	#here = deq.popleft()
	for nxt in G[here]:
		if dist[here]+1 < dist[nxt]:
			dist[nxt] = dist[here]+1
			deq.append(nxt)
ans = dist[(1<<N)-1]
if ans == INF: ans = -1
print(ans)
