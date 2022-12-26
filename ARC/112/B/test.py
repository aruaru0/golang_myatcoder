b, c = map(int, input().split())

segments = []
# x x
segments.append((b, b-c//2))
# x o
segments.append((-b, -(b-(c-1)//2)))
# o x
segments.append((-b, -b-(c-1)//2))
# o o
if(c >= 2):segments.append((b, -(-b-(c-2)//2)))

print(segments)

imos = []
for (l, r) in segments:
    if l > r:l,r = r,l
    imos.append((l, +1))
    imos.append((r+1, -1))

last = -(1e20)
cnt = 0

imos.sort()
ans = 0
for (pos, diff) in imos:
    if cnt > 0:ans += pos - last
    last = pos
    cnt += diff

print(ans)
